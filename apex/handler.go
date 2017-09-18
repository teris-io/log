// Copyright 2017 teris.io. License MIT

package apex

import (
	"code.teris.io/go/log"
	"fmt"
	alog "github.com/apex/log"
)

type factory struct {
}

func Use() {
	log.SetFactory(&factory{})
}

func (f *factory) With(k string, v interface{}) log.Logger {
	return &logger{lvl: log.Info, ctx: alog.WithField(k, v)}
}

type logger struct {
	lvl log.Level
	ctx *alog.Entry
}

func (l *logger) With(k string, v interface{}) log.Logger {
	return &logger{lvl: l.lvl, ctx: l.ctx.WithField(k, v)}
}

func (l *logger) WithLevel(lvl log.Level) log.Logger {
	return &logger{lvl: lvl, ctx: l.ctx}
}

func (l *logger) WithError(err error) log.Logger {
	return &logger{lvl: l.lvl, ctx: l.ctx.WithError(err)}
}

func (l *logger) Log(msg string) log.Tracer {
	switch {
	case l.lvl <= log.Debug:
		l.ctx.Debug(msg)
	case l.lvl == log.Info:
		l.ctx.Info(msg)
	case l.lvl >= log.Error:
		l.ctx.Error(msg)
	}
	return log.NewTracer(&logger{lvl: l.lvl, ctx: l.ctx})
}

func (l *logger) Logf(format string, v ...interface{}) log.Tracer {
	return l.Log(fmt.Sprintf(format, v...))
}
