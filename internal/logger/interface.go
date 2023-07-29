package logger

import "context"

type ILog interface {
	Debug(ctx context.Context, msg string)
	Info(ctx context.Context, msg string)
	Error(ctx context.Context, msg string)
}
