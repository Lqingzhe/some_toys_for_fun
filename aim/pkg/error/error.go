package newerror

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap/zapcore"
)

const (
	LevelDebug = zapcore.DebugLevel // -1  调试
	LevelInfo  = zapcore.InfoLevel  //  0  一般信息
	LevelWarn  = zapcore.WarnLevel  //  1  警告
	LevelError = zapcore.ErrorLevel //  2  错误
	LevelFatal = zapcore.FatalLevel //  5  致命（谨慎使用）
)

type Error struct {
	HttpCode    int
	HttpMessage string

	StatueCode ErrorStatue
	LogMessage error
	LogLevel   zapcore.Level
}
type Option struct {
	id        uint64
	operation string
	ip        string
}

func MakeError(httpCode int, statueCode ErrorStatue, httpMessage string, err error, logLevel zapcore.Level) *Error {
	return &Error{
		HttpCode:    httpCode,
		StatueCode:  statueCode,
		HttpMessage: httpMessage,
		LogMessage:  fmt.Errorf("-> %w", err),
		LogLevel:    logLevel,
	}

}
func (e *Error) Error() string {
	if e == nil || e.LogMessage == nil {
		return ""
	}
	return e.LogMessage.Error()
}
func TranslateError(err error) *Error {
	if err == nil {
		return nil
	}
	err2, ok := err.(*Error)
	if !ok {
		return MakeError(http.StatusInternalServerError, CodeInternalError, "Type Assertion Error", fmt.Errorf("%s", `Type Assertion To "*newerror.Error" Error`), LevelFatal).AddErrorTrace("error:TranslateError")
	}
	return err2
}

func (e *Error) AddErrorTrace(trace string) *Error {
	if e == nil {
		return nil
	}
	e.LogMessage = fmt.Errorf(" %s / %w ", trace, e.LogMessage)
	return e
}
func (o *Option) OptionInfo() (uint64, string, string) {
	return o.id, o.operation, o.ip
}
func IsContextError(err error) (bool, *Error) {
	if errors.Is(err, context.DeadlineExceeded) {
		return true, MakeError(http.StatusGatewayTimeout, CodeNetworkTimeout, "Time Out", err, LevelWarn)
	}
	if errors.Is(err, context.Canceled) {
		return true, MakeError(http.StatusGatewayTimeout, CodeNetworkTimeout, "Time Out", err, LevelWarn)
	}
	return false, nil
}
