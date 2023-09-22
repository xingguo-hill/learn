package x

import (
	"testing"

	logrus "github.com/sirupsen/logrus"
)

func TestLogrus(t *testing.T) {
	// 设置日志等级
	logrus.Trace("Something very low level.")
	logrus.Debug("Useful debugging information.")
	logrus.Info("Something noteworthy happened!")
	logrus.Warn("You should probably take a look at this.")
	logrus.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logginglog.Fatal("Bye.")
	// Calls panic() after logginglog.Panic("I'm bailing.")
	logrus.SetLevel(logrus.DebugLevel)
}
