package trace

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

// HTTPHandler is a convenience function which helps attaching tracing
// functionality to conventional HTTP handlers.
func HTTPHandler(handler http.Handler, name string) http.Handler {
	return otelhttp.NewHandler(handler, name, otelhttp.WithTracerProvider(otel.GetTracerProvider()))
}

// HTTPHandlerFunc is a convenience function which helps attaching tracing
// functionality to conventional HTTP handlers.
func HTTPHandlerFunc(handler http.HandlerFunc, name string) http.HandlerFunc {
	return otelhttp.NewHandler(handler, name, otelhttp.WithTracerProvider(otel.GetTracerProvider())).ServeHTTP
}
