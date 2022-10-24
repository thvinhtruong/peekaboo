package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// startSpanParent starts a new span from a parent span.
func startSpanParent(parent opentracing.SpanContext, ctx *gin.Context) opentracing.Span {
	options := []opentracing.StartSpanOption{
		opentracing.Tag{Key: "component", Value: "gin"},
		opentracing.Tag{Key: "http.url", Value: ctx.Request.URL.String()},
		opentracing.Tag{Key: "http.method", Value: ctx.Request.Method},
		opentracing.Tag{Key: "http.remote_ip", Value: ctx.ClientIP()},
	}
	if parent != nil {
		options = append(options, opentracing.ChildOf(parent))
	}
	return opentracing.StartSpan(ctx.Request.URL.Path, options...)
}

// startSpanHeader starts span based on header information, looking for the parent and create span from it
func startSpanHeader(header *http.Header, ctx *gin.Context) opentracing.Span {
	var parent opentracing.SpanContext
	if parentSpanContext, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(*header)); err == nil {
		parent = parentSpanContext
	}
	return startSpanParent(parent, ctx)
}

func Tracer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// create span
		var span opentracing.Span

		if ctx_span, ok := ctx.Get("tracing_span"); ok {
			span = startSpanParent(ctx_span.(opentracing.SpanContext), ctx)
		} else {
			span = startSpanHeader(&ctx.Request.Header, ctx)
		}
		defer span.Finish()
		ctx.Set("tracing_span", span) // add span to ctx to use it when requests happen
		ctx.Next()
		span.SetTag("http.status_code", ctx.Writer.Status())
	}
}
