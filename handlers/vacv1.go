package handlers

import (
    "fmt"
    "github.com/mosliu/gowalkwindow/mq"
)

type VacV1Handler struct {
    FixedStyleHandler
}

//51->pc
//AA AA X1 X2 X3 X4 X5 X6 X7 X8 X9 SM BB BB

//in: 51 -> pc
var VacV1InPi = PacketInfo{
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
var VacV1OutPi = PacketInfo{
    BodyLength:   10,
    HasFixedHead: false,
    HasFixedTail: false,
    HeadStyle:    nil,
    TailStyle:    nil,
}

func NewVacV1Handler() *VacV1Handler {
    rtn := &VacV1Handler{
        FixedStyleHandler: FixedStyleHandler{
            BaseHandler: BaseHandler{
                Name: "VacV1Handler",
            },
            MinLength: VacV1InPi.CalcLen(),
            InPi:      &VacV1InPi,
        },
    }
    return rtn
}

func (h *VacV1Handler) GetName() string {
    return "VacV1Handler"
}

func (h *VacV1Handler) Handle(pkt Packet) {
    //解析数据
    log.Debugf("do handle vac-v1 in packet %+v", pkt)
    body := pkt.GetBody()
    if body[0] == 0xCC {
        mainv := body[1]
        subv := body[2]
        orderv := body[3]
        mq.MQ.UiMsgQuene <- mq.UIMSG{
            ToUi:  mq.VERSION_LINEEDIT_SETSTR,
            Msg:   fmt.Sprintf("VAC-V1 v%d.%d.%d", mainv, subv, orderv),
            Value: 0,
        }

    }

}
