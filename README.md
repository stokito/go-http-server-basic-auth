# go-http-server-basic-auth

Simplest basic authenticator for HTTP server.
Also contains CORS handler to allow access from everywhere. You must be careful and check auth.
A useful addition is a Recovery handler to handle panic and print it into log.

## Install

    go get -u github.com/stokito/go-http-server-basic-auth

## Usage

See [example](./examples) folder

## License

[0BSD](https://opensource.org/licenses/0BSD) (similar to Public Domain)

## See also

* Most popular library https://github.com/abbot/go-http-auth
* Some popular gist that many users just copy and paste https://gist.github.com/elithrar/9146306 It's mentioned in the first googled StackOverflow topic :) It's written by the @elithrar
* Many others libs but none seems good to me https://pkg.go.dev/search?q=basic+auth
* One of them looks easy to use https://github.com/99designs/basicauth-go/blob/2a93ba0f464d/basicauth.go
* Prometheus also has basic auth handler with BCrypt support https://github.com/prometheus/exporter-toolkit/blob/master/web/handler.go#L105
* https://github.com/gorilla/handlers also has a recovery and CORS handlers mut more complicated and slow
