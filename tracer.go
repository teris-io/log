// Copyright 2017 teris.io. License MIT

package log

import "time"

type Tracer interface {
	Trace()
}

func NewTracer(logger Logger, start time.Time) Tracer {
	return &tracer{logger: logger, start: start}
}

type tracer struct {
	logger Logger
	start  time.Time
}

func (t *tracer) Trace() {
	d := float64(time.Since(t.start)) * 1e-9
	t.logger.With("duration", d).Log("traced")
}
