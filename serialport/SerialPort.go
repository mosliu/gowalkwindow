package serialport

import (
    "time"
    "github.com/tarm/serial"
    "bytes"
    "reflect"
)

/*
    get raw data from serial port.
    serialPort:serial port
    recvChan: receive chan which put the raw data in
 */
func RecvSerialData(serialPort *serial.Port, recvChan chan []byte) {
    buf := make([]byte, 128)
    for {
        //buf := make([]byte, 128)
        n, err := serialPort.Read(buf)
        if err != nil {
            log.Fatal(err)
        }
        log.Debugf("get %d bytes at %s ", n, time.Now())
        log.Debug("Receive :", buf[:n])
        var s = make([]byte, n)
        copy(s, buf[:n])
        recvChan <- s
        //for i := 0; i < n; i++ {
        //    recvChan <- buf[i];
        //}
    }
}

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
        for buffer.Len() >= pi.MinLength {
            //值传递！
            pkt, ok := ExtractPacket(buffer.Bytes(), pi)
            if !ok {
                abyte,_ := buffer.ReadByte()
                log.Debugf("Drop one byte: %X",abyte)
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
    pkt.New(pktbuf.Bytes())
    return pkt, true
}

func SendSerialData(serialPort *serial.Port, sendChan chan []byte) {

    for b := range sendChan {

        n, err := serialPort.Write(b)
        if err != nil {
            log.Fatal(err)
        }
        log.Debugf("Send [%d] bytes :%X ", n, b)
    }
}

func OpenPort(PortNum string) (serialPort *serial.Port){
    serialConfig := &serial.Config{Name: PortNum, Baud: 115200}
    serialPort, err := serial.OpenPort(serialConfig)
    if err != nil {
        log.Fatal(err)
    }
    return serialPort

}
