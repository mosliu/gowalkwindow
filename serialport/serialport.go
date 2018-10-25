package serialport

import (
    "bytes"
    "errors"
    . "github.com/mosliu/gowalkwindow/handlers"
    "github.com/tarm/serial"
    "reflect"
    "time"
)

// 保存一个串口的信息。配置信息，对应实例，发送和接收通道
type CommPort struct {
    //port   *serial.Port
    Port       *serial.Port   `desc:"串口"`
    Config     *serial.Config `desc:"串口配置"`
    IsOpen     bool           `desc:"串口是否打开"`
    RecvChan   chan []byte    `desc:"接收通道"`
    SendChan   chan []byte    `desc:"发送通道"`
    PacketChan chan Packet    `desc:"解析后数据包通道？"`
    handlers   []Handler
}

// 打开串口
// 初始化发送接收通道
func (p *CommPort) OpenPort() (err error) {
    //FIXME p.Config = nil?
    if p.Config == nil {
        err = errors.New("port config is null")
        return err
    }
    // 打开串口
    pt, err := serial.OpenPort(p.Config)
    if err != nil {
        logf.WithField("config", p.Config).Errorf("Port Open Failure:%v\r\n", err)
        p.IsOpen = false
        return err
    } else {
        p.IsOpen = true
        p.Port = pt
    }
    // 初始化发送接收通道
    p.RecvChan = make(chan []byte)
    p.SendChan = make(chan []byte)
    p.PacketChan = make(chan Packet)

    go p.StartSendChannel()
    go p.StartRecvChannel()
    return err
}

// 关闭串口
func (p *CommPort) ClosePort() (err error) {
    // 关闭串口
    if p.IsOpen == false {
        err = errors.New("port is closed and can not close again")
        return err
    }
    p.IsOpen = false
    err = p.Port.Flush()
    if err != nil {
        logf.Warnf("Port Flush Failure:%v\r\n", err)
        return err
    }
    err = p.Port.Close()
    if err != nil {
        logf.Warnf("Port Close Failure:%v\r\n", err)
        return err
    }
    //关闭channel
    //对于不再使用的通道不必显示关闭。如果没有goroutine引用这个通道，这个通道就会被垃圾回收? 是否不需要关闭么？
    close(p.RecvChan)
    close(p.SendChan)
    close(p.PacketChan)

    return err
}

//Open send Channel
// SendChan 的消费者
func (p *CommPort) StartSendChannel() {
    //FIXME if port is nil?
    for b := range p.SendChan {
        n, err := p.Port.Write(b)
        if err != nil {
            log.Fatal(err)
        }
        log.Debugf("Send [%d] bytes :%X ", n, b)
    }

    log.Debug("Send Channel Closed.")
}

/*
    get raw data from serial port.
    serialPort:serial port
    recvChan: receive chan which put the raw data in
 */
// RecvChan 的生产者
func (p *CommPort) StartRecvChannel() {
    //FIXME if port is nil?
    if p.IsOpen == false {
        log.Debug("Can Not Start Receive Channel For Close Port.")
        return
    }
    buf := make([]byte, 128)
    for {
        if p.IsOpen == false {
            log.Debug("Port closed. Stop receiving Serial Data.")
            break
        }
        //buf := make([]byte, 128)
        // 会阻塞
        n, err := p.Port.Read(buf)
        if err != nil {
            log.Error(err)
        }
        if n == 0 {
            continue
        }
        log.Debugf("get %d bytes at %s ", n, time.Now())
        log.Debug("Receive :", buf[:n])
        var s = make([]byte, n)
        copy(s, buf[:n])
        p.RecvChan <- s

        //for i := 0; i < n; i++ {
        //    recvChan <- buf[i];
        //}
    }
}

//开始处理队列运行
func (p *CommPort) StartHandle() {

    p.prepareHandlerQuene()

    go p.startFilter()

}

//准备处理器序列
func (p *CommPort) prepareHandlerQuene() {
    //p.addHandler(&SimplePrintHandler{})
    //p.addHandler(NewSimplePrintHandler())
    p.addHandler(NewVacV1Handler())
}

//adds a handler to the queue
func (p *CommPort) addHandler(handler Handler) {
    p.handlers = append(p.handlers, handler)
}

