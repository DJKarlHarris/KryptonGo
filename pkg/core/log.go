package core

import (
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func (app *App) InitLog() (err error) {
	logRotate := &lumberjack.Logger{
		Filename:   "./log/KryptonGo.log", // 日志路径
		MaxSize:    100,                   // 单个文件最大大小（MB）
		MaxBackups: 30,                    // 保留旧文件的最大数量
		MaxAge:     30,                    // 保留旧文件的最大天数
		Compress:   true,                  // 是否压缩/归档旧文件
	}

	fileSyncer := zapcore.AddSync(logRotate)

	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}

	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {

		funcName := ""
		idx := strings.LastIndexByte(caller.Function, '.')
		if idx != -1 {
			funcName = caller.Function[idx+1:]
		}

		enc.AppendString(caller.TrimmedPath() + " " + funcName)
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
