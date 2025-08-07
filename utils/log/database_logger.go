package logger

import (
	"context"
	"log"
	"time"

	tracelog "github.com/jackc/pgx/v5/tracelog"
)

var DATABASE_LOGGER_COUNT int = 0

func GetLogFilePath() string {
	timeStamp := time.Now().Format("2006-01-02")
	return "logs/" + timeStamp + ".log"
}

type DatabaseLogger struct {
	Logger *log.Logger
}

func (cl *DatabaseLogger) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]any) {
	DATABASE_LOGGER_COUNT++
	cl.Logger.Printf("[%s] %s - %+v\n", level, msg, data)
}
