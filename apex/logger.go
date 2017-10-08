// Copyright (c) 2017 Oleg Sklyar & teris.io

// Package apex provides a logger implementation using the github.com/apex/log backends.
package apex

import (
	"fmt"
	"time"

	"code.teris.io/util/log"
	alog "github.com/apex/log"
)

var _ log.Logger = (*logger)(nil)

type logger struct {
	*factory
	lvl log.LogLevel
	ctx *alog.Entry
}

func (l *logger) Level(lvl log.LogLevel) log.Logger {
	return &logger{factory: l.factory, lvl: lvl, ctx: l.ctx}
}

func (l *logger) Field(k string, v interface{}) log.Logger {
	return &logger{factory: l.factory, lvl: l.lvl, ctx: l.ctx.WithField(k, v)}
}

func (l *logger) Fields(data map[string]interface{}) log.Logger {
	fields := alog.Fields{}
	for k, v := range data {
		fields[k] = v
	}
	return &logger{factory: l.factory, lvl: l.lvl, ctx: l.ctx.WithFields(fields)}
}

func (l *logger) WithError(err error) log.Logger {
	return &logger{factory: l.factory, lvl: l.lvl, ctx: l.ctx.WithError(err)}
}

func (l *logger) Log(msg string) log.Tracer {
	if l.lvl >= l.factory.min {
		switch {
		case l.lvl <= log.Debug:
			l.ctx.Debug(msg)
		case l.lvl == log.Info:
			l.ctx.Info(msg)
		case l.lvl >= log.Error:
			l.ctx.Error(msg)
		}
	}
	return log.NewTracer(&logger{factory: l.factory, lvl: l.lvl, ctx: l.ctx}, time.Now())
}

func (l *logger) Logf(format string, v ...interface{}) log.Tracer {
	if l.lvl >= l.factory.min {
		return l.Log(fmt.Sprintf(format, v...))
	} else {
		return l.Log("")
	}
}
