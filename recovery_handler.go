package basicauth

import (
	"bytes"
	"log"
	"net/http"
	"runtime/debug"
)

type RecoveryHandlerWrapper struct {
	Handler  http.Handler
	ErrorLog *log.Logger
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
			h.ErrorLog.Printf("<3> FAIL: %s: %s\n", panicErr, stackTrace)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()
	h.Handler.ServeHTTP(w, r)
}
