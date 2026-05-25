package xlog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
)

func InitXLog() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
}

func Info(ctx context.Context, msg string) {
	slog.Info(msg, getSource(), getTraceId(ctx))
}

func Debug(ctx context.Context, msg string) {
	slog.Debug(msg, getSource(), getTraceId(ctx))
}

func Warn(ctx context.Context, msg string) {
	slog.Warn(msg, getSource(), getTraceId(ctx))
}

func Error(ctx context.Context, err error, msg string) {
	slog.Error(msg, getSource(), getTraceId(ctx), slog.String("error", err.Error()))
}

func Fatal(ctx context.Context, err error, msg string) {
	Error(ctx, err, msg)
	os.Exit(1)
}

func getSource() slog.Attr {
	_, file, line, ok := runtime.Caller(2)
	source := "unknown source"
	if ok {
		source = fmt.Sprintf("%s:%d", file, line)
	}
	return slog.String("source", source)
}

func getTraceId(ctx context.Context) slog.Attr {
	var traceId string
	if ctx != nil && ctx.Value("traceId") != nil {
		traceId = ctx.Value("traceId").(string)
	}
	if traceId == "" {
		traceId = "no traceId"
	}
	return slog.String("traceId", traceId)
}
