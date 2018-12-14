package ui

import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
)

func VACV2ValvesControlDialog(owner walk.Form) (int, error) {
    var v2dlg *walk.Dialog
    var acceptPB, cancelPB *walk.PushButton

    return Dialog{
        AssignTo:      &v2dlg,
        Title:         "V2 valves control",
        DefaultButton: &acceptPB,
        CancelButton:  &cancelPB,
        MinSize:       Size{300, 300},
        Layout:        VBox{},
        Children: []Widget{
            GroupBox{
                Title:  "Valves",
                Layout: HBox{Spacing: 2},
                Children: []Widget{
                    PushButton{
                        Text:      "Reset Rows",
                    },PushButton{
                        Text:      "Reset Rows",
                    },PushButton{
                        Text:      "Reset Rows",
                    },
                    PushButton{
                        Text:      "Reset Rows",
                    },PushButton{
                        Text:      "Reset Rows",
                    },PushButton{
                        Text:      "Reset Rows",
                    },
                    PushButton{
                        Text:      "Reset Rows",
                    },PushButton{
                        Text:      "Reset Rows",
                    },PushButton{
                        Text:      "Reset Rows",
                    },
                    PushButton{
                        Text:      "Reset Rows",
                    },PushButton{
                        Text:      "Reset Rows",
                    },PushButton{
                        Text:      "Reset Rows",
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
                            v2dlg.Accept()
                        },
                    },
                    PushButton{
                        AssignTo: &cancelPB,
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
