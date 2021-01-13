package exporter

import (
	"aliyun-exporter/collector"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	rdsnamespace = "aliyun"
)

type RdsCloudmonitorExporter struct {
	client *cms.Client
	// rds dashbaord
	cpuUsage        *prometheus.Desc
	connectionUsage *prometheus.Desc
	memoryUsage     *prometheus.Desc
	dataDelay       *prometheus.Desc
	diskUsage       *prometheus.Desc
	iopsUsage       *prometheus.Desc
}

// NewExporter instantiate an CloudmonitorExport
func RdsExporter(c *cms.Client) *RdsCloudmonitorExporter {
	return &RdsCloudmonitorExporter{
		client: c,

		// rds dashbaord
		cpuUsage: prometheus.NewDesc(
			prometheus.BuildFQName(rdsnamespace, "rds", "cpu_usage_average"),
			"CPU usage per minute",
			[]string{
				"instanceid",
			},
			nil,
		),

		connectionUsage: prometheus.NewDesc(
			prometheus.BuildFQName(rdsnamespace, "rds", "connection_usage"),
			"Connection usage per minute",
			[]string{
				"instanceid",
			},
			nil,
		),
		memoryUsage: prometheus.NewDesc(
			prometheus.BuildFQName(rdsnamespace, "rds", "memory_usage"),
			"Memory usage per minute",
			[]string{
				"instanceid",
			},
			nil,
		),
		dataDelay: prometheus.NewDesc(
			prometheus.BuildFQName(rdsnamespace, "rds", "data_delay"),
			"Slave data delay minute",
			[]string{
				"instanceid",
			},
			nil,
		),
		diskUsage: prometheus.NewDesc(
			prometheus.BuildFQName(rdsnamespace, "rds", "disk_usage"),
			"Disk Usage per minute",
			[]string{
				"instanceid",
			},
			nil,
		),
		iopsUsage: prometheus.NewDesc(
			prometheus.BuildFQName(rdsnamespace, "rds", "iops_usage"),
			"Iops Usage per minute",
			[]string{
				"instanceid",
			},
			nil,
		),
	}
}

// Describe describes all the metrics exported by the cloudmonitor exporter.
// It implements prometheus.Collector.
func (e *RdsCloudmonitorExporter) Describe(ch chan<- *prometheus.Desc) {
	// rds dashbaord
	ch <- e.cpuUsage
	ch <- e.connectionUsage
	ch <- e.memoryUsage
	ch <- e.dataDelay
	ch <- e.diskUsage
	ch <- e.iopsUsage
}

// Collect fetches the metrics from Aliyun cms
// It implements prometheus.Collector.
func (e *RdsCloudmonitorExporter) Collect(ch chan<- prometheus.Metric) {
	rdsDashboard := collector.NewRDSDashboard(e.client)

	// rds
	for _, point := range rdsDashboard.RetrieveCPUUsage() {
		ch <- prometheus.MustNewConstMetric(
			e.cpuUsage,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
		)
	}

	for _, point := range rdsDashboard.RetrieveConnectionUsage() {
		ch <- prometheus.MustNewConstMetric(
			e.connectionUsage,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
		)
	}
	for _, point := range rdsDashboard.RetrieveMemoryUsage() {
		ch <- prometheus.MustNewConstMetric(
			e.memoryUsage,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
		)
	}

	for _, point := range rdsDashboard.RetrieveDataDelay() {
		ch <- prometheus.MustNewConstMetric(
			e.dataDelay,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
		)
	}

	for _, point := range rdsDashboard.RetrieveDiskUsage() {
		ch <- prometheus.MustNewConstMetric(
			e.diskUsage,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
		)
	}
	for _, point := range rdsDashboard.RetrieveIOPSUsage() {
		ch <- prometheus.MustNewConstMetric(
			e.iopsUsage,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
		)
	}
}
