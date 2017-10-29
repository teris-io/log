// Copyright (c) 2017. Oleg Sklyar & teris.io. All rights reserved.
// See the LICENSE file in the project root for licensing information.

package std_test

import (
	"fmt"
	stdlog "log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/teris-io/log"
	"github.com/teris-io/log/std"
)

type strwriter struct {
	entries []string
}

func (sw *strwriter) Write(p []byte) (n int, err error) {
	sw.entries = append(sw.entries, string(p))
	return len(p), nil
}

func TestLoggerStaticLogsDirectly(t *testing.T) {
	defer std.Use(os.Stderr, log.UnsetLevel, std.DefaultFmtFun)
	w := &strwriter{}
	std.Use(w, log.UnsetLevel, std.DefaultFmtFun)
	log.Log("logs directly")
	assert.True(t, strings.Contains(w.entries[0], " logs directly"))
}

func TestLoggerStaticLogsFormattedDirectly(t *testing.T) {
	defer std.Use(os.Stderr, log.UnsetLevel, std.DefaultFmtFun)
	w := &strwriter{}
	std.Use(w, log.UnsetLevel, std.DefaultFmtFun)
	log.Logf("%s %s", "logs", "directly")
	assert.True(t, strings.Contains(w.entries[0], " logs directly"))
}

func TestLoggerStaticNewLoggerOnNew(t *testing.T) {
	defer std.Use(os.Stderr, log.UnsetLevel, std.DefaultFmtFun)
	w := &strwriter{}
	std.Use(w, log.UnsetLevel, std.DefaultFmtFun)
	logger := log.New()
	logger.Log("new logger")
	log.Log("logs directly")
	assert.True(t, strings.Contains(w.entries[0], " new logger"))
	assert.True(t, strings.Contains(w.entries[1], " logs directly"))
}

func TestLoggerStaticNewLoggerOnLevel(t *testing.T) {
	defer std.Use(os.Stderr, log.UnsetLevel, std.DefaultFmtFun)
	w := &strwriter{}
	std.Use(w, log.UnsetLevel, std.DefaultFmtFun)
	logger := log.Level(log.DebugLevel)
	logger.Log("new logger")
	log.Log("logs directly")
	assert.True(t, strings.Contains(w.entries[0], " DBG new logger"))
	assert.True(t, strings.Contains(w.entries[1], " logs directly"))
}

func TestLoggerStaticNewLoggerOnWith(t *testing.T) {
	defer std.Use(os.Stderr, log.UnsetLevel, std.DefaultFmtFun)
	w := &strwriter{}
	std.Use(w, log.UnsetLevel, std.DefaultFmtFun)
	logger := log.Field("ctx", "context")
	logger.Log("new logger")
	log.Log("logs directly")
	assert.True(t, strings.Contains(w.entries[0], " new logger {ctx: context}"))
	assert.True(t, strings.Contains(w.entries[1], " logs directly"))
}

func TestLoggerStaticNewLoggerOnFields(t *testing.T) {
	defer std.Use(os.Stderr, log.UnsetLevel, std.DefaultFmtFun)
	w := &strwriter{}
	std.Use(w, log.UnsetLevel, std.DefaultFmtFun)
	logger := log.Fields(map[string]interface{}{
		"key1": 25,
		"key2": "value2",
	})
	logger.Log("new logger")
	log.Log("logs directly")
	// hash maps are unsorted
	assert.True(t, strings.Contains(w.entries[0], "key1: 25") && strings.Contains(w.entries[0], "key2: value2"))
	assert.True(t, strings.Contains(w.entries[1], " logs directly"))
}

func TestLoggerStaticNewLoggerOnError(t *testing.T) {
	defer std.Use(os.Stderr, log.UnsetLevel, std.DefaultFmtFun)
	w := &strwriter{}
	std.Use(w, log.UnsetLevel, std.DefaultFmtFun)
	logger := log.Error(errors.Wrap(fmt.Errorf("inner"), "outer"))
	logger.Log("new logger")
	log.Log("logs directly")
	assert.True(t, strings.Contains(w.entries[0], " ERR new logger {error: outer: inner}, {source: "))
	assert.True(t, strings.Contains(w.entries[1], " logs directly"))
}

func TestLoggerFactoryNewLoggerOnNew(t *testing.T) {
	w := &strwriter{}
	factory := std.NewFactory(stdlog.New(w, "", 0), log.UnsetLevel, std.DefaultFmtFun)
	logger1 := factory.New()
	logger2 := factory.New()
	logger2.Log("logger2")
	logger1.Log("logger1")
	assert.True(t, strings.Contains(w.entries[0], " logger2"))
	assert.True(t, strings.Contains(w.entries[1], " logger1"))
}

func TestLoggerChainingCreatesNewLoggers(t *testing.T) {
	w := &strwriter{}
	factory := std.NewFactory(stdlog.New(w, "", 0), log.UnsetLevel, std.DefaultFmtFun)
	logger1 := factory.New()
	logger2 := logger1.Level(log.DebugLevel)
	logger3 := logger2.Field("ctx", "context")
	logger4 := logger3.Level(log.InfoLevel)
	logger5 := logger4.Error(fmt.Errorf("failed %s", "badly"))
	logger5.Log("5th")
	logger4.Log("4th")
	logger3.Log("3rd")
	logger2.Log("2nd")
	logger1.Log("1st")
	assert.True(t, strings.Contains(w.entries[0], " ERR 5th {error: failed badly}, {ctx: context}"))
	assert.True(t, strings.Contains(w.entries[1], " INF 4th {ctx: context}"))
	assert.True(t, strings.Contains(w.entries[2], " DBG 3rd {ctx: context}"))
	assert.True(t, strings.Contains(w.entries[3], " DBG 2nd"))
	assert.True(t, strings.Contains(w.entries[4], " 1st"))
}

func TestLoggerLevelsBelowMinFilteredOut(t *testing.T) {
	w := &strwriter{}
	factory := std.NewFactory(stdlog.New(w, "", 0), log.InfoLevel, std.DefaultFmtFun)
	logger1 := factory.New()
	logger2 := logger1.Level(log.DebugLevel)
	logger3 := logger2.Field("ctx", "context")
	logger4 := logger3.Level(log.InfoLevel)
	logger5 := logger4.Error(fmt.Errorf("failed %s", "badly"))
	logger5.Log("5th")
	logger4.Log("4th")
	logger3.Log("3rd")
	logger2.Log("2nd")
	logger1.Log("1st")
	assert.Equal(t, 2, len(w.entries))
	assert.True(t, strings.Contains(w.entries[0], " ERR 5th {error: failed badly}, {ctx: context}"))
	assert.True(t, strings.Contains(w.entries[1], " INF 4th {ctx: context}"))
}

func TestLoggerTrace(t *testing.T) {
	w := &strwriter{}
	factory := std.NewFactory(stdlog.New(w, "", 0), log.DebugLevel, std.DefaultFmtFun)
	logger := factory.New().Level(log.DebugLevel).Field("key", "value").Log("start")
	time.Sleep(100 * time.Millisecond)
	logger.Trace()
	assert.Equal(t, 2, len(w.entries))
	assert.True(t, strings.Contains(w.entries[0], "DBG start {key: value}"))
	assert.True(t, strings.Contains(w.entries[1], "DBG traced {duration: 0"))

}
