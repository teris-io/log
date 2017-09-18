// Copyright 2017 teris.io. License MIT

package log

type Logger interface {

	With(k string, v interface{}) Logger

	Log(msg string) Tracer

	Logf(format string, v... interface{}) Tracer
}
