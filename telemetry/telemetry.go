// Package telemetry provides a simplified wrapper for opentelemetry.
package telemetry

import (
	"context"

	otelmetric "go.opentelemetry.io/otel/metric"
	"interview_3_31_2026/config"
	//"go.opentelemetry.io/otel/sdk/metric"
	//"go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type TelemetryProvider struct {
	meter  otelmetric.Meter
	tracer oteltrace.Tracer
}

// NewTelemetryProvider is a singleton, and will provide a thread safe reference to the singular instance
// of the telemetry. Implementation-wise, this is lame, but it will work.
func NewTelemetryProvider(ctx context.Context, cfg *config.Config) *TelemetryProvider {
	return &TelemetryProvider{}
}
