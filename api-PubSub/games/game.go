package games

import (
	context "context"
	"log"
)

type Server struct {
}

func (s *Server) Play(ctx context.Context, in *ServerRequest) (*ServerResponse, error) {
	log.Printf("Receive message body from client: %s", in.Request)
	return &ServerResponse{}, nil
}
