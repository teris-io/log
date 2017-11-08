// Copyright (c) 2017. Oleg Sklyar & teris.io. All rights reserved.
// See the LICENSE file in the project root for licensing information.

package std

import (
	"io"
	stdlog "log"

	"github.com/teris-io/log"
)

// NewFactory creates a new log Factory instance for the implementation based
// on the Go standard log.
func NewFactory(std *stdlog.Logger, min log.LoggerLevel, fmt FmtFun) log.Factory {
	return &factory{std: std, min: min, fmt: fmt}
}

// Use activates the logger implementation (based on the Go standard log) to be
// used for static logging via the log package static functions.
func Use(out io.Writer, min log.LoggerLevel, fmt FmtFun) {
	std := stdlog.New(out, "", 0)
	log.SetFactory(NewFactory(std, min, fmt))
}

type factory struct {
	std *stdlog.Logger
	min log.LoggerLevel
	fmt FmtFun
}

var _ log.Factory = (*factory)(nil)

func (f *factory) New() log.Logger {
	return &logger{factory: f, lvl: log.UnsetLevel}
}

func (f *factory) Threshold(min log.LoggerLevel) {
	f.min = min
}
