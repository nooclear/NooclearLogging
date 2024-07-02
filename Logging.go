package NooclearLogging

import (
	"fmt"
	"os"
	"time"
)

type Log struct {
	Timestamp string
	Category  string
	Msg       string
}

var TimestampFormat = "2006-01-02 15:04:05"
var LogDir = "./logs"
var LogDb = "logs.db"

func InitLogSystem() error {
	if err := os.MkdirAll(LogDir, os.ModePerm); err != nil {
		return err
	}
	if err := initDB(LogDir + "/" + LogDb); err != nil {
		return err
	}
	return nil
}

func formatLog(log Log) string {
	msg := fmt.Sprintf("%s %s %s", log.Timestamp, log.Category, log.Msg)
	return msg
}

func (l Log) Info(msg string) {
	l.Category = "INFO"
	logger(l, msg)
}
func (l Log) Warn(msg string) {
	l.Category = "WARN"
	logger(l, msg)
}
func (l Log) Error(err error) {
	l.Category = "ERROR"
	logger(l, err.Error())
	panic(err)
}
func (l Log) Success(msg string) {
	l.Category = "SUCCESS"
	logger(l, msg)
}
func (l Log) User(msg string) {
	l.Category = "USER"
	logger(l, msg)
}

func logger(l Log, msg string) {
	l.Timestamp = time.Now().Format(TimestampFormat)
	l.Msg = msg
	if l.Category != "USER" {
		fmt.Println(formatLog(l))
	}
	_, err := addLog(&l)
	if err != nil {
		panic(err)
	}
}
