package basicauth

import "net/http"

type CorsHandlerWrapper struct {
	Handler http.Handler
}

func (h *CorsHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		corsAllowHeaders     = "authorization"
		corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
		corsAllowOrigin      = "*"
		corsAllowCredentials = "true"
	)
	h.Handler.ServeHTTP(w, r)
	w.Header().Set("Access-Control-Allow-Credentials", corsAllowCredentials)
	w.Header().Set("Access-Control-Allow-Headers", corsAllowHeaders)
	w.Header().Set("Access-Control-Allow-Methods", corsAllowMethods)
	w.Header().Set("Access-Control-Allow-Origin", corsAllowOrigin)
}
