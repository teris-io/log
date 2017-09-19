// Copyright 2017 teris.io. License MIT

package std_test

import (
	"code.teris.io/go/log"
	"testing"
	"time"
	"code.teris.io/go/log/std"
	"fmt"
	stdlog "log"
	"os"
	"github.com/pkg/errors"
	"github.com/fatih/color"
)

func init() {
	std.Use()
	std.SetFormatter(func(start time.Time, lvl log.Level, msg string, ctx interface{}) string {
		lvlstr := ""
		switch lvl {
		case log.Debug:
			lvlstr = fmt.Sprintf(" [%s]", color.CyanString("DBG"))
		case log.Info:
			lvlstr = fmt.Sprintf(" [%s]", color.GreenString("INF"))
		case log.Error:
			lvlstr = fmt.Sprintf(" [%s]", color.RedString("ERR"))
		}
		return fmt.Sprintf("%s%s %s %v", start.Format("02/01 15:04:05.999"), lvlstr, msg, ctx)
	})
	stdlog.SetOutput(os.Stdout)
}

func TestLogger_Log(t *testing.T) {
	tracer := log.With("context", "Foo").Log("something happened")
	time.Sleep(time.Second)
	tracer.Trace()
	time.Sleep(time.Second)
}

func TestLogger_LogError(t *testing.T) {
	origerr := fmt.Errorf("some error")
	err := errors.Wrap(origerr, "wrapped")
	tracer := log.With("context", "Foo").WithError(err).Log("something happened")
	time.Sleep(time.Second)
	tracer.Trace()
	time.Sleep(time.Second)
}
