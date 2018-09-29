package main

import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
    "log"
)

type SerialPortMainWindow struct {
    *walk.MainWindow
}

var isSpecialMode = walk.NewMutableCondition()

func showWindow2() {
    MustRegisterCondition("isSpecialMode", isSpecialMode)

    mw := new(SerialPortMainWindow)

    var openAction, showAboutBoxAction *walk.Action
    var recentMenu *walk.Menu
    var toggleSpecialModePB *walk.PushButton

    if err := (MainWindow{
        AssignTo: &mw.MainWindow,
        Title:    "Walk Actions Example",
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
                        OnTriggered: mw.openAction_Triggered,
                    },
                    Menu{
                        AssignTo: &recentMenu,
                        Text:     "Recent",
                    },
                    Separator{},
                    Action{
                        Text:        "E&xit",
                        OnTriggered: func() { mw.Close() },
                    },
                },
            },
            Menu{
                Text: "&Help",
                Items: []MenuItem{
                    Action{
                        AssignTo:    &showAboutBoxAction,
                        Text:        "About",
                        OnTriggered: mw.showAboutBoxAction_Triggered,
                    },
                },
            },
        },
        ToolBar: ToolBar{
            ButtonStyle: ToolBarButtonImageBeforeText,
            Items: []MenuItem{
                ActionRef{&openAction},
                Menu{
                    Text:  "New A",
                    Image: "./assets/imgs/document-new.png",
                    Items: []MenuItem{
                        Action{
                            Text:        "A",
                            OnTriggered: mw.newAction_Triggered,
                        },
                        Action{
                            Text:        "B",
                            OnTriggered: mw.newAction_Triggered,
                        },
                        Action{
                            Text:        "C",
                            OnTriggered: mw.newAction_Triggered,
                        },
                    },
                    OnTriggered: mw.newAction_Triggered,
                },
                Separator{},
                Menu{
                    Text:  "View",
                    Image: "./assets/imgs/document-properties.png",
                    Items: []MenuItem{
                        Action{
                            Text:        "X",
                            OnTriggered: mw.changeViewAction_Triggered,
                        },
                        Action{
                            Text:        "Y",
                            OnTriggered: mw.changeViewAction_Triggered,
                        },
                        Action{
                            Text:        "Z",
                            OnTriggered: mw.changeViewAction_Triggered,
                        },
                    },
                },
                Separator{},
                Action{
                    Text:        "Special",
                    Image:       "./assets/imgs/system-shutdown.png",
                    Enabled:     Bind("isSpecialMode && enabledCB.Checked"),
                    OnTriggered: mw.specialAction_Triggered,
                },
            },
        },
        ContextMenuItems: []MenuItem{
            ActionRef{&showAboutBoxAction},
        },
        MinSize: Size{300, 200},
        Layout:  VBox{},
        Children: []Widget{
            CheckBox{
                Name:    "enabledCB",
                Text:    "Open / Special Enabled",
                Checked: true,
            },
            CheckBox{
                Name:    "openHiddenCB",
                Text:    "Open Hidden",
                Checked: true,
            },
            PushButton{
                AssignTo: &toggleSpecialModePB,
                Text:     "Enable Special Mode",
                OnClicked: func() {
                    isSpecialMode.SetSatisfied(!isSpecialMode.Satisfied())

                    if isSpecialMode.Satisfied() {
                        toggleSpecialModePB.SetText("Disable Special Mode")
                    } else {
                        toggleSpecialModePB.SetText("Enable Special Mode")
                    }
                },
            },
        },
    }.Create()); err != nil {
        log.Fatal(err)
    }

    addRecentFileActions := func(texts ...string) {
        for _, text := range texts {
            a := walk.NewAction()
            a.SetText(text)
            a.Triggered().Attach(mw.openAction_Triggered)
            recentMenu.Actions().Add(a)
        }
    }

    addRecentFileActions("Foo", "Bar", "Baz")

    mw.Run()
}

func (mw *SerialPortMainWindow) openAction_Triggered() {
    walk.MsgBox(mw, "Open", "Pretend to open a file...", walk.MsgBoxIconInformation)
}

func (mw *SerialPortMainWindow) newAction_Triggered() {
    walk.MsgBox(mw, "New", "Newing something up... or not.", walk.MsgBoxIconInformation)
}

func (mw *SerialPortMainWindow) changeViewAction_Triggered() {
    walk.MsgBox(mw, "Change View", "By now you may have guessed it. Nothing changed.", walk.MsgBoxIconInformation)
}

func (mw *SerialPortMainWindow) showAboutBoxAction_Triggered() {
    walk.MsgBox(mw, "About", "Walk Actions Example", walk.MsgBoxIconInformation)
}

func (mw *SerialPortMainWindow) specialAction_Triggered() {
    walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
}
