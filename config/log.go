package config

import(
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

func Logger() *log.Logger {
	logger := log.New()

	logFile := LogFile()
	logger.Out = logFile


	return logger
}

func LogFile() io.Writer {
	day := time.Now().Format("2006-01-02")
	path := fmt.Sprintf("%s/%s", "dota2_logs", day)

	err := os.MkdirAll(path, 0755)

	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	logPath := path+"/d2.log"
	logFile, err := os.OpenFile(logPath, os.O_RDWR | os.O_APPEND, os.ModeAppend)


	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	return logFile
}