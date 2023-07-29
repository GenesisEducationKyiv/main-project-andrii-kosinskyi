package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	DebugLevel = iota + 1
	InfoLevel
	ErrorLevel
)

type DefaultLog struct {
}

type LogMessage struct {
	Time     time.Time `json:"time"`
	LogLevel string    `json:"level"`
	Message  string    `json:"msg"`
}

var levelMap map[int]string

func NewLog() *DefaultLog {
	levelMap = map[int]string{
		1: "DEBUG",
		2: "INFO",
		3: "ERROR",
	}

	return &DefaultLog{}
}

func print(level int, msg string) {
	bytes, err := json.Marshal(&LogMessage{
		Time:     time.Now(),
		LogLevel: levelMap[level],
		Message:  msg,
	})
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
	}

	if level == ErrorLevel {
		fmt.Fprint(os.Stderr, string(bytes)+"\n")
	} else {
		fmt.Fprint(os.Stdout, string(bytes)+"\n")
	}
}

func (that *DefaultLog) Debug(ctx context.Context, msg string) {
	print(DebugLevel, msg)
}

func (that *DefaultLog) Info(ctx context.Context, msg string) {
	print(InfoLevel, msg)
}

func (that *DefaultLog) Error(ctx context.Context, msg string) {
	print(ErrorLevel, msg)
}
