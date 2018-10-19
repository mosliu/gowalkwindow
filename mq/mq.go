//消息系统 消息传递
package mq

type UiReceiver int

const (
	VERSION_LINEEDIT_SETSTR UiReceiver = iota //版本 文本框
	OUTPUT_TEXTEDIT_APPEND                    //输出框
)

type UIMSG struct {
	ToUi  UiReceiver
	Msg   string
	Value int
}
type MessageQuene struct {
	UiMsgQuene chan UIMSG
}

var MQ MessageQuene

func InitMQ() {
	MQ.UiMsgQuene = make(chan UIMSG)
}
