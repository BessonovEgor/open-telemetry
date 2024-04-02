package test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/BessonovEgor/open-telemetry/service"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
)

func TestCount(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	exporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithInsecure())
	require.NoError(t, err)

	srv := service.NewMetricService(exporter)
	err = srv.Count("test_otel_count",
		"A test count metric from open telemetry collector.",
		5, map[string]string{"environment": "dev"})
	require.NoError(t, err)
}

func TestGauge(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	exporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithInsecure())
	require.NoError(t, err)

	srv := service.NewMetricService(exporter)
	err = srv.Gauge("test_otel_gauge",
		"A test gauge metric from open telemetry collector.",
		2, map[string]string{"environment": "dev"})
	require.NoError(t, err)
}

func TestHistogram(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	exporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithInsecure())
	require.NoError(t, err)

	srv := service.NewMetricService(exporter)
	err = srv.Histogram("test_otel_histogram",
		"A test histogram metric from open telemetry collector.",
		3, map[string]string{"environment": "dev"})
	require.NoError(t, err)
}

func TestContinuesRandomValuesCount(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	exporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithInsecure())
	require.NoError(t, err)

	srv := service.NewMetricService(exporter)
	for i := 0; i < 100; i++ {
		err = srv.Count("test_otel_count",
			"A test count metric from open telemetry collector.",
			int64(rand.Intn(901)+100), map[string]string{"environment": "dev"})
		require.NoError(t, err)
		<-time.After(time.Second * 60)
	}
}
