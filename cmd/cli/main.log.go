package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	endcoder := getEncoderLog()
	sync := getWriterLog()
	core := zapcore.NewCore(endcoder, sync, zapcore.InfoLevel)
	logger := zap.New(core)
}


func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"

	return zapcore.NewConsoleEncoder(encodeConfig)
}

func getWriterLog() zapcore.WriteSyncer {
	file, _  := os.OpenFile("./log/log.txt", os.O_RDWR, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}