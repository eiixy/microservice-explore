package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/trace"
	"testing"
)

func TestTrackId(t *testing.T) {
	ctx := context.Background()
	span := trace.SpanContextFromContext(ctx)
	fmt.Printf("trace id: %s span id: %s \r\n", span.TraceID().String(), span.SpanID().String())

	fmt.Printf("log trace id: %s span id: %s \r\n", log.TraceID()(ctx), log.SpanID()(ctx))

}
