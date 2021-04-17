package main

import (
	"context"
	"net"

	"go_test_project/jaeger_test/proto"
	"google.golang.org/grpc"
)

type Server struct {

}

func (s *Server)SayHello(ctx context.Context,request *proto.HelloRequest)  (*proto.HelloReply,error){
	return &proto.HelloReply{
		Message:"hello," + request.Name,
	},nil
}

func main()  {
	g := grpc.NewServer()

	proto.RegisterGreeterServer(g,&Server{})
	lis,err := net.Listen("tcp","0.0.0.0:9000")
	if err != nil{
		panic("failed to listen:" + err.Error())
	}
	_ = g.Serve(lis)
}
