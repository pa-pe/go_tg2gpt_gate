package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Log = logrus.New()

type Config struct {
	LogLevel      string
	Env           string
	AccessLogFile string
}

func (c Config) getLogLevel() logrus.Level {
	lvl, _ := logrus.ParseLevel(c.LogLevel)
	return lvl
}

func Init(c Config) {
	Log.SetLevel(c.getLogLevel())
	Log.SetOutput(os.Stdout) // setting up output to console (default)
	Log.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   false,
		TimestampFormat: time.StampMicro,
	})

	// Init access log
	initAccessLog(c.AccessLogFile, c.Env)
}

func LaunchLog(s string) {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Print(s)
}