// Unsubscribe removes a handler.
func (p *CommPort) removeHandler(handler Handler) {
    for i, h := range p.handlers {
        if h == handler {
            p.handlers = append(p.handlers[:i], p.handlers[i+1:]...)
            return
        }
    }
}

// 先过滤原始数据，匹配相应的处理器。
func (p *CommPort) startFilter() {
    //var buffer bytes.Buffer
    buffer := new(bytes.Buffer)

    for b := range p.RecvChan {
        //fmt.Println("b:",b)
        //接收数据，放进buffer中。
        buffer.Write(b)
        log.Debugf("Receive %d bytes at %s ", len(b), time.Now())

        // 始终循环，用于一次接受多帧时，或者帧有问题时。
        for {
            //没有缓冲了, 跳出
            if buffer.Len() == 0 {
                break
            }
            inbytes := buffer.Bytes()
            log.Debugf("current buf count:%d,content: %X", len(inbytes), inbytes)
            var judge = REJECT
            //遍历所有handlers
            for _, h := range p.handlers {
                ok, pkt := h.Judge(inbytes)
                switch ok {
                case REJECT:
                    log.Debugf("%v rejected", h.GetName())
                case NEEDMORE:
                    log.Debugf("%v need more bytes", h.GetName())
                    if judge == REJECT {
                        judge = NEEDMORE
                    }
                case ACCEPT:
                    log.Infof("%v accepted:%+v", h.GetName(),pkt)
                    judge = ACCEPT
                }
                if judge == ACCEPT {
                    h.Handle(pkt)
                    buffer.Next(pkt.Len())
                    break
                }
            }

            //否则下一个循环，对应一次收多条信息
            //所有的处理器都不满足，可能有不需要的字节，抛弃一位
            if judge == REJECT {
                buffer.ReadByte()
                log.Debug("All handlers rejected, drop one byte")
            }
            // 没有accept的 有needmore的，等待更多字节来处理。
            if judge == NEEDMORE {
                break
            }
        }
        // 判断是否处理

        //for buffer.Len() >= pi.MinLength {
        //    //值传递！
        //    pkt, ok := ExtractPacket(buffer.Bytes(), pi)
        //    if !ok {
        //        abyte, _ := buffer.ReadByte()
        //        log.Debugf("Drop one byte: %X", abyte)
        //    } else {
        //        buffer.Next(pkt.Len())
        //        packetChan <- pkt
        //    }
        //}
    }
}

//--------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------------------------------------------------------------------

/**
    judge if the raw data match the packet format.
    push the packet into the packetChan
 */
func HandleRawData(recvChan chan []byte, packetChan chan Packet, pi PacketInfo) {
    //var buffer bytes.Buffer
    buffer := new(bytes.Buffer)

    for b := range recvChan {
        //fmt.Println("b:",b)
        buffer.Write(b)
        log.Debugf("handle %d bytes at %s ", len(b), time.Now())
        log.Debug(buffer.Bytes())
        //for buffer.Len() >= pi.MinLength {
        for buffer.Len() >= 10 {
            //值传递！
            pkt, ok := ExtractPacket(buffer.Bytes(), pi)
            if !ok {
                abyte, _ := buffer.ReadByte()
                log.Debugf("Drop one byte: %X", abyte)
            } else {
                buffer.Next(pkt.Len())
                packetChan <- pkt
            }
        }
    }
}

/**
   extract packet from comm buffer
 */
func ExtractPacket(bufByte []byte, pi PacketInfo) (pkt Packet, ok bool) {
    buf := bytes.NewBuffer(bufByte)
    pktbuf := new(bytes.Buffer)
    var head, body, tail []byte
    //if has fixed formatted head
    if pi.HasFixedHead {
        head = buf.Next(len(pi.HeadStyle))
        if !reflect.DeepEqual(head, pi.HeadStyle) {
            return pkt, false
        }
        pktbuf.Write(head)
    }
    if pi.BodyLength > 0 {
        body = buf.Next(pi.BodyLength)
        pktbuf.Write(body)
    } else {
        //do nothing
    }

    if pi.HasFixedTail {
        tail = buf.Next(len(pi.TailStyle))
        if !reflect.DeepEqual(tail, pi.TailStyle) {
            return pkt, false
        }
        pktbuf.Write(tail)
    }
    //pkt.New(pktbuf.Bytes())
    return pkt, true
}
