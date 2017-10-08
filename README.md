# Structured leveled log interface

Package `log` decouples the logger backend from your application. It defines
a simple, lightweight and comprehensive `Logger` and `Factory` interfaces which 
can be used through your applications without any knowledge of the particular
implemeting backend and can be configured at the application wiring point to
bind a particular backend, such as Go's standard logger, `apex/log`, `logrus`, 
with ease.

To complement the facade, the package `code.teris.io/util/log/std` provides an 
implementation using the standard Go logger. The default log formatter for
this implementation uses colour coding for log levels and logs the date
leaving out the month and the year on the timestamp. However, the formatter
is fully configurable.

Similarly, the package `code.teris.io/util/log/apex` provides and implementation 
using the `apex/log` logger backend.


## Interface details

The `Logger` interface defines a facade for a structured leveled log: 

```go
type Logger interface {
	Level(lvl LogLevel) Logger
	Field(k string, v interface{}) Logger
	Fields(data map[string]interface{}) Logger
	WithError(err error) Logger
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
	std.Use(os.Stderr, log.Info, std.DefaultFmtFun)
}

// elsewhere	
logger := log.Level(log.Info).Field("key", "value")
logger.Log("message")
```

and dynamically by always going via a factory:

```go
factory := std.NewFactory(os.Stderr, log.Info, std.DefaultFmtFun)
logger := factory.Level(log.Info).Field("key", "value")
logger.Log("message")
```

By default a NoOp (no-operation) implementation is bound to the static factory.

## Tracing

To simplify debugging with execution time tracing, the `Log` and `Logf` methods
return a tracer that can be used to measure and log the execution time:

```go
logger := log.Level(log.Debug).Field("key", "value")

defer logger.Log("start").Trace()
// code to trace the execution time of
```

The above code snippet would output two log entries (provided the threshold permits)
the selected `Debug` level (her for the default formatter of the `std` logger):

	08 16:31:42.023798 DBG start {key: value}
	08 16:31:45.127619 DBG traced {duration: 3.103725832}, {key: value}

# License and copyright

    Copyright (c) 2017 Oleg Sklyar & teris.io

    MIT License

Permission is hereby granted, free of charge, to any person obtaining a copy of this 
software and associated documentation files (the "Software"), to deal in the Software 
without restriction, including without limitation the rights to use, copy, modify, 
merge, publish, distribute, sublicense, and/or sell copies of the Software, and to 
permit persons to whom the Software is furnished to do so, subject to the following 
conditions:

The above copyright notice and this permission notice shall be included in all copies 
or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, 
INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A 
PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT 
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION 
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE 
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
