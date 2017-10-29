// Copyright (c) 2017. Oleg Sklyar & teris.io. All rights reserved.
// See the LICENSE file in the project root for licensing information.

package log

// LoggerLevel defines level markers for log entries.
type LoggerLevel int

const (
	// UnsetLevel should not be output by logger implementation.
	UnsetLevel = iota - 2
	// DebugLevel marks detailed output for design purposes.
	DebugLevel
	// InfoLevel is the default log output marker.
	InfoLevel
	// ErrorLevel marks an error output.
	ErrorLevel
)

// Factory defines a utility to create new loggers and set the log level threshold.
type Factory interface {

	//New creates a new logger.
	New() Logger

	// Threshold sets the minimum logger level threshold for messages to be output.
	Threshold(min LoggerLevel)
}

// Logger defines the logger interface.
type Logger interface {

	// Level creates a new logger instance from the current one setting its log level to the value supplied.
	Level(lvl LoggerLevel) Logger

	// Field creates a new logger instance from the current one adding a new field value.
	Field(k string, v interface{}) Logger

	// Fields creates a new logger instance from the current one adding a collection of field values.
	Fields(data map[string]interface{}) Logger

	// Error creates a new logger instance from the current one adding an error
	// and setting the level to ErrorLevel.
	Error(err error) Logger

	// Log outputs the log structure along with a message if the logger level is above or matching
	// the threshold set in the factory.
	Log(msg string) Tracer

	// Logf outputs the log structure along with a formatted message if the logger level is above or
	// matching the threshold set in the factory.
	Logf(format string, v ...interface{}) Tracer
}
