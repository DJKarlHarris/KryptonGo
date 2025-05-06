package main

import (
	"KryptonGo/pkg/luban"
	"KryptonGo/pkg/res"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type App struct {
	m_logFile     *os.File
	m_tables      *res.Tables
	m_logger      *zap.Logger
	m_sugarLogger *zap.SugaredLogger
}

var app App

func exit() {
	app.m_logger.Sync()
	app.m_logFile.Close()
}

func InitLog() (err error) {
	app.m_logFile, err = os.OpenFile("./KryptonGo.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		println(err.Error())
		return
	}

	fileSyncer := zapcore.AddSync(app.m_logFile)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	core := zapcore.NewCore(encoder, fileSyncer, zap.DebugLevel)

	app.m_logger = zap.New(core, zap.AddCaller())
	app.m_sugarLogger = app.m_logger.Sugar()

	return err
}

func main() {
	var err error

	//load table
	if app.m_tables, err = res.NewTables(luban.Loader); err != nil {
		println(err.Error())
		return
	}

	for idx, reward := range app.m_tables.TbReward.GetDataList() {
		fmt.Printf("idx:%d id:%d name:%s count:%d\n", idx, reward.Id, reward.Name, reward.Count)
	}

	if err = InitLog(); err != nil {
		println(err.Error())
		return
	}

	var url = "www.4399.com"
	app.m_sugarLogger.Infow("test",
		"url", url,
		"time", 3,
		"backoff", time.Second,
	)

	app.m_sugarLogger.Infof("test, url:%s, time:%d, backoff:%s", url, 3, time.Second)

	exit()
}
