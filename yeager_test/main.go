package main

import (
	"time"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
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
	span := tracer.StartSpan("go-grpc-web")
	time.Sleep(time.Second)
	defer span.Finish()
}
