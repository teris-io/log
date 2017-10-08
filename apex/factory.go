// Copyright (c) 2017 Oleg Sklyar & teris.io.

package apex

import (
	"code.teris.io/util/log"
	alog "github.com/apex/log"
)

func NewFactory(root *alog.Logger) log.Factory {
	return &factory{root: root, min: log.Unset}
}

func Use(root *alog.Logger) log.Factory {
	factory := NewFactory(root)
	log.SetFactory(factory)
	return factory
}

type factory struct {
	root *alog.Logger
	min  log.LogLevel
}

var _ log.Factory = (*factory)(nil)

func (f *factory) New() log.Logger {
	return &logger{factory: f, lvl: log.Unset, ctx: alog.NewEntry(f.root)}
}

func (f *factory) Threshold(min log.LogLevel) {
	f.min = min
}
