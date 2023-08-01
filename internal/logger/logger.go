package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"bitcoin_checker_api/internal/pkg/broker"
)

const (
	DebugLevel = iota + 1
	InfoLevel
	ErrorLevel
)

type DefaultLog struct {
	brokerSrv broker.Service
}

type LogMessage struct {
	Time     time.Time `json:"time"`
	LogLevel string    `json:"level"`
	Message  string    `json:"msg"`
}

var levelMap map[int]string //nolint:gochecknoglobals

func NewLog(brokerSrv broker.Service) *DefaultLog {
	levelMap = map[int]string{
		1: "DEBUG",
		2: "INFO",
		3: "ERROR",
	}

	return &DefaultLog{brokerSrv: brokerSrv}
}

func (that *DefaultLog) printLog(level int, msg string) {
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
		if that.brokerSrv != nil { //nolint:typecheck
			err = that.brokerSrv.SendErr(bytes)
			if err != nil {
				fmt.Fprint(os.Stderr, err.Error()+"\n")
			}
		}
	} else {
		fmt.Fprint(os.Stdout, string(bytes)+"\n")
	}

	if that.brokerSrv != nil { //nolint:typecheck
		err = that.brokerSrv.Send(bytes)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error()+"\n")
		}
	}
}

func (that *DefaultLog) Debug(_ context.Context, msg string) {
	that.printLog(DebugLevel, msg)
}

func (that *DefaultLog) Info(_ context.Context, msg string) {
	that.printLog(InfoLevel, msg)
}

func (that *DefaultLog) Error(_ context.Context, msg string) {
	that.printLog(ErrorLevel, msg)
}
