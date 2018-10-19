package handlers

type Judgement int

const (
	REJECT   Judgement = iota //拒绝
	ACCEPT                    //接受
	NEEDMORE                  //需要更多字节判断
)

// The Handler interfaces defines a method to receive a frame.
type Handler interface {
	// 处理器名字
	GetName() string
	// 可以处理一个包
	Handle(pkt Packet)
	// 可以检查是否可处理一个包
	Judge(inbytes []byte) (ok Judgement, pkt Packet)
}

type BaseHandler struct {
	Name string
}

func (h *BaseHandler) GetName() string {
	return h.Name
}

func (h BaseHandler) Handle(pkt Packet) {
	log.Debug("BaseHandler handle packet")
	//do nothing
}

func (h *BaseHandler) Judge(inbytes []byte) (ok Judgement, pkt Packet) {
	return NEEDMORE, pkt
}

// HandlerFunc defines the function type to handle a frame.
type HandlerFunc func(pkt Packet)

// HandlerFunc defines the function type to handle a frame.
type JudgeFunc func(inbytes []byte) (ok Judgement, pkt Packet)

type handler struct {
	name     string
	doHandle HandlerFunc
	doJudge  JudgeFunc
}

// NewHandler returns a new handler which calls fn when a frame is received.
func NewHandler(name string, doHandle HandlerFunc, doJudge JudgeFunc) Handler {
	return &handler{name, doHandle, doJudge}
}

func (h *handler) GetName() string {
	return h.name
}

func (h *handler) Handle(pkt Packet) {
	h.doHandle(pkt)
}
func (h *handler) Judge(inbytes []byte) (ok Judgement, pkt Packet) {
	return h.doJudge(inbytes)
}
