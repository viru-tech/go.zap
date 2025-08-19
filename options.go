package logger

import (
	"log"

	"go.uber.org/zap"
)

// Option allows to configure logger.
type Option func(*builder)

// WithConfig sets zap config for the logger.
// Note: this option should be used before any others since
// they may change config's fields.
func WithConfig(config zap.Config) Option { //nolint:gocritic
	return func(b *builder) {
		b.config = &config
	}
}

// WithEncoding sets the logger's encoding. Valid values are "json" and
// "console", as well as any third-party encodings registered via
// zap.RegisterEncoder.
func WithEncoding(enc string) Option {
	return func(b *builder) {
		b.config.Encoding = enc
	}
}

// WithFields adds structured context to all log entries.
func WithFields(ff ...zap.Field) Option {
	return func(b *builder) {
		b.buildOpts = append(b.buildOpts, zap.Fields(ff...))
	}
}

// WithLevel sets level, if not empty is passed.
// Zap config's default is used otherwise.
func WithLevel(level string) Option {
	return func(b *builder) {
		if level == "" {
			return
		}

		if err := b.config.Level.UnmarshalText([]byte(level)); err != nil && level != "" {
			log.Printf("failed to set log level %q: %v", level, err)
		}
	}
}
