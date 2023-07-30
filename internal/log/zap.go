package log

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log = Nop
var Nop = zap.NewNop().Sugar()

func NewLogger(project, name string) (*zap.SugaredLogger, error) {
	//timeString := time.Now().Format("2006_01_02-15_04_05")
	//pathName := strings.ToLower("logs_" + project + "_" + name)
	//cwd, err := os.Getwd()
	//if err != nil {
	//	return nil, errors.Wrap(err, "os.Getwd")
	//}

	//path := filepath.Join(cwd, pathName, timeString+".log")
	//newFilePath := filepath.FromSlash(path)
	//
	//if err := os.Mkdir(pathName, os.ModePerm); err != nil {
	//	if errors.Is(err, os.ErrExist) {
	//
	//	} else {
	//		return nil, errors.Wrap(err, "os.Mkdir")
	//	}
	//
	//}
	//
	//f, err := os.Create(newFilePath)
	//if err != nil {
	//	return nil, errors.Wrap(err, "os.Create")
	//}

	//writers := []zapcore.WriteSyncer{f, os.Stdout}
	writers := []zapcore.WriteSyncer{os.Stdout}

	c := zapcore.NewCore(zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		FunctionKey:    "fn",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(time.DateTime),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}), zapcore.NewMultiWriteSyncer(writers...), zapcore.DebugLevel)

	l := zap.New(c)
	l = l.WithOptions(zap.AddStacktrace(zap.ErrorLevel))

	Log = l.Sugar()

	return Log, nil
}
