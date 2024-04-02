package service

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	oTelMetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/resource"
)

type MetricExporter interface {
	Count(name, description string, value int64, tags map[string]string) error
	Gauge(name, description string, value float64, tags map[string]string) error
	Histogram(name, description string, value float64, tags map[string]string) error
}

// MetricService is a service to send metrics via oTelMetric.Exporter.
type MetricService struct {
	exporter oTelMetric.Exporter
}

func NewMetricService(exporter oTelMetric.Exporter) MetricExporter {
	return &MetricService{
		exporter: exporter,
	}
}

// Count exports count metric on open telemetry collector.
func (t *MetricService) Count(name, description string, value int64, tags map[string]string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var attributes []attribute.KeyValue
	for key, tag := range tags {
		attributes = append(attributes, attribute.String(key, tag))
	}

	res, err := resource.New(ctx, resource.WithAttributes(attributes...))
	if err != nil {
		return err
	}
	return t.exporter.Export(ctx, &metricdata.ResourceMetrics{
		Resource: res,
		ScopeMetrics: []metricdata.ScopeMetrics{{
			Scope: instrumentation.Scope{},
			Metrics: []metricdata.Metrics{{
				Name:        name,
				Description: description,
				Unit:        "",
				Data: metricdata.Sum[int64]{
					DataPoints: []metricdata.DataPoint[int64]{{
						Attributes: attribute.Set{},
						StartTime:  time.Now().UTC(),
						Time:       time.Now().UTC(),
						Value:      value,
						Exemplars:  nil,
					}},
					Temporality: metricdata.CumulativeTemporality,
					IsMonotonic: false,
				},
			}},
		}},
	})
}

// Gauge exports gauge metric on open telemetry collector.
func (t *MetricService) Gauge(name, description string, value float64, tags map[string]string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var attributes []attribute.KeyValue
	for key, tag := range tags {
		attributes = append(attributes, attribute.String(key, tag))
	}

	res, err := resource.New(ctx, resource.WithAttributes(attributes...))
	if err != nil {
		return err
	}
	return t.exporter.Export(ctx, &metricdata.ResourceMetrics{
		Resource: res,
		ScopeMetrics: []metricdata.ScopeMetrics{{
			Scope: instrumentation.Scope{},
			Metrics: []metricdata.Metrics{{
				Name:        name,
				Description: description,
				Unit:        "",
				Data: metricdata.Gauge[float64]{
					DataPoints: []metricdata.DataPoint[float64]{{
						Attributes: attribute.Set{},
						StartTime:  time.Now().UTC(),
						Time:       time.Now().UTC(),
						Value:      value,
						Exemplars:  nil,
					}},
				},
			}},
		}},
	})
}

// Histogram exports histogram metric on open telemetry collector.
func (t *MetricService) Histogram(name, description string, value float64, tags map[string]string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var attributes []attribute.KeyValue
	for key, tag := range tags {
		attributes = append(attributes, attribute.String(key, tag))
	}

	res, err := resource.New(ctx, resource.WithAttributes(attributes...))
	if err != nil {
		return err
	}
	return t.exporter.Export(ctx, &metricdata.ResourceMetrics{
		Resource: res,
		ScopeMetrics: []metricdata.ScopeMetrics{{
			Scope: instrumentation.Scope{},
			Metrics: []metricdata.Metrics{{
				Name:        name,
				Description: description,
				Unit:        "",
				Data: metricdata.Histogram[float64]{
					DataPoints: []metricdata.HistogramDataPoint[float64]{{
						Attributes:   attribute.Set{},
						StartTime:    time.Now().UTC(),
						Time:         time.Now().UTC(),
						Count:        1,
						Bounds:       nil,
						BucketCounts: nil,
						Min:          metricdata.Extrema[float64]{},
						Max:          metricdata.Extrema[float64]{},
						Sum:          value,
						Exemplars:    nil,
					}},
					Temporality: metricdata.CumulativeTemporality,
				},
			}},
		}},
	})
}
