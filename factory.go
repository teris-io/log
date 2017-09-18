// Copyright 2017 teris.io. License MIT

package log

var factory Factory

type Factory interface {
	With(k string, v interface{}) Logger
}

func SetFactory(f Factory) {
	factory = f
}

func With(k string, v interface{}) Logger {
	return factory.With(k, v)
}
