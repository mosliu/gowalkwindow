package main

import (
    "fmt"
    "github.com/lxn/walk"
    "log"
)
import . "github.com/lxn/walk/declarative"

type AppWindow struct {
    *walk.MainWindow
    recentMenu *walk.Menu
    outPutTextEdit *walk.TextEdit
    //status bar items
}

func showWindow() {
    appw := new(AppWindow)
    var btn *walk.PushButton
    var openAction, showAboutBoxAction *walk.Action
    //var recentMenu *walk.Menu

    icon1, err := walk.NewIconFromFile("./assets/imgs/check.ico")
    if err != nil {
        log.Fatal(err)
    }
    icon2, err := walk.NewIconFromFile("./assets/imgs/stop.ico")
    if err != nil {
        log.Fatal(err)
    }
    var sbi *walk.StatusBarItem
    err = MainWindow{
        AssignTo: &appw.MainWindow,
        Name:     "haliluya",
        Title:    "兰光版本号下位机读取工具",
        MinSize:  Size{800, 600},
        Layout: VBox{
            MarginsZero: true,
        },
        MenuItems: []MenuItem{
            Menu{
                Text: "&File",
                Items: []MenuItem{
                    Action{
                        AssignTo:    &openAction,
                        Text:        "&Open",
                        Image:       "./assets/imgs/open.png",
                        Enabled:     Bind("enabledCB.Checked"),
                        Visible:     Bind("!openHiddenCB.Checked"),
                        Shortcut:    Shortcut{walk.ModControl, walk.KeyO},
                        OnTriggered: appw.openAction_Triggered,
                    },
                    Menu{
                        AssignTo: &appw.recentMenu,
                        Text:     "Recent",
                    },
                    Separator{},
                    Action{
                        Text:        "E&xit",
                        OnTriggered: func() { appw.Close() },
                    },
                },
            },
            Menu{
                Text: "&Help",
                Items: []MenuItem{
                    Action{
                        AssignTo:    &showAboutBoxAction,
                        Text:        "About",
                        OnTriggered: appw.showAboutBoxAction_Triggered,
                    },
                },
            },
        },
        StatusBarItems: []StatusBarItem{
            StatusBarItem{
                AssignTo: &sbi,
                Icon:     icon1,
                Text:     "click",
                Width:    80,
                OnClicked: func() {
                    if sbi.Text() == "click" {
                        sbi.SetText("again")
                        sbi.SetIcon(icon2)
                    } else {
                        sbi.SetText("click")
                        sbi.SetIcon(icon1)
                    }
                },
            },
            StatusBarItem{
                Text:        "left",
                ToolTipText: "no tooltip for me",
            },
            StatusBarItem{
                Text: "\tcenter",
            },
            StatusBarItem{
                Text: "\t\tright",
            },
            StatusBarItem{
                Icon:        icon1,
                ToolTipText: "An icon with a tooltip",
            },
        },
        Children: []Widget{
            VSpacer{},
            //主区域
            Composite{
                Layout: HBox{},
                Children: []Widget{
                    Composite{
                        MinSize: Size{Width: 250,},
                        Layout:  Grid{Columns: 2},
                        Children: []Widget{
                            VSpacer{ColumnSpan: 2},
                            PushButton{
                                Text: "Open",
                            },
                            PushButton{
                                Text: "Send",
                            },
                            VSpacer{ColumnSpan: 2},
                            LineEdit{
                                Font:       Font{PointSize: 20},
                                Text:       "AKAKAA",
                                ColumnSpan: 2,
                                ReadOnly:   true,
                            },
                            VSpacer{ColumnSpan: 2},
                        },
                    },
                    HSplitter{},
                    Composite{
                        MinSize: Size{Width: 550,},
                        Layout:  VBox{},
                        Children: []Widget{
                            TextEdit{
                                AssignTo: &appw.outPutTextEdit,
                                Text: "aha",
                            },
                            PushButton{
                                Text:     "push me",
                                AssignTo: &btn,
                                OnClicked: func() {
                                    btnClicked(btn)
                                },
                            },
                            Label{
                                Text: "点击按钮后，可拖动标题栏测试界面状态。",
                            },
                        },
                    },
                },
            },

            VSpacer{},
        },
    }.Create()
    if err != nil {
        log.Fatal(err)
    }
    appw.addRecentFileActions("Foo", "Bar", "Baz")
    appw.Run()
}

func btnClicked(btn *walk.PushButton) {
    btn.SetText("哈哈")
}

func (appw *AppWindow) openAction_Triggered() {
    //walk.MsgBox(appw, "Open", "Pretend to open a file...", walk.MsgBoxIconInformation)
    currentSerialPort := SerialPort{
        "COM1",
        9600,
        ParityNone,
        Stop1,
    }
    if cmd, err := RunSerialPortSetDialog(appw, &currentSerialPort); err != nil {
        log.Print(err)
    } else if cmd == walk.DlgCmdOK {
        appw.outPutTextEdit.SetText(fmt.Sprintf("%+v", &currentSerialPort))
    }
}

func (appw *AppWindow) showAboutBoxAction_Triggered() {
    walk.MsgBox(appw, "About", "Walk Actions Example", walk.MsgBoxIconInformation)
}

func (appw *AppWindow) addRecentFileActions(texts ...string) {
    for _, text := range texts {
        a := walk.NewAction()
        a.SetText(text)
        a.Triggered().Attach(appw.openAction_Triggered)
        appw.recentMenu.Actions().Add(a)
    }
}
