// Copyright (c) 2017 Oleg Sklyar & teris.io

package log

// LogLevel defines level markers for log entries.
type LogLevel int

const (
	// Unset level should not be output by logger implementation.
	Unset = iota - 2
	// Debug level marks detailed output for design purposes.
	Debug
	// Info level is the default log output marker.
	Info
	// Error level marks an error output.
	Error
)

// Factory defines a utility to create new loggers and set the log level threshold.
type Factory interface {

	//New creates a new logger.
	New() Logger

	// Threshold sets the minimum log level threshold for messages to be output.
	Threshold(min LogLevel)
}

// Logger defines the logger interface.
type Logger interface {

	// Level creates a new logger instance from the current one setting its log level to the value supplied.
	Level(lvl LogLevel) Logger

	// Field creates a new logger instance from the current one adding a new field value.
	Field(k string, v interface{}) Logger

	// Fields creates a new logger instance from the current one adding a collection of field values.
	Fields(data map[string]interface{}) Logger

	// WithError creates a new logger instance from the current one adding an error
	// and setting the level to Error.
	WithError(err error) Logger

	// Log outputs the log structure along with a message if the logger level is above or matching
	// the threshold set in the factory.
	Log(msg string) Tracer

	// Logf outputs the log structure along with a formatted message if the logger level is above or
	// matching the threshold set in the factory.
	Logf(format string, v ...interface{}) Tracer
}
