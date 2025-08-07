package logger

import (
	"log"
	"os"
	"time"
)

func GetBLogFilePath() string {
	timeStamp := time.Now().Format("2006-01-02")
	return "logs/" + timeStamp + ".b.log"
}

func AppendToBLog(message ...any) {
	filePath := GetBLogFilePath()
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Error("Error opening B log file:", err)
		return
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	logger.Println(message...)
}
