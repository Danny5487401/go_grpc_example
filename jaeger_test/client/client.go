package main
import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"google.golang.org/grpc"

	"go_test_project/jaeger_test/otgrpc"
	"go_test_project/jaeger_test/proto"
)

func main()  {
	cfg := jaegercfg.Configuration{
		// 采样
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
			LocalAgentHostPort:"81.68.197.3:6831",
		},
		ServiceName: "danny_test",
	}
	// 打印日志
	jLogger := jaegerlog.StdLogger
	tracer,closer,err :=cfg.NewTracer(jaegercfg.Logger(jLogger))
	if err != nil{
		panic(err)
	}
	defer closer.Close()

	// 注册全局tracer，方便调用--->opentracing.StartSpan, 不需要一直传tracer
	opentracing.SetGlobalTracer(tracer)

	// 方式一：使用实例tracer
	//conn,err := grpc.Dial("127.0.0.1:9000",grpc.WithInsecure(),
	//	grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer))))

	// 方式二：使用全局tracer
	conn,err := grpc.Dial("127.0.0.1:9000",grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))

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
