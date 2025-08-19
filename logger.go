package logger

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type builder struct {
	config    *zap.Config
	buildOpts []zap.Option
}

// New returns new zap logger with built-in level toggler.
func New(opts ...Option) (*zap.Logger, error) {
	defaultConfig := zap.NewProductionConfig()
	b := builder{config: &defaultConfig}
	for _, opt := range opts {
		opt(&b)
	}

	b.config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := b.config.Build(b.buildOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to build logger: %w", err)
	}

	go func() {
		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
		defer cancel()

		usr1 := make(chan os.Signal, 1)
		signal.Notify(usr1, syscall.SIGUSR1)
		initialLevel := b.config.Level.Level()

		for {
			select {
			case <-ctx.Done():
				return
			case <-usr1:
				logger.Info("caught SIGUSR1 signal, toggling log level")
				var nextLevel zapcore.Level
				if b.config.Level.Level() == zap.DebugLevel {
					nextLevel = initialLevel
				} else {
					nextLevel = zap.DebugLevel
				}
				b.config.Level.SetLevel(nextLevel)
				logger.Info("log level changed", zap.String("level", nextLevel.String()))
			}
		}
	}()

	return logger, nil
}
