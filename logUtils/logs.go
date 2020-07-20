package logUtils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var SugarLogger *zap.SugaredLogger

func Infof(template string, args ...interface{}) {
	 SugarLogger.Infof(template,args...)
}
func Info(args ...interface{})  {
	SugarLogger.Info(args...)
}

func Debugf(template string, args ...interface{})  {
	SugarLogger.Debugf(template,args...)
}
func Debug(args ...interface{})  {
	SugarLogger.Debug(args...)
}

func Errorf(template string, args ...interface{})  {
	SugarLogger.Errorf(template,args...)
}
func Error(args ...interface{})  {
	SugarLogger.Error(args...)
}

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core,zap.AddCaller())
	SugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.OpenFile(" ./logs/monkey.log",os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	return zapcore.AddSync(file)
}