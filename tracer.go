// Copyright 2017 teris.io. License MIT

package log

import "time"

type Tracer interface {
	Trace()
}

func NewTracer(logger Logger) Tracer {
	return &tracer{logger: logger, start: time.Now()}
}

type tracer struct {
	logger Logger
	start  time.Time
}

func (t *tracer) Trace() {
	d := time.Since(t.start)
	t.logger.With("duration", d).Log("traced")
}
