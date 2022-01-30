package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/semconv"
	"go.opentelemetry.io/otel/trace"

	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// ProviderConfig represents the provider configuration and used to create a new
// `Provider` type.
type ProviderConfig struct {
	JaegerEndpoint string
	ServiceName    string
	ServiceVersion string
	Environment    string
	// Set this to `true` if you want to disable tracing completly.
	Disabled bool
}

// Provider represents the tracer provider. Depending on the `config.Disabled`
// parameter, it will either use a "live" provider or a "no operations" version.
// The "no operations" means, tracing will be globally disabled.
type Provider struct {
	provider trace.TracerProvider
}

// New returns a new `Provider` type. It uses Jaeger exporter and globally sets
// the tracer provider as well as the global tracer for spans.
func NewProvider(config ProviderConfig) (Provider, error) {
	if config.Disabled {
		return Provider{provider: trace.NewNoopTracerProvider()}, nil
	}

	exp, err := jaeger.NewRawExporter(
		jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(config.JaegerEndpoint)),
	)
	if err != nil {
		return Provider{}, err
	}

	prv := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(sdkresource.NewWithAttributes(
			semconv.ServiceNameKey.String(config.ServiceName),
			semconv.ServiceVersionKey.String(config.ServiceVersion),
			semconv.DeploymentEnvironmentKey.String(config.Environment),
		)),
	)

	otel.SetTracerProvider(prv)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	tracer = struct{ trace.Tracer }{otel.Tracer(config.ServiceName)}

	return Provider{provider: prv}, nil
}

// Close shuts down the tracer provider only if it was not "no operations"
// version.
func (p Provider) Close(ctx context.Context) error {
	if prv, ok := p.provider.(*sdktrace.TracerProvider); ok {
		return prv.Shutdown(ctx)
	}

	return nil
}
