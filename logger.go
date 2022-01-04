package  dailyLogger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func NewLogger(logsDir, prefix string) *log.Logger {
	// os.PathSeparator
	var err error
	var logger *log.Logger

	logsDir = filepath.Clean(logsDir)
	_, err = os.Stat(logsDir)
	if err != nil {
		err = os.Mkdir(logsDir, os.ModePerm)
		if err != nil {
			panic(err.Error())
		}
	}

	loggerFilePath := filepath.Join(logsDir, prefix+time.Now().Format("20060102")+".log")
	loggerFile, err := os.OpenFile(loggerFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	//fmt.Printf("path="+loggerFilePath +" \r\n")
	if err != nil {
		fmt.Println(err)
	}
	//defer  loggerFile.Close()
	//logger = log.New(loggerFile, "", log.LstdFlags)
	// 	if runtime.GOOS == "windows" {

	if os.Getenv("debug") != "" {
		logger = log.New(io.MultiWriter(os.Stdout, loggerFile), "", log.LstdFlags)
	} else {
		logger = log.New(loggerFile, "", log.LstdFlags)
		//logger = log.New(os.Stdout, "", log.LstdFlags)
	}

	go func(logsDir, prefix string, loggerFile *os.File) {
		for {
			intervalTime := 24*60*60 - time.Now().Hour()*60*60 - time.Now().Minute()*60 - time.Now().Second()
			<-time.NewTimer(time.Duration(intervalTime) * time.Second).C
			var err error
			//loggerFilePath = logsDir + string(os.PathSeparator) + prefix + time.Now().Format("20060102") + ".log"
			loggerFilePath := filepath.Join(logsDir, prefix+time.Now().Format("20060102")+".log")

			if loggerFile.Name() != loggerFilePath {

				//fmt.Println(loggerFilePath + "  " + loggerFile.Name())
				loggerFile.Close()
				loggerFile, err = os.OpenFile(loggerFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
				if err != nil {
					fmt.Println(err)
				}
				//logger.SetOutput(loggerFile)
				//logger.SetOutput(io.MultiWriter(os.Stdout, loggerFile))
				if os.Getenv("debug") != "" {
					logger.SetOutput(io.MultiWriter(os.Stdout, loggerFile))
				} else {
					logger.SetOutput(loggerFile)
					//logger.SetOutput(os.Stdout)

				}
				//defer  loggerFile.Close()
			}
		}

	}(logsDir, prefix, loggerFile)
	return logger
}
