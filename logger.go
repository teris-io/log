// Copyright 2017 teris.io. License MIT

package log

type Level int

const (
	Debug = iota - 1
	Info
	Error
)

type Logger interface {
	With(k string, v interface{}) Logger

	WithLevel(lvl Level) Logger

	WithError(err error) Logger

	Log(msg string) Tracer

	Logf(format string, v ...interface{}) Tracer
}
