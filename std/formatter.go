// Copyright (c) 2017. Oleg Sklyar & teris.io. All rights reserved.
// See the LICENSE file in the project root for licensing information.

package std

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/teris-io/log"
)

// FmtFun defines the log formatter type for the logger implementation based on the Go standard log.
type FmtFun func(start time.Time, level log.LoggerLevel, msg string, fields []Field) string

// DefaultFmtFun defines the default formatter.
var DefaultFmtFun = func(start time.Time, lvl log.LoggerLevel, msg string, fields []Field) string {
	timestr := color.CyanString(start.Format("02 15:04:05.000000"))
	lvlstr := ""
	switch lvl {
	case log.DebugLevel:
		lvlstr = fmt.Sprintf(" %s", color.New(color.Bold, color.FgYellow).Sprint("DBG"))
	case log.InfoLevel:
		lvlstr = fmt.Sprintf(" %s", color.New(color.Bold, color.FgGreen).Sprint("INF"))
	case log.ErrorLevel:
		lvlstr = fmt.Sprintf(" %s", color.New(color.Bold, color.FgRed).Sprint("ERR"))
	}
	fieldstr := ""
	for _, f := range fields {
		if fieldstr != "" {
			fieldstr += ", "
		}
		fieldstr += fmt.Sprintf("{%s: %v}", color.New(color.Bold).Sprint(f.Name), f.Value)
	}
	return fmt.Sprintf("%s%s %s %s", timestr, lvlstr, msg, fieldstr)
}
