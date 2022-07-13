package log

import (
	"path/filepath"
	"shapeservice/configs"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"

	"os"
)

var log *logrus.Logger

func GetLogger() *logrus.Logger {
	return log
}

func InitLogger() error {

	configLog := configs.GetConfig().LogConfig
	log = logrus.New()
	switch configLog.Format {
	case "json":
		log.SetFormatter(new(logrus.JSONFormatter))
	default:
		log.SetFormatter(new(logrus.TextFormatter))
	}
	if configLog.Output != "" {
		switch configLog.Output {
		case "stdout":
			log.SetOutput(os.Stdout)
		case "stderr":
			log.SetOutput(os.Stderr)
		case "file":
			wd, _ := os.Getwd()
			dir := filepath.Join(wd, configLog.Path)
			_ = os.Mkdir(dir, 0777)
			linkName := filepath.Join(dir, "shapeservice")
			write, err := rotatelogs.New(
				filepath.Join(dir, "shapeservice-%Y-%m-%d.log"),
				rotatelogs.WithMaxAge(time.Duration(24*configLog.Expired)*time.Hour),
				rotatelogs.WithLinkName(linkName),
				rotatelogs.WithRotationTime(24*time.Hour),
			)
			if err != nil {
				return err
			}
			log.SetOutput(write)
		}
	}
	return nil
}
