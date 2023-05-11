package basicauth

import (
	"bytes"
	"net/http"
	"runtime/debug"
)

type LogPrinterFunc func(format string, v ...interface{})

// RecoveryHandlerWrapper Middleware to catch and log panics
type RecoveryHandlerWrapper struct {
	Handler  http.Handler
	ErrorLog LogPrinterFunc
}

func (h *RecoveryHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// catch panic
	defer func() {
		panicErr := recover()
		if panicErr != nil {
			stackTrace := debug.Stack()
			// linearize stacktrace
			stackTrace = bytes.Replace(stackTrace, []byte("\n"), []byte("|"), -1)
			// the "<3>" is ERR in syslog
			h.ErrorLog("<3> FAIL: %s: %s\n", panicErr, stackTrace)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()
	h.Handler.ServeHTTP(w, r)
}
