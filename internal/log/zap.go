package log

import (
	"context"
	core "crypto_scripts/internal/tg"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Log = Nop
var Nop = zap.NewNop().Sugar()
var Alert = Nop

func NewLogger(project, name string) (*zap.SugaredLogger, error) {
	timeString := time.Now().Format("2006_01_02-15_04_05")
	pathName := strings.ToLower("logs_" + project + "_" + name)
	cwd, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, "os.Getwd")
	}

	path := filepath.Join(cwd, pathName, timeString+".log")
	newFilePath := filepath.FromSlash(path)

	if err := os.Mkdir(pathName, os.ModePerm); err != nil {
		if errors.Is(err, os.ErrExist) {

		} else {
			return nil, errors.Wrap(err, "os.Mkdir")
		}

	}

	f, err := os.Create(newFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "os.Create")
	}

	writers := []zapcore.WriteSyncer{f, os.Stdout}

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

	conn, err := grpc.Dial("170.64.160.93:912", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err == nil && conn != nil {

		w := NewTgWriter(conn, project, name)

		c = zapcore.RegisterHooks(c, func(entry zapcore.Entry) error {
			switch entry.Level {
			case zap.WarnLevel, zap.DPanicLevel, zap.ErrorLevel, zap.FatalLevel:
				_, _ = w.WriteE(entry)
			}
			return nil
		})
	}

	l := zap.New(c)
	l = l.WithOptions(zap.AddStacktrace(zap.ErrorLevel))

	Log = l.Sugar()

	return Log, nil
}

func NewTelegramAlert(project, name string) *zap.SugaredLogger {
	conn, err := grpc.Dial("170.64.160.93:912", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err == nil && conn != nil {
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
		}), zapcore.NewMultiWriteSyncer(NewTgWriter(conn, project, name)), zapcore.DebugLevel)
		l := zap.New(c)
		ll := l.Sugar()
		Alert = ll
		return Alert
	}

	return Nop
}

type TgWriter struct {
	ch            chan []byte
	cli           core.CoreServiceClient
	project, name string
}

func NewTgWriter(conn *grpc.ClientConn, project, name string) *TgWriter {
	w := &TgWriter{
		ch:      make(chan []byte, 100),
		project: project,
		name:    name,
	}

	w.cli = core.NewCoreServiceClient(conn)

	return w
}

func (w *TgWriter) WriteE(p zapcore.Entry) (n int, err error) {

	_, err = w.cli.Send(context.Background(), &core.SendReq{
		Name: w.project + "_" + w.name,
		Text: p.Message,
	})

	return len(p.Message), nil
}

func (w *TgWriter) Write(p []byte) (n int, err error) {

	_, err = w.cli.Send(context.Background(), &core.SendReq{
		Name: w.project + "_" + w.name,
		Text: string(p),
	})

	return len(p), nil
}

func (w *TgWriter) Sync() error {
	for {
		if len(w.ch) == 0 {
			close(w.ch)
			return nil
		}
	}
}
