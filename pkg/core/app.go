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
}

func Init() (err error) {
	// init log
	err = app.InitLog()

	// init table
	app.InitRes()

	return err
}

func Exit() {
	app.FreeLog()
}
