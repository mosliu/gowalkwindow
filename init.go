package main
import (
    "github.com/mosliu/gowalkwindow/logs"
    "github.com/sirupsen/logrus"
)

var log = logs.Log.WithFields(logrus.Fields{
    "pkg":"libs",
})