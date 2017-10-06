// Copyright 2017 teris.io. License MIT

package std

import (
	"code.teris.io/go/log"
	"fmt"
	"github.com/fatih/color"
	"time"
)

var formatter = func(start time.Time, lvl log.Level, fields []Field, msg string) string {
	timestr := color.CyanString(start.Format("02 15:04:05.000000"))
	lvlstr := ""
	switch lvl {
	case log.Debug:
		lvlstr = fmt.Sprintf(" %s", color.New(color.Bold, color.FgYellow).Sprint("DBG"))
	case log.Info:
		lvlstr = fmt.Sprintf(" %s", color.New(color.Bold, color.FgGreen).Sprint("INF"))
	case log.Error:
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
