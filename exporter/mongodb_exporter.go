package exporter

import (
	"aliyun-exporter/collector"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	mongodbnamespace = "aliyun"
)

type MongoDBCloudmonitorExporter struct {
	client *cms.Client

	mongoCPUUtilization        *prometheus.Desc
	mongoMemoryUtilization     *prometheus.Desc
	mongoDiskUtilization       *prometheus.Desc
	mongoIOPSUtilization       *prometheus.Desc
	mongoConnectionUtilization *prometheus.Desc
}

func MongoDBExporter(c *cms.Client) *MongoDBCloudmonitorExporter {
	return &MongoDBCloudmonitorExporter{
		client: c,

		mongoCPUUtilization: prometheus.NewDesc(
			prometheus.BuildFQName(mongodbnamespace, "mongodb", "cpu_utilization"),
			"Cpu Utilization per minute",
			[]string{
				"instanceid",
				"userid",
				"role",
			},
			nil,
		),

		mongoMemoryUtilization: prometheus.NewDesc(
			prometheus.BuildFQName(mongodbnamespace, "mongodb", "memory_utilization"),
			"Memory Utilization per minute",
			[]string{
				"instanceid",
				"userid",
				"role",
			},
			nil,
		),
		mongoDiskUtilization: prometheus.NewDesc(
			prometheus.BuildFQName(mongodbnamespace, "mongodb", "disk_utilization"),
			"Disk Utilization per minute",
			[]string{
				"instanceid",
				"userid",
				"role",
			},
			nil,
		),
		mongoIOPSUtilization: prometheus.NewDesc(
			prometheus.BuildFQName(mongodbnamespace, "mongodb", "iops_utilization"),
			"Iops Utilization per minute",
			[]string{
				"instanceid",
				"userid",
				"role",
			},
			nil,
		),
		mongoConnectionUtilization: prometheus.NewDesc(
			prometheus.BuildFQName(mongodbnamespace, "mongodb", "connection_utilization"),
			"Connection Utilization per minute",
			[]string{
				"instanceid",
				"userid",
				"role",
			},
			nil,
		),
	}
}

func (e *MongoDBCloudmonitorExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.mongoCPUUtilization
	ch <- e.mongoConnectionUtilization
	ch <- e.mongoDiskUtilization
	ch <- e.mongoIOPSUtilization
	ch <- e.mongoMemoryUtilization
}

func (e *MongoDBCloudmonitorExporter) Collect(ch chan<- prometheus.Metric) {
	mongoDashboard := collector.NewCloudMonitorMongoDB(e.client)

	for _, point := range mongoDashboard.RetrieveCPUUtilization() {
		ch <- prometheus.MustNewConstMetric(
			e.mongoCPUUtilization,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
			point.UserId,
			point.Role,
		)
	}

	for _, point := range mongoDashboard.RetrieveConnectionUtilization() {
		ch <- prometheus.MustNewConstMetric(
			e.mongoConnectionUtilization,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
			point.UserId,
			point.Role,
		)
	}

	for _, point := range mongoDashboard.RetrieveDiskUtilization() {
		ch <- prometheus.MustNewConstMetric(
			e.mongoDiskUtilization,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
			point.UserId,
			point.Role,
		)
	}

	for _, point := range mongoDashboard.RetrieveIOPSUtilization() {
		ch <- prometheus.MustNewConstMetric(
			e.mongoIOPSUtilization,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
			point.UserId,
			point.Role,
		)
	}

	for _, point := range mongoDashboard.RetrieveMemoryUtilization() {
		ch <- prometheus.MustNewConstMetric(
			e.mongoMemoryUtilization,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
			point.UserId,
			point.Role,
		)
	}
}
