package modules

import (
	"broadcaster/internal/config/fields"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Logger() *zap.Logger {
	return Register("logger", func(s string) (Module, error) {
		cfg := Config()

		cfg.Environment = fields.EnvironmentName(os.Getenv("BROADCASTER_ENV"))
		if cfg.Environment.IsEmpty() {
			cfg.Environment = fields.Dev
		}

		logConfig := zap.Config{
			Level:            zap.NewAtomicLevelAt(zapcore.Level(cfg.Logger.Level)),
			Development:      false,
			Encoding:         "json",
			EncoderConfig:    zap.NewProductionEncoderConfig(),
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
		}

		if !cfg.Environment.IsDev() {
			logConfig.Development = false
			logConfig.EncoderConfig = zap.NewProductionEncoderConfig()
			logConfig.Sampling = &zap.SamplingConfig{
				Initial:    100,
				Thereafter: 100,
			}
		} else {
			logConfig.Development = true
			logConfig.EncoderConfig = zap.NewDevelopmentEncoderConfig()
		}

		logConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		logConfig.EncoderConfig.TimeKey = "time"
		logConfig.EncoderConfig.NameKey = "name"
		logConfig.EncoderConfig.CallerKey = "caller"
		logConfig.EncoderConfig.LevelKey = "level"
		logConfig.EncoderConfig.MessageKey = "msg"
		logConfig.EncoderConfig.StacktraceKey = "trace"

		if cfg.Logger.Format.IsJSON() {
			logConfig.Encoding = "json"
		} else {
			logConfig.Encoding = "console"
		}

		logger, err := logConfig.Build(zap.AddStacktrace(zap.ErrorLevel))
		if err != nil {
			return nil, err
		}

		return logger.With(zap.String("env", string(cfg.Environment))), nil
	}).(*zap.Logger)
}
