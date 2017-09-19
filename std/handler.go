// Copyright 2017 teris.io. License MIT

package std

import (
	"code.teris.io/go/log"
	"fmt"
	"strings"
	"github.com/pkg/errors"
	stdlog "log"
	"time"
	"github.com/fatih/color"
)

func Use() {
	stdlog.SetFlags(0)
	log.SetFactory(&factory{})
}

type factory struct {
}

func (f *factory) With(k string, v interface{}) log.Logger {
	return &logger{lvl: log.Unset, ctx: []kv{{k: k, v: v}}}
}

func SetFormatter(f func(start time.Time, level log.Level, msg string, ctx interface{}) string) {
	formatter = f
}

var formatter = func(start time.Time, lvl log.Level, msg string, ctx interface{}) string {
	lvlstr := ""
	switch lvl {
	case log.Debug:
		lvlstr = fmt.Sprintf(" [%s]", color.CyanString("DBG"))
	case log.Info:
		lvlstr = fmt.Sprintf(" [%s]", color.GreenString("INF"))
	case log.Error:
		lvlstr = fmt.Sprintf(" [%s]", color.RedString("ERR"))
	}
	return fmt.Sprintf("%s%s %s %v", start.Format("02 15:04:05.999999"), lvlstr, msg, ctx)
}

type logger struct {
	lvl log.Level
	ctx []kv
}

type kv struct {
	k string
	v interface{}
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func (l *logger) With(k string, v interface{}) log.Logger {
	return &logger{lvl: l.lvl, ctx: append([]kv{{k: k, v: v}}, l.ctx...)}
}

func (l *logger) WithLevel(lvl log.Level) log.Logger {
	return &logger{lvl: lvl, ctx: append([]kv{}, l.ctx...)}
}

func (l *logger) WithError(err error) log.Logger {
	ctx := []kv{{k: "error", v: err.Error()}}

	if s, ok := err.(stackTracer); ok {
		frame := s.StackTrace()[0]

		name := fmt.Sprintf("%n", frame)
		file := fmt.Sprintf("%+s", frame)
		line := fmt.Sprintf("%d", frame)

		parts := strings.Split(file, "\n\t")
		if len(parts) > 1 {
			file = parts[1]
		}

		ctx = append(ctx, kv{k: "source", v: fmt.Sprintf("%s(%s:%s)", name, file, line)})
	}
	return &logger{lvl: log.Error, ctx: append(ctx, l.ctx...)}
}

func (l *logger) Log(msg string) log.Tracer {
	start := time.Now()
	stdlog.Println(formatter(start, l.lvl, msg, l.ctx))
	return log.NewTracer(&logger{lvl: l.lvl, ctx: append([]kv{}, l.ctx...)}, start)
}

func (l *logger) Logf(format string, v ...interface{}) log.Tracer {
	return l.Log(fmt.Sprintf(format, v...))
}
