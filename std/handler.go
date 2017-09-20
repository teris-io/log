// Copyright 2017 teris.io. License MIT

package std

import (
	"code.teris.io/go/log"
	"fmt"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	stdlog "log"
	"strings"
	"time"
)

func Use() {
	stdlog.SetFlags(0)
	log.SetFactory(&factory{})
}

type factory struct {
}

func (f *factory) With(k string, v interface{}) log.Logger {
	return &logger{lvl: log.Unset, fields: []Field{{Name: k, Value: v}}}
}

type Formatter func(start time.Time, level log.Level, fields []Field, msg string) string

func SetFormatter(f Formatter) {
	formatter = f
}

var formatter = func(start time.Time, lvl log.Level, fields []Field, msg string) string {
	timestr := color.CyanString(start.Format("02 15:04:05.000000"))
	lvlstr := ""
	switch lvl {
	case log.Debug:
		lvlstr = fmt.Sprintf(" %s", color.New(color.Bold, color.FgYellow).Sprint("DBG"))
	case log.Info:
		lvlstr = fmt.Sprintf(" %s", color.New(color.Bold, color.FgGreen).Sprint("INF"))
	case log.Error:
		lvlstr = fmt.Sprintf(" %s", color.New(color.Bold, color.FgRed).Sprint("ERR"))
	}
	fieldstr := ""
	for _, f := range fields {
		if fieldstr != "" {
			fieldstr += ", "
		}
		fieldstr += fmt.Sprintf("{%s: %v}", color.New(color.Bold).Sprint(f.Name), f.Value)
	}
	return fmt.Sprintf("%s%s %s %s", timestr, lvlstr, msg, fieldstr)
}

type logger struct {
	lvl    log.Level
	fields []Field
}

type Field struct {
	Name  string
	Value interface{}
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func (l *logger) With(k string, v interface{}) log.Logger {
	return &logger{lvl: l.lvl, fields: append([]Field{{Name: k, Value: v}}, l.fields...)}
}

func (l *logger) WithLevel(lvl log.Level) log.Logger {
	return &logger{lvl: lvl, fields: append([]Field{}, l.fields...)}
}

func (l *logger) WithError(err error) log.Logger {
	ctx := []Field{{Name: "error", Value: err.Error()}}

	if s, ok := err.(stackTracer); ok {
		frame := s.StackTrace()[0]

		name := fmt.Sprintf("%n", frame)
		file := fmt.Sprintf("%+s", frame)
		line := fmt.Sprintf("%d", frame)

		parts := strings.Split(file, "\n\t")
		if len(parts) > 1 {
			file = parts[1]
		}

		ctx = append(ctx, Field{Name: "source", Value: fmt.Sprintf("%s(%s:%s)", name, file, line)})
	}
	return &logger{lvl: log.Error, fields: append(ctx, l.fields...)}
}

func (l *logger) Log(msg string) log.Tracer {
	start := time.Now()
	stdlog.Println(formatter(start, l.lvl, l.fields, msg))
	return log.NewTracer(&logger{lvl: l.lvl, fields: append([]Field{}, l.fields...)}, start)
}

func (l *logger) Logf(format string, v ...interface{}) log.Tracer {
	return l.Log(fmt.Sprintf(format, v...))
}
