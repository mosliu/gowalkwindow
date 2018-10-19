package ui

import (
	"github.com/mosliu/gowalkwindow/logs"
	"github.com/sirupsen/logrus"
)

var log = logs.Log.WithFields(logrus.Fields{
	"pkg": "ui",
})
var logf = logs.LogF.WithFields(logrus.Fields{
	"pkg": "ui",
})
