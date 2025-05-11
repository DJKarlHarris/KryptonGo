package core

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (app *App) InitLog() (err error) {
	app.logFile, err = os.OpenFile("./KryptonGo.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		println(err.Error())
		return
	}

	fileSyncer := zapcore.AddSync(app.logFile)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	core := zapcore.NewCore(encoder, fileSyncer, zap.DebugLevel)

	app.logger = zap.New(core, zap.AddCaller())
	app.sugarLogger = app.logger.Sugar()

	return err
}

func SLOG() *zap.SugaredLogger {
	return app.sugarLogger
}

func LOG() *zap.Logger {
	return app.logger
}
