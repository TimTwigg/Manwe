package logger

import (
	"fmt"
	"os"
	"time"
)

func getLogFilePath() string {
	timeStamp := time.Now().Format("2006-01-02")
	return "logs/" + timeStamp + ".log"
}

func AppendToLog(message ...any) {
	messageStr := ""
	for _, m := range message {
		messageStr += fmt.Sprintf("%v ", m)
	}
	messageStr = messageStr[:len(messageStr)-1]

	timeStamp := time.Now().Format("2006-01-02 15:04:05")

	logFilePath := getLogFilePath()
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("[%s] %s\n", timeStamp, messageStr)); err != nil {
		fmt.Printf("Error writing to log file: %v\n", err)
	}
}
