package server

import (
	"context"
	"fmt"
	"testing"

	pb "github.com/nikunjy/go/protos/hello"
	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	srv := &Server{}
	req := &pb.GreetingRequest{Name: "Test"}
	ctx := context.Background()
	resp, _ := srv.Greet(ctx, req)

	fmt.Println(resp.Greeting)
	require.EqualValues(t, "Hello Test", resp.Greeting)
}
