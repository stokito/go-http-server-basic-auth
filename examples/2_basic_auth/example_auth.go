package main

import (
	basicauth "github.com/stokito/go-http-server-basic-auth"
	"net/http"
)

func main() {
	credentials := map[string]string{"admin": "secret"}
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/test", handleReq)
	handler := basicauth.NewAuthHandlerWrapper(
		serveMux,
		credentials,
		"Admin Dashboard",
		[]string{"/robots.txt", "/favicon.ico", "/.well-known/"},
	)
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	server.ListenAndServe()
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		username := r.Header.Get(basicauth.HeaderSubscriberId)
		body := "Hello " + username
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(body))
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
