package service

import (
	"context"

	v1 "industry/api/industry/v1"
	"industry/internal/biz"
)

type IndustryServer struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

func NewGreeterService(uc *biz.GreeterUsecase) *IndustryServer {
	return &IndustryServer{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *IndustryServer) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
