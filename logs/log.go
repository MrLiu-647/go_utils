package logs

import (
	"context"
	"fmt"
)

type Logger interface {
	CtxDebug(ctx context.Context, format string, v ...interface{})
	CtxInfo(ctx context.Context, format string, v ...interface{})
	CtxWarn(ctx context.Context, format string, v ...interface{})
	CtxError(ctx context.Context, format string, v ...interface{})
	CtxFatal(ctx context.Context, format string, v ...interface{})
}

type internalLogger struct {
}

func (i *internalLogger) CtxDebug(ctx context.Context, format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (i *internalLogger) CtxInfo(ctx context.Context, format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (i *internalLogger) CtxWarn(ctx context.Context, format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (i *internalLogger) CtxError(ctx context.Context, format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (i *internalLogger) CtxFatal(ctx context.Context, format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

var logger Logger = &internalLogger{}

func SetLogger(log Logger) {
	logger = log
}

func CtxDebug(ctx context.Context, format string, v ...interface{}) {
	logger.CtxDebug(ctx, format, v...)
}

func CtxInfo(ctx context.Context, format string, v ...interface{}) {
	logger.CtxInfo(ctx, format, v...)
}

func CtxWarn(ctx context.Context, format string, v ...interface{}) {
	logger.CtxWarn(ctx, format, v...)
}

func CtxError(ctx context.Context, format string, v ...interface{}) {
	logger.CtxError(ctx, format, v...)
}

func CtxFatal(ctx context.Context, format string, v ...interface{}) {
	logger.CtxFatal(ctx, format, v...)
}
