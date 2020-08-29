package main

import (
	"fmt"
	"go.uber.org/zap"
	"log"
)

type reporting struct {
	logger *zap.Logger
}

func (r *reporting) getLogger() *zap.Logger {
	return r.logger
}

const dev = "dev"

func newReporting(env string) *reporting {
	return &reporting{
		logger: getLogger(env),
	}
}

func getLogger(env string) *zap.Logger {
	var logger *zap.Logger
	var err error

	if env == dev {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		fmt.Println("mid", err)
		log.Fatal(err)
	}

	defer func() {
		if err := logger.Sync(); err != nil {
			fmt.Println(err)
		}
	}()

	return logger
}
