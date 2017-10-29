// Copyright (c) 2017. Oleg Sklyar & teris.io. All rights reserved.
// See the LICENSE file in the project root for licensing information.

package apex

import (
	alog "github.com/apex/log"
	"github.com/teris-io/log"
)

func NewFactory(root *alog.Logger) log.Factory {
	return &factory{root: root, min: log.UnsetLevel}
}

func Use(root *alog.Logger) log.Factory {
	factory := NewFactory(root)
	log.SetFactory(factory)
	return factory
}

type factory struct {
	root *alog.Logger
	min  log.LoggerLevel
}

var _ log.Factory = (*factory)(nil)

func (f *factory) New() log.Logger {
	return &logger{factory: f, lvl: log.UnsetLevel, ctx: alog.NewEntry(f.root)}
}

func (f *factory) Threshold(min log.LoggerLevel) {
	f.min = min
}
