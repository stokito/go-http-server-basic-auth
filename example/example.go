package main

import (
	basicauth "github.com/stokito/go-http-server-basic-auth"
	"net/http"
)

func main() {
	credentials := map[string]string{"admin": "secret"}
	serveMux := http.NewServeMux()
	serveMux.Handle("/", http.FileServer(http.Dir("./www")))
	serveMux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			username := r.Header.Get(basicauth.HeaderSubscriberId)
			body := "Hello " + username
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(body))
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})
	server := &http.Server{
		Addr: ":8000",
		Handler: &basicauth.AuthHandlerWrapper{
			Handler:     serveMux,
			Realm:       "Admin Dashboard",
			Credentials: credentials,
			IgnorePath:  []string{"/robots.txt", "/favicon.ico", "/.well-known/"},
		},
	}
	server.ListenAndServe()
}
