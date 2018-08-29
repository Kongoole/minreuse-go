package log

import (
	"os"
	"log"
	"time"
)

func Fatal(v ...interface{}) error {
	file, err := logFile()
	if err != nil {
		log.Fatal("not found log file")
	}
	defer file.Close()
	logger := log.New(file, "Fatal\t", log.LstdFlags)
	logger.Fatal(v)
	return nil
}

func Debug(v ...interface{}) error {
	file, err := logFile()
	if err != nil {
		log.Fatal("not found log file")
	}
	defer file.Close()
	logger := log.New(file, "Debug\t", log.LstdFlags)
	logger.Println(v)
	return nil
}

func logFile() (*os.File, error) {
	logPath := os.Getenv("LOG_PATH")
	_, err := os.Stat(logPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("cannot find log file dir which is '" + logPath + "'")
		}
	}
	fileName := logPath + "/" + time.Now().Format("2006-01-02")
	file, err := os.OpenFile(fileName, os.O_CREATE | os.O_APPEND|os.O_RDWR, os.ModeAppend|os.ModePerm)
	if err != nil {
		return nil, err
	}
	return file, nil
}