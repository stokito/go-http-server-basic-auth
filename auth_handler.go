// SPDX-License-Identifier: 0BSD
package basicauth

import (
	"crypto/subtle"
	"net/http"
	"strings"
)

// HeaderSubscriberId Authenticated username or email
const HeaderSubscriberId = "X-Subscriber-ID"

// AuthHandlerWrapper Wrapper for a http.Handler that adds basic auth
type AuthHandlerWrapper struct {
	Handler     http.Handler
	Credentials map[string]string
	Realm       string
	wwwAuthHdr  string
	IgnorePath  []string
}

func NewAuthHandlerWrapper(handler http.Handler, credentials map[string]string, realm string, ignorePath []string) *AuthHandlerWrapper {
	return &AuthHandlerWrapper{
		Handler:     handler,
		Credentials: credentials,
		Realm:       realm,
		wwwAuthHdr:  `Basic realm="` + realm + `", charset="UTF-8"`,
		IgnorePath:  ignorePath,
	}
}

func (bah *AuthHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// The OPTIONS must be always allowed for CORS requests
	if r.Method != http.MethodOptions && !bah.pathIsIgnored(r.URL.Path) {
		authorized := false
		username, password, ok := r.BasicAuth()
		if ok {
			userPassword := bah.Credentials[username]
			if userPassword != "" && subtle.ConstantTimeCompare([]byte(password), []byte(userPassword)) == 1 {
				authorized = true
				r.Header.Set(HeaderSubscriberId, username)
			}
		}
		if !authorized {
			w.Header().Set("WWW-Authenticate", bah.wwwAuthHdr)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}
	bah.Handler.ServeHTTP(w, r)
}

func (bah AuthHandlerWrapper) pathIsIgnored(path string) bool {
	for _, ignoredPath := range bah.IgnorePath {
		if strings.HasPrefix(path, ignoredPath) {
			return true
		}
	}
	return false
}
