package main
import (
	"context"
	"fmt"
	"go_test_project/grpc_test/proto"
	"google.golang.org/grpc"
)

func main()  {
	conn,err := grpc.Dial("127.0.0.1:9000",grpc.WithInsecure())
	if err != nil{
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	r,err := c.SayHello(context.Background(),&proto.HelloRequest{
		Name: "danny",
	})
	if err != nil{
		panic(err)
	}
	fmt.Println(r.Message)
}
