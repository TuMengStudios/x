package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/context"

	"github.com/TuMengStudios/x/xotel"
)

func main() {
	var wg = sync.WaitGroup{}
	var engine = gin.New()
	var tp = tracesdk.NewTracerProvider(tracesdk.WithSampler(tracesdk.NeverSample()))
	engine.Use(xotel.NewNullTracer("test", tp).CapTrace())
	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	})
	// 随机监听一个端口
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	var port = listener.Addr().(*net.TCPAddr).Port

	ctx, cancelFunc := context.WithCancel(context.Background())
	wg.Add(2)
	go func() {
		defer wg.Done()
		cancelFunc()

		if err := engine.RunListener(listener); err != nil {
			fmt.Printf("error %v\n", err)
		}

	}()
	go func() {
		defer wg.Done()
		<-ctx.Done()
		time.Sleep(time.Second)
		resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%v/", port))
		if err != nil {
			panic(err)
		}

		if resp.Body != nil {
			defer resp.Body.Close()
		}

		mapPropagator := otel.GetTextMapPropagator()
		extractCtx := mapPropagator.Extract(context.Background(), propagation.HeaderCarrier(resp.Header))
		t := trace.SpanFromContext(extractCtx).SpanContext()
		fmt.Println(t.IsValid(), t.TraceID().String())
		fmt.Println(resp.Header.Get("RequestId"))
		listener.Close()
	}()
	wg.Wait()

}
