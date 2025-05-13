package core

import (
	"KryptonGo/pkg/res"
	"os"

	"go.uber.org/zap"
)

var app App

type App struct {
	logFile     *os.File
	tables      *res.Tables
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
	config      Config
}

func Init() (err error) {
	// init log
	if err = app.InitLog(); err != nil {
		SLOG().Infof("init log fail err", err.Error())
		return err
	}

	// init table
	app.InitRes()

	// load config
	if err = app.LoadConfig("./cfg/config.yaml"); err != nil {
		SLOG().Infof("load config fail %s", err.Error())
		return err
	}

	SLOG().Info("app Init succ!")

	return err
}

func Exit() {
	app.FreeLog()
}
