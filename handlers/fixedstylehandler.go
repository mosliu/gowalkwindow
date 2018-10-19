package handlers

import "bytes"

type FixedStyleHandler struct {
	BaseHandler
	//减少循环中的计算量
	MinLength int
	InPi      *PacketInfo
}

//51->pc
//AA AA X1 X2 X3 X4 X5 X6 X7 X8 X9 SM BB BB

//in: 51 -> pc
var FixedStyleInPi = PacketInfo{
	BodyLength:    9,
	HasFixedHead:  true,
	HasFixedTail:  true,
	HeadStyle:     []byte{0xAA, 0xAA},
	TailStyle:     []byte{0xBB, 0xBB},
	HasCheckSum:   true,
	CheckSumStyle: []byte{0x0},
}

//pc->51
//AA AA X1 X2 X3 X4 SM BB BB

//out: pc -> 51
var FixedStyleOutPi = PacketInfo{
	BodyLength:   10,
	HasFixedHead: false,
	HasFixedTail: false,
	HeadStyle:    nil,
	TailStyle:    nil,
}

func NewFixedStyleHandler() *FixedStyleHandler {
	rtn := &FixedStyleHandler{
		BaseHandler{Name: "FixedStyleHandler"},
		FixedStyleInPi.CalcLen(),
		&FixedStyleInPi,
	}
	return rtn
}

func (h *FixedStyleHandler) Judge(inbytes []byte) (ok Judgement, pkt Packet) {
	if len(inbytes) < h.MinLength {
		return NEEDMORE, pkt
	}

	pkt.PacketInfo = h.InPi
	headFromInbytes := inbytes[0:len(h.InPi.HeadStyle)]
	if h.InPi.HasFixedHead && bytes.Compare(h.InPi.HeadStyle, headFromInbytes) != 0 {
		//不符合协议头
		return REJECT, pkt
	}

	tailFromInbytes := inbytes[pkt.TailStartAt():]

	if h.InPi.HasFixedTail && bytes.Compare(h.InPi.TailStyle, tailFromInbytes) != 0 {
		//不符合协议尾
		return REJECT, pkt
	}

	pktlen := FixedStyleInPi.CalcLen()
	pkt.Content = make([]byte, pktlen)

	copy(pkt.Content, inbytes[:pktlen])
	//FIXME should check copy num equals pktlen?

	//计算校验和
	if h.InPi.HasCheckSum {
		inByteCheckSum := inbytes[pkt.CheckSumAt() : pkt.CheckSumAt()+len(pkt.CheckSumStyle)]
		pktCalcChecksum, err := pkt.CalcCheckSum()
		if err != nil {
			log.Debugf("Erroe occurd at calc check sum:%v", err)
			return REJECT, pkt
		}
		if bytes.Compare(inByteCheckSum, pktCalcChecksum) != 0 {
			return REJECT, pkt
		}
	}

	return ACCEPT, pkt
}
