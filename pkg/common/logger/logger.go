package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

// 集成zap日志管理工具

var lg *zap.Logger

func init() {
	var err error
	//cfg := conf.Cfg.LogConfig
	//mode := conf.Cfg.Mode
	mode := "dev"

	//todo 配置
	writeSyncer := getLogWriter("./logs/im.log", 500, 3, 28)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte("debug"))
	if err != nil {
		panic(err)
	}

	var core zapcore.Core
	if mode == "dev" {
		// 开发者模式 日志输出到终端
		log.Println("debug mode")
		consoleEncoderCfg := zap.NewDevelopmentEncoderConfig()
		consoleEncoderCfg.EncodeCaller = zapcore.FullCallerEncoder
		consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderCfg)
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(log.Writer()), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}
	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)
	//Logger().Info(fmt.Sprintf("logger level: %s", lg.Level()))
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//encoderConfig.TimeKey = "time"
	//encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	//encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 集成lumberjack 滚动日志功能
func getLogWriter(fileName string, maxSize, maxBuckup, maxAge int) zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBuckup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberjackLogger)
}

func L() *zap.Logger {
	return zap.L()
}
