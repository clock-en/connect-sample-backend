package main

import (
	"github.com/clock-en/connect-sample-backend/pbgen/submodules/protobuf/greet/v1/greetv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"

	"github.com/clock-en/connect-sample-backend/internal/greet"
	"github.com/rs/cors"
)

func main() {
	greeter := &greet.GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	log.Println(path)
	mux.Handle(path, handler)

	// TODO: 接続確認のため、雑な設定
	corsHandler := cors.AllowAll().Handler(h2c.NewHandler(mux, &http2.Server{}))

	http.ListenAndServe(
		":8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		corsHandler,
	)
}
