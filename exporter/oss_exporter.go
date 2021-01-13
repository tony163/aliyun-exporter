package exporter

import (
	"aliyun-exporter/collector"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	ossnamespace = "aliyun"
)

type OssCloudmonitorExporter struct {
	client *cms.Client

	// oss
	ossUserAvailability      *prometheus.Desc
	ossUserRequestValidRate  *prometheus.Desc
	ossUserTotalRequestCount *prometheus.Desc
}

func OssExporter(c *cms.Client) *OssCloudmonitorExporter {
	return &OssCloudmonitorExporter{
		client: c,
		// oss
		ossUserAvailability: prometheus.NewDesc(
			prometheus.BuildFQName(ossnamespace, "oss", "user_availability"),
			"Availability per minute",
			nil,
			nil,
		),

		ossUserRequestValidRate: prometheus.NewDesc(
			prometheus.BuildFQName(ossnamespace, "oss", "user_request_valid_rate"),
			"RequestValidRate per minute",
			nil,
			nil,
		),

		ossUserTotalRequestCount: prometheus.NewDesc(
			prometheus.BuildFQName(ossnamespace, "oss", "total_request_count"),
			"TotalRequestCount per minute",
			nil,
			nil,
		),
	}
}

// Describe describes all the metrics exported by the cloudmonitor exporter.
// It implements prometheus.Collector.
func (e *OssCloudmonitorExporter) Describe(ch chan<- *prometheus.Desc) {
	// oss dashboard
	ch <- e.ossUserAvailability
	ch <- e.ossUserRequestValidRate
	ch <- e.ossUserTotalRequestCount
}

// Collect fetches the metrics from Aliyun cms
// It implements prometheus.Collector.
func (e *OssCloudmonitorExporter) Collect(ch chan<- prometheus.Metric) {
	ossDashboard := collector.NewCloudMonitorOss(e.client)

	//oss
	for _, point := range ossDashboard.RetrieveUserAvailability() {
		ch <- prometheus.MustNewConstMetric(
			e.ossUserAvailability,
			prometheus.GaugeValue,
			float64(point.UserAvailability),
		)
	}

	for _, point := range ossDashboard.RetrieveUserRequestValidRate() {
		ch <- prometheus.MustNewConstMetric(
			e.ossUserRequestValidRate,
			prometheus.GaugeValue,
			float64(point.UserRequestValidRate),
		)
	}

	for _, point := range ossDashboard.RetrieveUserTotalRequestCount() {
		ch <- prometheus.MustNewConstMetric(
			e.ossUserTotalRequestCount,
			prometheus.GaugeValue,
			float64(point.UserTotalRequestCount),
		)
	}
}
