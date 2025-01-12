package util

import "github.com/sirupsen/logrus"

var Log *logrus.Logger


func init() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	Log.SetLevel(logrus.InfoLevel)
}

func LogDefault(chargingStationID string, feature string) *logrus.Entry {
	return Log.WithFields(logrus.Fields{"client": chargingStationID, "message": feature})
}