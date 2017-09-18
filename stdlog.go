// Copyright 2017 teris.io. License MIT

package log

import (
	stdlog "log"
)

type stdlogger struct {

	lock chan int

	fields map[string]interface{}
}

func newStdlogger(ctx string) Logger {
	return nil
}

func (l *stdlogger) With(k string, v interface{}) Logger {
	go func() {
		if l.fields == nil {
			l.fields = make(map[string]interface{})
		}
		l.fields[k] = v
	}()
	return l
}

func (l *stdlogger) Log(msg string) Tracer {

	return NewTracer(l)
}

func (l *stdlogger) Logf(format string, v... interface{}) Tracer {
	stdlog.Printf(format, v...)
	return NewTracer(l)
}
