package log

import (
	"os"

	"github.com/evandroferreiras/gopher-city-bank/app/common/envvar"
	echoLog "github.com/labstack/gommon/log"
	echoLogrus "github.com/neko-neko/echo-logrus/v2/log"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func getFormatter() logrus.Formatter {
	if isLocalEnv := envvar.IsLocalEnv(); isLocalEnv {
		return &prefixed.TextFormatter{
			ForceFormatting: true,
		}
	}
	return &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
	}
}

// Init sets the default config of logrus
func Init() {
	logrus.SetFormatter(getFormatter())
	logrus.SetLevel(logrus.DebugLevel)
}

// EchoLogrus sets the default config of echo logrus
func EchoLogrus() *echoLogrus.MyLogger {
	echoLogrus.Logger().SetOutput(os.Stdout)
	echoLogrus.Logger().SetLevel(echoLog.DEBUG)
	echoLogrus.Logger().SetFormatter(getFormatter())
	return echoLogrus.Logger()
}
