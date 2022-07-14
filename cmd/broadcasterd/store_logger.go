package main

import (
	"strings"

	"github.com/seletskiy/go-log"
)

type StoreLogger struct {
	verbose bool
}

func (logger StoreLogger) Printf(format string, args ...interface{}) {
	log.Infof(nil, strings.TrimSpace("{migrations} "+format), args...)
}

func (logger StoreLogger) Verbose() bool {
	return logger.verbose
}
