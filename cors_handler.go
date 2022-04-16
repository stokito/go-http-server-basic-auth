package basicauth

import "net/http"

type CorsHandlerWrapper struct {
	Handler http.Handler
}

func (h *CorsHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		corsAllowOrigin      = "*"
		corsExposeHeaders    = "Date"
		corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
		corsAllowHeaders     = "Accept,Authorization,Date,Content-Type,Origin"
		corsAllowCredentials = "true"
	)
	w.Header().Set("Access-Control-Allow-Origin", corsAllowOrigin)
	w.Header().Set("Access-Control-Expose-Headers", corsExposeHeaders)
	w.Header().Set("Access-Control-Allow-Methods", corsAllowMethods)
	w.Header().Set("Access-Control-Allow-Headers", corsAllowHeaders)
	w.Header().Set("Access-Control-Allow-Credentials", corsAllowCredentials)
	h.Handler.ServeHTTP(w, r)
}
