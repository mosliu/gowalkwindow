package ui

import (
    "encoding/binary"
    "errors"
    "fmt"
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
    "strconv"
)

type V2Dialog struct {
    *walk.Dialog
    outPutTextEdit     *walk.TextEdit
    acceptPB, cancelPB *walk.PushButton
    valveStatus        []bool
    valvesStatus       uint16
    pkt                []byte
}

var v2dlg = V2Dialog{
    valveStatus:[]bool{false,true,false,true,false,true,true,false,true,false,true,false,false,false,false,false},
    pkt: []byte{0xAA, 0xAA, 05, 00, 00, 00, 05, 0xBB, 0xBB},
}

func VACV2ValvesControlDialog(owner walk.Form) (int, error) {

    return Dialog{
        AssignTo:      &v2dlg.Dialog,
        Title:         "V2 valves control",
        DefaultButton: &v2dlg.acceptPB,
        CancelButton:  &v2dlg.cancelPB,
        MinSize:       Size{600, 300},
        Layout:        VBox{},
        Children: []Widget{
            GroupBox{
                Title:  "Valves",
                Layout: HBox{Spacing: 2},
                Children: []Widget{
                    PushButton{
                        Text: "1",
                        OnClicked: func() {
                            valveClick(1)
                        },
                    }, PushButton{
                        Text: "2",
                        OnClicked: func() {
                            valveClick(2)
                        },
                    }, PushButton{
                        Text: "3",
                        OnClicked: func() {
                            valveClick(3)
                        },
                    }, PushButton{
                        Text: "4",
                        OnClicked: func() {
                            valveClick(4)
                        },
                    }, PushButton{
                        Text: "5",
                        OnClicked: func() {
                            valveClick(5)
                        },
                    }, PushButton{
                        Text: "6",
                        OnClicked: func() {
                            valveClick(6)
                        },
                    }, PushButton{
                        Text: "7",
                        OnClicked: func() {
                            valveClick(7)
                        },
                    }, PushButton{
                        Text: "8",
                        OnClicked: func() {
                            valveClick(8)
                        },
                    }, PushButton{
                        Text: "9",
                        OnClicked: func() {
                            valveClick(9)
                        },
                    }, PushButton{
                        Text: "10",
                        OnClicked: func() {
                            valveClick(10)
                        },
                    },
                },
            },
            GroupBox{
                Title:  "Status",
                Layout: HBox{Spacing: 2},
                Children: []Widget{
                    TextEdit{
                        Font:     Font{Family: "consolas",},
                        AssignTo: &v2dlg.outPutTextEdit,
                        Text:     "",
                    },
                },
            },
            Composite{
                Layout: HBox{},
                Children: []Widget{
                    HSpacer{},
                    PushButton{
                        AssignTo: &v2dlg.acceptPB,
                        Text:     "OK",
                        OnClicked: func() {
                            v2dlg.Accept()
                        },
                    },
                    PushButton{
                        AssignTo: &v2dlg.cancelPB,
                        Text:     "Cancel",
                        OnClicked: func() {
                            v2dlg.Cancel()
                        },
                    },
                },
            },
        },
    }.Run(owner)
}

//打开串口按钮
func valveClick(no int) error {
    if no <= 0 || no > 10 {
        msg := fmt.Sprintf("Error Number: %n input", no)
        return errors.New(msg)
        log.Warnln(msg)
    }
    switch no {
    case 1, 2, 3:
        if v2dlg.valveStatus[(no-1)*2] == true {
            v2dlg.valveStatus[(no-1)*2] = false
            v2dlg.valveStatus[(no-1)*2+1] = true
        } else {
            v2dlg.valveStatus[(no-1)*2] = true
            v2dlg.valveStatus[(no-1)*2+1] = false
        }
    case 4, 5, 6:
        //4B 判断
        if v2dlg.valveStatus[(no-1)*2+1] == true {
            v2dlg.valveStatus[(no-1)*2+1] = false
            v2dlg.valveStatus[(no-1)*2] = true
        } else {
            v2dlg.valveStatus[(no-1)*2+1] = true
            v2dlg.valveStatus[(no-1)*2] = false
        }
    case 7, 8, 9, 10:
        if v2dlg.valveStatus[no-1+6] == true {
            v2dlg.valveStatus[no-1+6] = false
        } else {
            v2dlg.valveStatus[no-1+6] = true
        }
    }
    //calcStatus(&v2dlg)
    v2dlg.calcStatus()
    v2dlg.showCurrentValves()
    log.Debugln(v2dlg.valvesStatus)
    //start("COM1")
    return nil
}


//在下面的输出框展示当前阀状态
func (dlg *V2Dialog) showCurrentValves() {
    var s string
    s += ""
    for i := 1; i < 11; i++ {
        s += " 阀" + strconv.Itoa(i)
    }
    s += "\r\n"
    for i := 0; i < 6; i += 2 {
        if dlg.valveStatus[i] {
            s += " 开 "
        } else {
            s += " 关 "
        }
    }
    for i := 7; i < 12; i += 2 {
        if dlg.valveStatus[i] {
            s += " 开 "
        } else {
            s += " 关 "
        }
    }
    for i := 12; i < 16; i++ {
        if dlg.valveStatus[i] {
            s += " 开 "
        } else {
            s += " 关 "
        }
    }

    binary.BigEndian.PutUint16(dlg.pkt[3:5], dlg.valvesStatus)
    var sum byte
    for i := 2; i < 6; i++ {
        sum += dlg.pkt[i]
    }
    dlg.pkt[6] = sum
    hexStr := fmt.Sprintf("%X", dlg.pkt)
    s += "\r\n" + hexStr
    dlg.outPutTextEdit.SetText(s)
}

//通过阀状态16位的状态表，计算2字节的用于发送的16进制状态
func (dlg *V2Dialog) calcStatus() {
    var tmp uint16
    for i := 15; i > -1; i-- {
        if dlg.valveStatus[i] {
            //dlg.valvesStatus = dlg.valvesStatus | 1
            tmp = tmp | 1
        }
        if i != 0 {
            //dlg.valvesStatus = dlg.valvesStatus << 1
            tmp = tmp << 1
        }
    }
    dlg.valvesStatus = tmp
}
