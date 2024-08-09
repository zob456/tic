package utils

import (
	"fmt"

	"github.com/zob456/tic/internal/config"
)

func InfoLogger(args ...any) {
	config.Zapper.Info(fmt.Sprintf("%v", args...))
}

func ErrorLogger(err error) {
	config.Zapper.Error(fmt.Sprintf("%v", err))
}

func ErrorLoggerWithReturn(err error) error {
	config.Zapper.Error(fmt.Sprintf("%v", err))
	return err
}

func PanicLogger(err error) {
	config.Zapper.Panic(fmt.Sprintf("%v", err))
}
