// Copyright (c) 2017. Oleg Sklyar & teris.io. All rights reserved.
// See the LICENSE file in the project root for licensing information.

package log

import "time"

// implements a no-operation logger so that the logger interface can be used
// out of the box without providing an actual implementation
type noop struct{}

var (
	_ Factory = (*noop)(nil)
	_ Logger  = (*noop)(nil)
)

func (n *noop) New() Logger {
	return &noop{}
}

func (n *noop) Threshold(lvl LoggerLevel) {
}

func (n *noop) Level(lvl LoggerLevel) Logger {
	return n
}

func (n *noop) Field(k string, v interface{}) Logger {
	return n
}

func (n *noop) Fields(data map[string]interface{}) Logger {
	return n
}

func (n *noop) Error(err error) Logger {
	return n
}

func (n *noop) Log(msg string) Tracer {
	return NewTracer(n, time.Now())
}

func (n *noop) Logf(format string, v ...interface{}) Tracer {
	return NewTracer(n, time.Now())
}
