// Author：Mosliu
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.


package ui

import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
    "strconv"
)

type StopBits byte
type Parity byte

const (
    Stop1     StopBits = 1
    Stop1Half StopBits = 15
    Stop2     StopBits = 2
)

const (
    ParityNone  Parity = 'N'
    ParityOdd   Parity = 'O'
    ParityEven  Parity = 'E'
    ParityMark  Parity = 'M' // parity bit is always 1
    ParitySpace Parity = 'S' // parity bit is always 0
)

type SerialPort struct {
    Name string
    Baud int
    // Parity is the bit to use and defaults to ParityNone (no parity bit).
    Parity Parity
    // Number of stop bits to use. Default is 1 (1 stop bit).
    StopBits StopBits
}

func RunSerialPortSetDialog(owner walk.Form, serialPort *SerialPort) (int, error) {
    var dlg *walk.Dialog
    var db *walk.DataBinder
    var acceptPB, cancelPB *walk.PushButton

    return Dialog{
        AssignTo:      &dlg,
        Title:         "SerialPort Setting",
        DefaultButton: &acceptPB,
        CancelButton:  &cancelPB,
        DataBinder: DataBinder{
            AssignTo:       &db,
            DataSource:     serialPort,
            ErrorPresenter: ToolTipErrorPresenter{},
        },
        MinSize: Size{300, 300},
        Layout:  VBox{},
        Children: []Widget{
            Composite{
                Layout: Grid{Columns: 2},
                Children: []Widget{
                    Label{
                        Text: "Comm No:",
                    },

                    ComboBox{
                        Value: Bind("Name", SelRequired{}),
                        Model: getComs(),
                    },

                    Label{
                        Text: "Baud:",
                    },
                    NumberEdit{
                        Value:    Bind("Baud", Range{600, 4608000}),
                        Decimals: 0,
                    },

                    RadioButtonGroupBox{
                        ColumnSpan: 2,
                        Title:      "Parity",
                        Layout:     HBox{},
                        DataMember: "Parity",
                        Buttons: []RadioButton{
                            {Text: "None", Value: ParityNone},
                            {Text: "Odd", Value: ParityOdd},
                            {Text: "Even", Value: ParityEven},
                            {Text: "Mark", Value: ParityMark},
                            {Text: "Space", Value: ParitySpace},
                        },
                    },


                    RadioButtonGroupBox{
                        ColumnSpan: 2,
                        Title:      "StopBits",
                        Layout:     HBox{},
                        DataMember: "StopBits",
                        Buttons: []RadioButton{
                            {Text: "1位", Value: Stop1},
                            {Text: "1.5位", Value: Stop1Half},
                            {Text: "2位", Value: Stop2},
                        },
                    },
                },
            },
            Composite{
                Layout: HBox{},
                Children: []Widget{
                    HSpacer{},
                    PushButton{
                        AssignTo: &acceptPB,
                        Text:     "OK",
                        OnClicked: func() {
                            if err := db.Submit(); err != nil {
                                log.Print(err)
                                return
                            }

                            dlg.Accept()
                        },
                    },
                    PushButton{
                        AssignTo:  &cancelPB,
                        Text:      "Cancel",
                        OnClicked: func() { dlg.Cancel() },
                    },
                },
            },
        },
    }.Run(owner)
}

//生成选择用的COM口

func getComs() (rtn []string) {
    comsize := 40
    rtn = make([]string, comsize)
    for i := 1; i < comsize+1; i++ {
        rtn[i-1] = "COM" + strconv.Itoa(i)
    }
    //return rtn
    return
}
