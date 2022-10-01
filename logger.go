package dailyLogger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type StructLogger struct {
	logger         *log.Logger
	writer         io.Writer
	logsDir        string
	prefix         string
	loggerFilePath string
	loggerFile     *os.File
}

func NewLogger(logsDir, prefix string) *log.Logger {
	return new(StructLogger).init(logsDir, prefix)

	//or a := &StructLogger{}
	//return  a.init("loggs", "test")

}

func (structLogger *StructLogger) getLogger() *log.Logger {
	return structLogger.logger
}
func (structLogger *StructLogger) changeDayFile() {
	go func() {
		for {
			intervalTime := 24*60*60 - time.Now().Hour()*60*60 - time.Now().Minute()*60 - time.Now().Second()
			<-time.NewTimer(time.Duration(intervalTime) * time.Second).C
			var err error
			//loggerFilePath = logsDir + string(os.PathSeparator) + prefix + time.Now().Format("20060102") + ".log"
			loggerFilePath := filepath.Join(structLogger.logsDir, structLogger.prefix+time.Now().Format("20060102")+".log")

			if structLogger.loggerFilePath != loggerFilePath {
				structLogger.loggerFilePath = loggerFilePath
				//fmt.Println(loggerFilePath + "  " + loggerFile.Name())
				structLogger.loggerFile.Close()
				structLogger.loggerFile, err = os.OpenFile(structLogger.loggerFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
				if err != nil {
					fmt.Println(err)
				}
				//logger.SetOutput(loggerFile)
				//logger.SetOutput(io.MultiWriter(os.Stdout, loggerFile))
				if os.Getenv("debug") != "" {
					structLogger.logger.SetOutput(io.MultiWriter(os.Stdout, structLogger.loggerFile))
				} else {
					structLogger.logger.SetOutput(structLogger.loggerFile)
					//logger.SetOutput(os.Stdout)

				}
				//defer  loggerFile.Close()
			}
		}
	}()

}

func (structLogger *StructLogger) init(logsDir, prefix string) *log.Logger {
	var err error
	logsDir = filepath.Clean(logsDir)
	_, err = os.Stat(logsDir)
	if err != nil {
		err = os.Mkdir(logsDir, os.ModePerm)
		if err != nil {
			panic(err.Error())
		}
	}
	structLogger.prefix = prefix
	structLogger.logsDir = logsDir

	structLogger.loggerFilePath = filepath.Join(structLogger.logsDir, prefix+time.Now().Format("20060102")+".log")
	structLogger.loggerFile, err = os.OpenFile(structLogger.loggerFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	//fmt.Printf("path="+loggerFilePath +" \r\n")
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(os.Stderr, err)
		//fmt.Fprintf(os.Stderr, "%s", err)
	}

	if os.Getenv("debug") != "" {
		structLogger.logger = log.New(io.MultiWriter(os.Stdout, structLogger.loggerFile), "", log.LstdFlags)
	} else {
		structLogger.logger = log.New(structLogger.loggerFile, "", log.LstdFlags)
		//logger = log.New(os.Stdout, "", log.LstdFlags)
	}
	structLogger.changeDayFile()
	return structLogger.logger

}

func NewLogger_old(logsDir, prefix string) *log.Logger {
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
		fmt.Fprintln(os.Stderr, err)
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
