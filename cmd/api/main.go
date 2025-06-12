package main

import (
	"github.com/clock-en/connect-sample-backend/pbgen/submodules/protobuf/greet/v1/greetv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"

	"github.com/clock-en/connect-sample-backend/internal/greet"
)

func main() {
	greeter := &greet.GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	log.Println(path)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
