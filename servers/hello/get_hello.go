package hello

import (
	"context"
)

// GetHello implements hello.GreeterServer
func (s *Server) GetHello(ctx context.Context, in *GetHelloRequest) (*GetHelloReply, error) {
	greetings, err := s.Bundle.HelloSvc.Get()
	if err != nil {
		return &GetHelloReply{
			Greetings: nil,
		}, nil
	}

	var greetingsStrings []string

	for _, greeting := range greetings {
		greetingsStrings = append(greetingsStrings, greeting.Value)
	}

	return &GetHelloReply{
		Greetings: greetingsStrings,
	}, nil
}
