// Copyright (c) 2017 Oleg Sklyar & teris.io.

package log

var factory Factory = &noop{}

// SetFactory sets the static logger factory.
func SetFactory(f Factory) {
	factory = f
}

// New returns a new logger instance from the static factory.
func New() Logger {
	return factory.New()
}

// Level returns a new logger instance from the factory setting its log level to the value supplied.
func Level(lvl LogLevel) Logger {
	return factory.New().Level(lvl)
}

// Field returns a new logger instance from the factory setting a field value as supplied.
func Field(k string, v interface{}) Logger {
	return factory.New().Field(k, v)
}

// Fields returns a new logger instance from the factory setting field values as supplied.
func Fields(data map[string]interface{}) Logger {
	return factory.New().Fields(data)
}

// WithError returns a new logger instance from the factory setting the error as supplied.
func WithError(err error) Logger {
	return factory.New().WithError(err)
}

// Log constructs a new logger instance from the factory with no context and logs a message.
func Log(msg string) Tracer {
	return factory.New().Log(msg)
}

// Logf constructs a new logger instance from the factory with no context and logs a formatted message.
func Logf(format string, v ...interface{}) Tracer {
	return factory.New().Logf(format, v...)
}
