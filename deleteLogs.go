package dailyLogger



import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func NewDelete(logsDir string, days int, logger *log.Logger) {
	logsDir = filepath.Base(logsDir)
	timeLong := time.Hour * 24 * time.Duration(days)
	delLogsDir(logsDir, timeLong, logger)
	go func(dir string, timeLong time.Duration) {
		for {
			intervalTime := 24*60*60 - time.Now().Hour()*60*60 - time.Now().Minute()*60 - time.Now().Second()
			<-time.NewTimer(time.Duration(intervalTime) * time.Second).C
			delLogsDir(dir, timeLong, logger)
		}
	}(logsDir, timeLong)
}
func delLogsDir(logsDir string, timeLong time.Duration, logger *log.Logger) {

	//logger.Printf("dir=%s\n", dir)
	dirList, e := ioutil.ReadDir(logsDir)
	if e != nil {
		return
	}
	for _, fileInfo := range dirList {
		fileName := fileInfo.Name()
		filePath := filepath.Join(logsDir, fileName)

		//logger.Printf("%s\r\n", filePath)
		if fileInfo.IsDir() {
			delLogsDir(filePath, timeLong, logger)
		} else {
			//logger.Printf("文件-> %s\r\n", filePath)
			//strings.LastIndex(fileName, ".log") == len(fileName)-len(".log")
			if strings.HasSuffix(fileName,".log") {
				if fileInfo.ModTime().Before(time.Now().Add(-timeLong)) {
					os.Remove(filePath)
					if logger !=nil {
						logger.Printf("delete old logs file %s\n", filePath)
					}

				}
			}
		}
	}
}
