// Copyright (c) 2017. Oleg Sklyar & teris.io. All rights reserved.
// See the LICENSE file in the project root for licensing information.

// Package apex provides a logger implementation using the github.com/apex/log backends.
package apex

import (
	"fmt"
	"time"

	alog "github.com/apex/log"
	"github.com/teris-io/log"
)

var _ log.Logger = (*logger)(nil)

type logger struct {
	*factory
	lvl log.LoggerLevel
	ctx *alog.Entry
}

func (l *logger) Level(lvl log.LoggerLevel) log.Logger {
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

func (l *logger) Error(err error) log.Logger {
	return &logger{factory: l.factory, lvl: l.lvl, ctx: l.ctx.WithError(err)}
}

func (l *logger) Log(msg string) log.Tracer {
	if l.lvl >= l.factory.min {
		switch {
		case l.lvl <= log.DebugLevel:
			l.ctx.Debug(msg)
		case l.lvl == log.InfoLevel:
			l.ctx.Info(msg)
		case l.lvl >= log.ErrorLevel:
			l.ctx.Error(msg)
		}
	}
	return log.NewTracer(&logger{factory: l.factory, lvl: l.lvl, ctx: l.ctx}, time.Now())
}

func (l *logger) Logf(format string, v ...interface{}) log.Tracer {
	if l.lvl >= l.factory.min {
		return l.Log(fmt.Sprintf(format, v...))
	}
	return l.Log("")
}
