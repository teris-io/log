// Copyright (c) 2017. Oleg Sklyar & teris.io. All rights reserved.
// See the LICENSE file in the project root for licensing information.

// Package std provides a logger implementation via the Go built-in logger.
package std

import (
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/teris-io/log"
)

type Field struct {
	Name  string
	Value interface{}
}

type tracer interface {
	StackTrace() errors.StackTrace
}

type logger struct {
	*factory
	lvl    log.LoggerLevel
	fields []Field
}

var _ log.Logger = (*logger)(nil)

func (l *logger) Level(lvl log.LoggerLevel) log.Logger {
	return &logger{factory: l.factory, lvl: lvl, fields: append([]Field{}, l.fields...)}
}

func (l *logger) Field(k string, v interface{}) log.Logger {
	return &logger{factory: l.factory, lvl: l.lvl, fields: append([]Field{{Name: k, Value: v}}, l.fields...)}
}

func (l *logger) Fields(data map[string]interface{}) log.Logger {
	var fields []Field
	for k, v := range data {
		fields = append(fields, Field{Name: k, Value: v})
	}
	return &logger{factory: l.factory, lvl: l.lvl, fields: fields}
}

func (l *logger) Error(err error) log.Logger {
	ctx := []Field{{Name: "error", Value: err.Error()}}

	if s, ok := err.(tracer); ok {
		frame := s.StackTrace()[0]

		name := fmt.Sprintf("%s", frame)
		file := fmt.Sprintf("%+s", frame)
		line := fmt.Sprintf("%d", frame)

		parts := strings.Split(file, "\n\t")
		if len(parts) > 1 {
			file = parts[1]
		}

		ctx = append(ctx, Field{Name: "source", Value: fmt.Sprintf("%s(%s:%s)", name, file, line)})
	}
	return &logger{factory: l.factory, lvl: log.ErrorLevel, fields: append(ctx, l.fields...)}
}

func (l *logger) Log(msg string) log.Tracer {
	start := time.Now()
	if l.lvl >= l.factory.min {
		l.factory.std.Println(l.factory.fmt(start, l.lvl, msg, l.fields))
	}
	tracelogger := &logger{factory: l.factory, lvl: l.lvl, fields: append([]Field{}, l.fields...)}
	return log.NewTracer(tracelogger, start)
}

func (l *logger) Logf(format string, v ...interface{}) log.Tracer {
	if l.lvl >= l.factory.min {
		return l.Log(fmt.Sprintf(format, v...))
	} else {
		return l.Log("")
	}
}
