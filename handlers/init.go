package handlers

import (
	"github.com/mosliu/gowalkwindow/logs"
	"github.com/sirupsen/logrus"
)

var log = logs.Log.WithFields(logrus.Fields{
	"pkg": "handlers",
})

var logf = logs.Log.WithFields(logrus.Fields{
	"pkg": "handlers",
})
