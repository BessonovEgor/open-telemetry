package opentelemetrytracer

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	oTelTrace "go.opentelemetry.io/otel/sdk/trace"
)

const retrieveFormat = "%s.%s"

type Tracer interface {
	TraceData(request string) func()
}

type OpenTelemetryTracer struct {
	serviceName string
	exporter    oTelTrace.SpanExporter
}

// NewTracer creates a new OpenTelemetryTracer instance.
func NewTracer(exporter oTelTrace.SpanExporter, serviceName string) Tracer {
	return &OpenTelemetryTracer{
		serviceName: serviceName,
		exporter:    exporter,
	}
}

// TraceData default trace function.
func (t *OpenTelemetryTracer) TraceData(sql string) func() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tCtx, span := otel.Tracer(t.serviceName).Start(ctx, fmt.Sprintf(retrieveFormat, t.serviceName, sql))
	defer span.End()

	err := t.exporter.ExportSpans(tCtx, []oTelTrace.ReadOnlySpan{})
	if err != nil {
		return nil
	}

	spanFinish := func() {
		span.End()
	}
	return spanFinish
}
