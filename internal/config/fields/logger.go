package fields

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
	"strings"
)

type LoggerLevel zapcore.Level
type LoggerFormat string

func (f *LoggerLevel) parseValue(rawValue string) error {
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(rawValue)); err != nil {
		return err
	}
	*f = LoggerLevel(level)
	return nil
}

func (f *LoggerLevel) UnmarshalYAML(value *yaml.Node) error {
	return f.parseValue(value.Value)
}

func (f *LoggerLevel) Decode(rawValue string) error {
	return f.parseValue(rawValue)
}

const (
	loggerFormatJSON = LoggerFormat("JSON")
	loggerFormatTEXT = LoggerFormat("TEXT")
)

func (f *LoggerFormat) parseValue(rawValue string) error {
	*f = LoggerFormat(strings.ToUpper(rawValue))
	if *f != "JSON" && *f != "TEXT" {
		return fmt.Errorf("invalid loffer format '%s', must be one of TEXT or JSON ", rawValue)
	}
	return nil
}

func (f *LoggerFormat) UnmarshalYAML(value *yaml.Node) error {
	return f.parseValue(value.Value)
}

func (f *LoggerFormat) Decode(value string) error {
	return f.parseValue(value)
}

func (f *LoggerFormat) IsJSON() bool {
	return *f == loggerFormatJSON
}

func (f *LoggerFormat) IsTEXT() bool {
	return *f == loggerFormatTEXT
}
