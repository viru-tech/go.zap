//nolint:lll
package logger

import "go.uber.org/zap"

func ExampleNew_production() {
	config := zap.NewProductionConfig()

	// next configuration is for example output only.
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.TimeKey = ""

	logger, err := New(
		WithConfig(config),
		WithFields(
			zap.String("version", "v1.0.0"),
			zap.String("environment", "production"),
		),
	)
	if err != nil {
		panic(err)
	}

	logger.Debug("this is a debug message", zap.Bool("debug", true))
	logger.Info("this is an info message", zap.Bool("debug", false))

	// Output:
	// {"level":"info","caller":"go.zap/logger_test.go:25","msg":"this is an info message","version":"v1.0.0","environment":"production","debug":false}
}

func ExampleNew_staging() {
	config := zap.NewProductionConfig()

	// next configuration is for example output only.
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.TimeKey = ""

	logger, err := New(
		WithConfig(config),
		WithLevel(zap.DebugLevel.String()),
		WithFields(
			zap.String("version", "v1.0.0"),
			zap.String("environment", "staging"),
		),
	)
	if err != nil {
		panic(err)
	}

	logger.Debug("this is a debug message", zap.Bool("debug", true))
	logger.Info("this is an info message", zap.Bool("debug", false))

	// Output:
	// {"level":"debug","caller":"go.zap/logger_test.go:50","msg":"this is a debug message","version":"v1.0.0","environment":"staging","debug":true}
	// {"level":"info","caller":"go.zap/logger_test.go:51","msg":"this is an info message","version":"v1.0.0","environment":"staging","debug":false}
}

func ExampleNew_localJSON() {
	config := zap.NewDevelopmentConfig()

	// next configuration is for example output only.
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.TimeKey = ""

	logger, err := New(
		WithConfig(config),
		WithEncoding("json"),
		WithFields(
			zap.String("version", "abacaba"),
			zap.String("environment", "local"),
		),
	)
	if err != nil {
		panic(err)
	}

	logger.Debug("this is a debug message", zap.Bool("debug", true))
	logger.Info("this is an info message", zap.Bool("debug", false))

	// Output:
	// {"L":"DEBUG","C":"go.zap/logger_test.go:77","M":"this is a debug message","version":"abacaba","environment":"local","debug":true}
	// {"L":"INFO","C":"go.zap/logger_test.go:78","M":"this is an info message","version":"abacaba","environment":"local","debug":false}
}
