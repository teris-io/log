[![Build status][buildimage]][build] [![Coverage][codecovimage]][codecov] [![GoReportCard][cardimage]][card] [![API documentation][docsimage]][docs]

# Structured log interface

Package `log` provides the separation of the logging interface from its 
implementation and decouples the logger backend from your application. It defines
a simple, lightweight and comprehensive `Logger` and `Factory` interfaces which 
can be used through your applications without any knowledge of the particular
implemeting backend and can be configured at the application wiring point to
bind a particular backend, such as Go's standard logger, `apex/log`, `logrus`, 
with ease.

To complement the facade, the package `github.com/teris-io/log/std` provides an 
implementation using the standard Go logger. The default log formatter for
this implementation uses colour coding for log levels and logs the date
leaving out the month and the year on the timestamp. However, the formatter
is fully configurable.

Similarly, the package `github.com/teris-io/log/apex` provides and implementation 
using the `apex/log` logger backend.


## Interface details

The `Logger` interface defines a facade for a structured leveled log: 

```go
type Logger interface {
	Level(lvl LogLevel) Logger
	Field(k string, v interface{}) Logger
	Fields(data map[string]interface{}) Logger
	Error(err error) Logger
	Log(msg string) Tracer
	Logf(format string, v ...interface{}) Tracer
}
```

The `Factory` defines a facade for the creation of logger instances and setting the
log output threshold for newly created instances:

```go
type Factory interface {
	New() Logger
	Threshold(min LogLevel)
}
```

The package further defines three log levels differentiating between the (normally hidden) 
`Debug`, (default) `Info` and (erroneous) `Error`.


## Usage

The log can be used both statically by binding a particular logger factory:

```go
func init() {
	std.Use(os.Stderr, log.InfoLevel, std.DefaultFmtFun)
}

// elsewhere	
logger := log.Level(log.InfoLevel).Field("key", "value")
logger.Log("message")
```

and dynamically by always going via a factory:

```go
factory := std.NewFactory(os.Stderr, log.InfoLevel, std.DefaultFmtFun)
logger := factory.Level(log.InfoLevel).Field("key", "value")
logger.Log("message")
```

By default a NoOp (no-operation) implementation is bound to the static factory.

## Tracing

To simplify debugging with execution time tracing, the `Log` and `Logf` methods
return a tracer that can be used to measure and log the execution time:

```go
logger := log.Level(log.DebugLevel).Field("key", "value")

defer logger.Log("start").Trace()
// code to trace the execution time of
```

The above code snippet would output two log entries (provided the threshold permits)
the selected `Debug` level (her for the default formatter of the `std` logger):

	08 16:31:42.023798 DBG start {key: value}
	08 16:31:45.127619 DBG traced {duration: 3.103725832}, {key: value}

### License and copyright

	Copyright (c) 2017. Oleg Sklyar and teris.io. MIT license applies. All rights reserved.


[build]: https://travis-ci.org/teris-io/log
[buildimage]: https://travis-ci.org/teris-io/log.svg?branch=master

[codecov]: https://codecov.io/github/teris-io/log?branch=master
[codecovimage]: https://codecov.io/github/teris-io/log/coverage.svg?branch=master

[card]: http://goreportcard.com/report/teris-io/log
[cardimage]: https://goreportcard.com/badge/github.com/teris-io/log

[docs]: https://godoc.org/github.com/teris-io/log
[docsimage]: http://img.shields.io/badge/godoc-reference-blue.svg?style=flat
