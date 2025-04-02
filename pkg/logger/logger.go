package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"main.go/pkg/setting"
)

type Logger struct {
	*zap.Logger
}

func NewLogger(setting setting.LoggerSetting) *Logger {
	var level zapcore.Level
	switch setting.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.DebugLevel
	}

	// 📌 Cấu hình log file với Lumberjack
	fileLogger := zapcore.AddSync(&lumberjack.Logger{
		Filename:   setting.LogFile,
		MaxSize:    setting.LogMaxSize,
		MaxBackups: setting.LogMaxBackups,
		MaxAge:     setting.LogMaxAge,
		Compress:   setting.LogCompress,
	})

	// 📌 Log ra terminal (stdout)
	consoleLogger := zapcore.AddSync(os.Stdout)

	// 📌 Cấu hình encoder
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:      "time",
		LevelKey:     "level",
		MessageKey:   "message",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	// 📌 Tạo encoder cho file (JSON) và console (Text)
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 📌 Tạo core để ghi log ra cả file và terminal
	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, fileLogger, level),       // Log ra file (JSON)
		zapcore.NewCore(consoleEncoder, consoleLogger, level), // Log ra terminal (text)
	)

	// 📌 Trả về logger
	return &Logger{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))}
}
