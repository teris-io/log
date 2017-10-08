// Copyright (c) 2017 Oleg Sklyar & teris.io.

package std

import (
	stdlog "log"

	"code.teris.io/util/log"
	"io"
)

func NewFactory(std *stdlog.Logger, min log.LogLevel, fmt FmtFun) log.Factory {
	return &factory{std: std, min: min, fmt: fmt}
}

func Use(out io.Writer, min log.LogLevel, fmt FmtFun) {
	std := stdlog.New(out, "", 0)
	log.SetFactory(NewFactory(std, min, fmt))
}

type factory struct {
	std *stdlog.Logger
	min log.LogLevel
	fmt FmtFun
}

var _ log.Factory = (*factory)(nil)

func (f *factory) New() log.Logger {
	return &logger{factory: f, lvl: log.Unset}
}

func (f *factory) Threshold(min log.LogLevel) {
	f.min = min
}
