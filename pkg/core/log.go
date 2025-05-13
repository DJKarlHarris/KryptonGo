package core

import (
	"os"
	"time"

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

	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}

	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(caller.TrimmedPath() + " " + caller.Function)
	}

	EncoderCfg := zap.NewProductionEncoderConfig()
	EncoderCfg.EncodeTime = customTimeEncoder
	EncoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	EncoderCfg.EncodeCaller = customCallerEncoder
	encoder := zapcore.NewJSONEncoder(EncoderCfg)

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

func (app *App) FreeLog() {
	app.logger.Sync()
	app.logFile.Close()
}
