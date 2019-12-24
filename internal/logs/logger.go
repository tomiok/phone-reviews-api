package logs

import (
	"fmt"
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func InitLogger() error {
	l, err := zap.NewDevelopment()

	if err != nil {
		_ = fmt.Errorf("cannot create zap logger %s", err.Error())
		return err
	}

	sugar = l.Sugar()

	return nil
}

func Log() *zap.SugaredLogger {
	return sugar
}
