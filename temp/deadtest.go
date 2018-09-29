package main
import (
    "log"
    "strconv"
    "time"
)
import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
)
var mw *walk.MainWindow
var acceptPB_1, acceptPB_2 *walk.PushButton
var timer, message *walk.LineEdit
func add() {
    t0 := time.Now()
    for i := 0; i < 100000; i++ {
        message.SetText(strconv.Itoa(i))
    }
    t1 := time.Now()
    timer.SetText(t1.Sub(t0).String())
}
func main() {
    if _, err := (MainWindow{
        AssignTo: &mw,
        Title:    "界面演示",
        MinSize:  Size{300, 200},
        Layout:   VBox{},
        Children: []Widget{
            Composite{
                Layout: Grid{Columns: 2},
                Children: []Widget{
                    Label{
                        Text: "运行耗时:",
                    },
                    LineEdit{
                        AssignTo: &timer,
                        ReadOnly: true,
                    },
                    Label{
                        Text: "中间结果:",
                    },
                    LineEdit{
                        AssignTo: &message,
                        ReadOnly: true,
                    },
                },
            },
            Composite{
                Layout: HBox{},
                Children: []Widget{
                    Label{
                        Text: "点击按钮后，可拖动标题栏测试界面状态。",
                    },
                },
            },
            Composite{
                Layout: HBox{},
                Children: []Widget{
                    HSpacer{},
                    PushButton{
                        AssignTo: &acceptPB_1,
                        Text:     "普通测试",
                        OnClicked: func() {
                            timer.SetText("")
                            add()
                        },
                    },
                    PushButton{
                        AssignTo: &acceptPB_2,
                        Text:     "防假死测试",
                        OnClicked: func() {
                            timer.SetText("")
                            go func() {
                                add()
                            }()
                        },
                    },
                },
            },
        },
    }.Run()); err != nil {
        log.Fatal(err)
    }
}