package core

import (
	"github.com/daiyuang/sshman/core/logger"
	"github.com/daiyuang/sshman/core/metadata"
	"github.com/daiyuang/sshman/core/ssh"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var commonModule = fx.Options(
	logger.Module,
	metadata.Module,
	ssh.Module,
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		fxLogger := &fxevent.ZapLogger{Logger: log}
		fxLogger.UseLogLevel(zapcore.DebugLevel)
		return fxLogger
	}),
)

func CreateContainer(op ...fx.Option) *fx.App {
	return fx.New(append([]fx.Option{commonModule}, op...)...)
}
