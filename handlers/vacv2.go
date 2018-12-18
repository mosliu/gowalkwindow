package handlers

import (
    "fmt"
    "github.com/mosliu/gowalkwindow/mq"
)

type VacV2Handler struct {
    FixedStyleHandler
}

//51->pc
//AA AA X1 X2 X3 X4 X5 X6 X7 X8 X9 XA XB XC XD SM BB BB (18 bytes)

//in: 51 -> pc
var VacV2InPi = PacketInfo{
    BodyLength:    9, //不含头尾合Sum
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
var VacV2OutPi = PacketInfo{
    BodyLength:   4,
    HasFixedHead: false,
    HasFixedTail: false,
    HeadStyle:    nil,
    TailStyle:    nil,
}

func NewVacV2Handler() *VacV2Handler {
    rtn := &VacV2Handler{
        FixedStyleHandler: FixedStyleHandler{
            BaseHandler: BaseHandler{
                Name: "VacV2Handler",
            },
            MinLength: VacV2InPi.CalcLen(),
            InPi:      &VacV2InPi,
        },
    }
    return rtn
}

func (h *VacV2Handler) GetName() string {
    return "VacV2Handler"
}

func (h *VacV2Handler) Handle(pkt Packet) {
    //解析数据
    log.Debugf("do handle vac-V2 in packet %+v", pkt)
    body := pkt.GetBody()
    if body[0] == 0xCC {
        mainv := body[1]
        subv := body[2]
        orderv := body[3]
        mq.MQ.UiMsgQuene <- mq.UIMSG{
            ToUi:  mq.VERSION_LINEEDIT_SETSTR,
            Msg:   fmt.Sprintf("VAC-V2 v%d.%d.%d", mainv, subv, orderv),
            Value: 0,
        }

    }

}
