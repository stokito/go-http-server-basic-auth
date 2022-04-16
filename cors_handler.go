package basicauth

import "net/http"

type CorsHandlerWrapper struct {
	Handler http.Handler
}

func (h *CorsHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	header := w.Header()
	header.Set("Access-Control-Allow-Origin", origin)
	header.Set("Access-Control-Allow-Credentials", "true")
	if r.Method == http.MethodOptions {
		header.Set("Access-Control-Allow-Methods", "HEAD,GET,POST,PUT,DELETE,OPTIONS")
		header.Set("Access-Control-Allow-Headers", "Accept,Authorization,Date,Content-Type,Origin")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	h.Handler.ServeHTTP(w, r)
}
