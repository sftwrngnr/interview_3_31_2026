// Package telemetry provides a simplified wrapper for opentelemetry.
package telemetry

import (
	"context"
	"sync"

	"interview_3_31_2026/config"

	otelmetric "go.opentelemetry.io/otel/metric"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var lock sync.Mutex
var telemInstance *TelemetryProvider

type TelemetryProvider struct {
	config *config.Config
	ctx    context.Context
	meter  otelmetric.Meter
	tracer oteltrace.Tracer
}

// NewTelemetryProvider is a singleton, and will provide a thread safe reference to the singular instance
// of the telemetry. Implementation-wise, this is lame, but it will work.
func NewTelemetryProvider(ctx context.Context, cfg *config.Config) *TelemetryProvider {
	if telemInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if telemInstance == nil {
			telemInstance = &TelemetryProvider{config: cfg, ctx: ctx}
		}
	}
	return telemInstance
}

func GetTelemetryProvider() *TelemetryProvider {
	return telemInstance
}
