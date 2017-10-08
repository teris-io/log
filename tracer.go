// Copyright (c) 2017 Oleg Sklyar & teris.io

package log

import "time"

// Tracer defines a utility to trace execution time as returned by the logger Log and Logf methods.
type Tracer interface {

	// Trace computes the time elapsed since the tracer was created. It then logs an entry
	// with message "traced" and field "duration" amounting to the execution duration in seconds.
	Trace() float64
}

// NewTracer creates a new tracer instance: assumed to be used by logger implementations
// when implementing the Log and Logf methods only.
func NewTracer(logger Logger, start time.Time) Tracer {
	return &tracer{logger: logger, start: start}
}

type tracer struct {
	logger Logger
	start  time.Time
}

func (t *tracer) Trace() float64 {
	d := float64(time.Since(t.start)) * 1e-9
	t.logger.Field("duration", d).Log("traced")
	return d
}
