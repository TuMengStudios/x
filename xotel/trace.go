package xotel

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type Tracer interface {
	CapTrace() gin.HandlerFunc
}

const (
	defaultTracerName = "http-tracing"
)

type NullTracer struct {
	tp          *tracesdk.TracerProvider
	props       propagation.TextMapPropagator
	serviceName string
}

func NewNullTracer(serviceName string, tp *tracesdk.TracerProvider) Tracer {
	otel.SetTextMapPropagator(propagation.TraceContext{})
	if serviceName == "" {
		serviceName = defaultTracerName
	}
	return &NullTracer{
		tp:          tp,
		props:       otel.GetTextMapPropagator(),
		serviceName: serviceName,
	}
}

//func (c *NullTracer) CapTrace() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		// 如果已经有 trace, 那就不用注入
//		if v := trace.SpanFromContext(ctx.Request.Context()).SpanContext(); v.IsValid() {
//			ctx.Header("RequestId", v.TraceID().String())
//			c.props.Inject(ctx.Request.Context(), propagation.HeaderCarrier(ctx.Writer.Header()))
//			ctx.Next()
//			return
//		}
//
//		extractCtx := c.props.Extract(ctx.Request.Context(), propagation.HeaderCarrier(ctx.Request.Header))
//
//		newCtx, span := c.tp.Tracer(c.serviceName).Start(extractCtx, "http-entry")
//		defer span.End()
//		ctx.Request = ctx.Request.WithContext(newCtx)
//		if v := trace.SpanFromContext(ctx.Request.Context()).SpanContext(); v.IsValid() {
//			ctx.Header("RequestId", v.TraceID().String())
//			c.props.Inject(ctx.Request.Context(), propagation.HeaderCarrier(ctx.Writer.Header()))
//		}
//
//		ctx.Next()
//	}
//}

func (c *NullTracer) CapTrace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 如果已经有 trace, 那就不用注入
		if v := trace.SpanFromContext(ctx.Request.Context()).SpanContext(); v.IsValid() {
			ctx.Header("RequestId", v.TraceID().String())
			ctx.Next()
			return
		}

		extractCtx := c.props.Extract(ctx.Request.Context(), propagation.HeaderCarrier(ctx.Request.Header))

		newCtx, span := c.tp.Tracer(c.serviceName).Start(extractCtx, "http-entry")
		defer span.End()
		ctx.Request = ctx.Request.WithContext(newCtx)
		if v := trace.SpanFromContext(ctx.Request.Context()).SpanContext(); v.IsValid() {
			ctx.Header("RequestId", v.TraceID().String())
		}

		ctx.Next()
	}
}
