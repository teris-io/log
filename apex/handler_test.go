// Copyright 2017 teris.io. License MIT

package apex_test

import (
	"code.teris.io/go/log"
	"code.teris.io/go/log/apex"
	alog "github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"os"
	"testing"
	"time"
)

func TestLogger_Log(t *testing.T) {
	apex.Use()
	alog.SetHandler(text.New(os.Stderr))
	defer log.With("context", "Foo").WithLevel(log.Error).Log("something happened").Trace()
	time.Sleep(time.Second * 5)
}
