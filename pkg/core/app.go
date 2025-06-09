package core

import (
	"KryptonGo/pkg/res"
	"database/sql"
	"os"

	"go.uber.org/zap"
)

var (
	app App
)

type App struct {
	logFile     *os.File
	tables      *res.Tables
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
	config      Config
	db          *sql.DB
}

func Init() (err error) {
	// init log
	if err = app.InitLog(); err != nil {
		SLOG().Errorf("init log fail err", err.Error())
		return err
	}

	// init table
	app.InitRes()

	// load config
	if err = app.LoadConfig("./cfg/config.yaml"); err != nil {
		return err
	}

	// init db
	if err = app.InitDB(); err != nil {
		return err
	}

	SLOG().Info("app Init succ!")

	return err
}

func Exit() {
	app.FreeLog()
	app.FreeDb()
}

func getApp() App {
	return app
}
