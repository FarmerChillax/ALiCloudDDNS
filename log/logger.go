package log

import (
	stdlog "log"
	"sync"
)

var (
	log  = stdlog.Logger{}
	once = sync.Once{}
)

func New() *stdlog.Logger {
	once.Do(func() {})
	return &stdlog.Logger{}
}
