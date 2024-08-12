package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v11"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Environment struct {
	BaseFileUrl   string `env:"BASE_FILE_URL,required"`
	FileUrlSuffix string `env:"FILE_URL_SUFFIX,required"`
	FileUrl       string
}

var ENV Environment

// LoadConfig loads all app configs necessary & sets ENV vars
func LoadConfig() {
	err := env.Parse(&ENV)
	if err != nil {
		log.Panicf("%+v\n", err)
	}

	// construct file URL based on current month
	// This is because the file URL returned an error on the GET request with the July month after the month change.
	// This way, the request is always made with the current month
	ENV.FileUrl = ENV.BaseFileUrl + time.Now().Format("2006-01") + ENV.FileUrlSuffix

	initLogger()
}

// Zapper is an exported variable to be used in the utils pkg for logging
// This set & export is designed to simplify the app start up process by initializing app configuration & logging
var Zapper *zap.Logger

// initLogger initializes the zap logger in development mode at Info level
func initLogger() {
	var err error
	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
		Encoding:          "console",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	Zapper, err = cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	defer Zapper.Sync()
}
