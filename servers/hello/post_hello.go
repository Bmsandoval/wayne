package hello

import (
	"context"
	"github.com/bmsandoval/wayne/db/models"
)

// PostHello implements hello.GreeterServer
func (s *Server) PostHello(ctx context.Context, in *PostHelloRequest) (*PostHelloReply, error) {
	greeting := models.Greetings{
		Value:     in.GetGreeting(),
	}

	result, err := s.Bundle.HelloSvc.Create(greeting)

	if err != nil {
		return &PostHelloReply{
			Greeting: "",
		}, err
	}

	return &PostHelloReply{
		Greeting: result.Value,
	}, nil
}
