package greet

import (
	"context"
	"fmt"
	v1greet "github.com/clock-en/connect-sample-backend/pbgen/submodules/protobuf/v1/greet"
	"log"

	"connectrpc.com/connect"
)

type GreetServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[v1greet.GreetRequest],
) (*connect.Response[v1greet.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&v1greet.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
