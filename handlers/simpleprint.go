package handlers

type SimplePrintHandler struct {
	BaseHandler
}

func NewSimplePrintHandler() *SimplePrintHandler {
	rtn := &SimplePrintHandler{
		BaseHandler{Name: "SimplePrintHandler"},
	}
	return rtn
}

var SimplePrintPi = PacketInfo{
	BodyLength:   10,
	HasFixedHead: false,
	HasFixedTail: false,
	HeadStyle:    nil,
	TailStyle:    nil,
}

//func (h *SimplePrintHandler) GetName() string {
//    return "SimplePrintHandler"
//}

func (h *SimplePrintHandler) Handle(pkt Packet) {
	log.Debugf("do handle packet %+v", pkt)
}
func (h *SimplePrintHandler) Judge(inbytes []byte) (ok Judgement, pkt Packet) {
	if len(inbytes) > 50 {
		//if len(inbytes) < 5 {
		ok = REJECT
		return
	}
	if len(inbytes) > 30 {
		ok = ACCEPT
		pkt.PacketInfo = &SimplePrintPi
		pkt.Content = make([]byte, SimplePrintPi.CalcLen())
	} else {
		ok = NEEDMORE
	}
	return
}
