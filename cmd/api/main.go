package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/clock-en/connect-sample-backend/internal/greet"
	"github.com/rs/xid"
)

func main() {
	xid := xid.New()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logText := fmt.Sprintf("XID: %s", xid)
		log.Println(logText)
		fmt.Fprintf(w, greet.Hello())
	})

	http.ListenAndServe(":8000", nil)
}
