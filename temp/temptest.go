package main

import "github.com/lxn/walk"

func main() {
    runMainWindow()
}

type AMainWindow struct {
    *walk.MainWindow
    ui mainWindowUI
}

func runMainWindow() (int, error) {
    mw := new(AMainWindow)
    if err := mw.init(); err != nil {
        return 0, err
    }
    defer mw.Dispose()

    // TODO: Do further required setup, e.g. for event handling, here.

    mw.Show()

    return mw.Run(), nil
}



type mainWindowUI struct {
    centralwidget        *walk.Composite
    verticalLayoutWidget *walk.Composite
    lineEdit             *walk.LineEdit
    pushButton           *walk.PushButton
    tableView            *walk.TableView
}

func (w *AMainWindow) init() (err error) {
    if w.MainWindow, err = walk.NewMainWindow(); err != nil {
        return err
    }

    succeeded := false
    defer func() {
        if !succeeded {
            w.Dispose()
        }
    }()

    var font *walk.Font
    if font == nil {
        font = nil
    }

    w.SetName("MainWindow")
    l := walk.NewVBoxLayout()
    if err := l.SetMargins(walk.Margins{0, 0, 0, 0}); err != nil {
        return err
    }
    if err := w.SetLayout(l); err != nil {
        return err
    }
    if err := w.SetClientSize(walk.Size{800, 600}); err != nil {
        return err
    }
    if err := w.SetTitle(`MainWindow`); err != nil {
        return err
    }

    // centralwidget
    if w.ui.centralwidget, err = walk.NewComposite(w); err != nil {
        return err
    }
    w.ui.centralwidget.SetName("centralwidget")

    // verticalLayoutWidget
    if w.ui.verticalLayoutWidget, err = walk.NewComposite(w.ui.centralwidget); err != nil {
        return err
    }
    w.ui.verticalLayoutWidget.SetName("verticalLayoutWidget")
    if err := w.ui.verticalLayoutWidget.SetBounds(walk.Rectangle{70, 40, 251, 411}); err != nil {
        return err
    }
    verticalLayout := walk.NewVBoxLayout()
    if err := w.ui.verticalLayoutWidget.SetLayout(verticalLayout); err != nil {
        return err
    }
    if err := verticalLayout.SetMargins(walk.Margins{9, 9, 9, 9}); err != nil {
        return err
    }
    if err := verticalLayout.SetSpacing(6); err != nil {
        return err
    }

    // anonymous spacer
    if _, err := walk.NewVSpacer(w.ui.verticalLayoutWidget); err != nil {
        return err
    }

    // lineEdit
    if w.ui.lineEdit, err = walk.NewLineEdit(w.ui.verticalLayoutWidget); err != nil {
        return err
    }
    w.ui.lineEdit.SetName("lineEdit")
    w.ui.lineEdit.SetReadOnly(true)

    // anonymous spacer
    if _, err := walk.NewVSpacer(w.ui.verticalLayoutWidget); err != nil {
        return err
    }

    // pushButton
    if w.ui.pushButton, err = walk.NewPushButton(w.ui.verticalLayoutWidget); err != nil {
        return err
    }
    w.ui.pushButton.SetName("pushButton")
    if err := w.ui.pushButton.SetText(`PushButton`); err != nil {
        return err
    }

    // anonymous spacer
    if _, err := walk.NewVSpacer(w.ui.verticalLayoutWidget); err != nil {
        return err
    }

    // tableView
    if w.ui.tableView, err = walk.NewTableView(w.ui.centralwidget); err != nil {
        return err
    }
    w.ui.tableView.SetName("tableView")
    if err := w.ui.tableView.SetBounds(walk.Rectangle{360, 200, 256, 192}); err != nil {
        return err
    }

    // toolBar

    // Tab order

    succeeded = true

    return nil
}
