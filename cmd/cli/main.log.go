package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	//1
	//sugar := zap.NewExample().Sugar()
	//sugar.Infof("Hello name:%s age:%d", "tipjs", 18)
	//logger := zap.NewExample()
	//logger.Info("Hello", zap.String("name", "tipjs"), zap.Int("age", 18))

	//2
	//logger := zap.NewExample()
	//logger.Info("hello")
	//logger, _ = zap.NewDevelopment()
	//logger.Info("hello new development")
	//logger, _ = zap.NewProduction()
	//logger.Info("hello new production")

	//3
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.DebugLevel)
	logger := zap.New(core)
	logger.Info("hello new production")
	logger.Error("bug nha ")
}
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}
func getWriterSync() zapcore.WriteSyncer {

	file, err := os.OpenFile("./log/log.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
