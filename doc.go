// Copyright (c) 2017. Oleg Sklyar & teris.io. All rights reserved.
// See the LICENSE file in the project root for licensing information.

// Package log defines the Logger interface for a structured leveled log, the Factory
// interface to create instances of the Logger and the Tracer interface used to trace
// and log the execution time since last Log.
//
// The package further defines three log levels differentiating between the (normally
// hidden) Debug, (default) Info and (erroneous) Error.
//
// The log can be used both statically by binding a particular logger factory, as in
//
//     std.Use(os.Stderr, log.Info, std.DefaultFmtFun)
//     logger := log.Level(log.Info).Field("key", "value")
//     logger.Log("message")
//
// and dynamically by always going via a factory, as in
//
//     factory := std.NewFactory(os.Stderr, log.Info, std.DefaultFmtFun)
//     logger := factory.Level(log.Info).Field("key", "value")
//     logger.Log("message")
package log
