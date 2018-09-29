package main

import (
    . "github.com/lxn/walk/declarative"
)

func main() {
    MainWindow{

        MinSize: Size{300, 200},
        Layout:  Grid{},
        Children: []Widget{
            Label{
                Row:    0,
                Column: 0,
                Text:   "I feel like someone is pulling me",
            },
            TextEdit{
                Row: 0,
                //RowSpan: 2,
                Column: 1,
            },
            HSpacer{
                Row:        1,
                Column:     0,
                ColumnSpan: 4,
            },
            Label{
                Row:    1,
                Column: 0,
                Text:   "a?10",
            },
            Label{
                Row:    2,
                Column: 0,
                Text:   "a?20",
            },
            Label{
                Row:    3,
                Column: 2,
                Text:   "a?32",
            },
            Label{
                Row:    4,
                Column: 1,
                Text:   "a?41",
            },
            Label{
                Row:    5,
                Column: 5,
                Text:   "a?55",
            },
        },
    }.Run()
}
