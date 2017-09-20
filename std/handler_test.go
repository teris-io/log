// Copyright 2017 teris.io. License MIT

package std_test

import (
	"code.teris.io/go/log"
	"code.teris.io/go/log/std"
	"fmt"
	"github.com/pkg/errors"
	stdlog "log"
	"os"
	"testing"
	"time"
)

func init() {
	std.Use()
	stdlog.SetOutput(os.Stdout)
}

func TestLogger_Log(t *testing.T) {
	logger := log.With("context", "Foo")
	logger.Log("something happened")
	logger.WithLevel(log.Debug).Log("some degub message")
	logger.WithLevel(log.Info).Log("some info message")
	logger.WithLevel(log.Error).Log("some error message")
}

func TestLogger_LogError(t *testing.T) {
	origerr := fmt.Errorf("some error")
	err := errors.Wrap(origerr, "wrapped")
	tracer := log.With("context", "Foo").WithError(err).Log("something happened")
	time.Sleep(time.Second)
	tracer.Trace()
	time.Sleep(time.Second)
}
