package cmd

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sman/internal/logger"
	"sman/internal/metadata"
)

var commonModule = fx.Options(
	logger.Module,
	metadata.Module,
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		fxLogger := &fxevent.ZapLogger{Logger: log}
		fxLogger.UseLogLevel(zapcore.DebugLevel)
		return fxLogger
	}),
)

func createContainer(op ...fx.Option) *fx.App {
	allOpts := []fx.Option{commonModule}
	allOpts = append(allOpts, op...)
	return fx.New(allOpts...)
}
