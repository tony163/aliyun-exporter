package exporter

import (
	"aliyun-exporter/collector"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	kvstorenamespace = "aliyun"
)

type RedisExporter struct {
	client *cms.Client

	// redis dashboard
	// Standard
	redisStandardMemoryUsage     *prometheus.Desc
	redisStandardConnectionUsage *prometheus.Desc
	redisStandardCpuUsage        *prometheus.Desc
	redisStandardAvgRt           *prometheus.Desc

	// Sharding
	redisShardingMemoryUsage     *prometheus.Desc
	redisShardingConnectionUsage *prometheus.Desc
	redisShardingCpuUsage        *prometheus.Desc
	redisShardingAvgRt           *prometheus.Desc

	// Splitrw
	redisSplitrwMemoryUsage     *prometheus.Desc
	redisSplitrwConnectionUsage *prometheus.Desc
	redisSplitrwCpuUsage        *prometheus.Desc
}

// NewExporter instantiate an KvstoreExporter
func KvstoreExporter(c *cms.Client) *RedisExporter {
	return &RedisExporter{
		client: c,

		// Standard
		redisStandardMemoryUsage: prometheus.NewDesc(
			prometheus.BuildFQName(kvstorenamespace, "redis", "standard_memory_usage"),
			"Standard Memory Usage per minute",
			[]string{
				"instanceid",
			},
			nil,
		),
		redisStandardConnectionUsage: prometheus.NewDesc(
			prometheus.BuildFQName(kvstorenamespace, "redis", "standard_connection_usage"),
			"Standard Connection Usage per minute",
			[]string{
				"instanceid",
			},
			nil,
		),
		redisStandardCpuUsage: prometheus.NewDesc(
			prometheus.BuildFQName(kvstorenamespace, "redis", "standard_cpu_usage"),
			"Standard Cpu Usage per minute",
			[]string{
				"instanceid",
			},
			nil,
		),
		redisStandardAvgRt: prometheus.NewDesc(
			prometheus.BuildFQName(kvstorenamespace, "redis", "standard_avg_rt"),
			"Standard avg rt us per minute",
			[]string{
				"instanceid",
			},
			nil,
		),

		// Sharding
		redisShardingMemoryUsage: prometheus.NewDesc(
			prometheus.BuildFQName(kvstorenamespace, "redis", "sharding_memory_usage"),
			"Sharding Memory Usage per minute",
			[]string{
				"instanceid",
				"node_id",
			},
			nil,
		),
		redisShardingConnectionUsage: prometheus.NewDesc(
			prometheus.BuildFQName(kvstorenamespace, "redis", "sharding_connection_usage"),
			"Sharding Connection Usage per minute",
			[]string{
				"instanceid",
				"node_id",
			},
			nil,
		),
		redisShardingCpuUsage: prometheus.NewDesc(
			prometheus.BuildFQName(kvstorenamespace, "redis", "sharding_cpu_usage"),
			"Sharding Cpu Usage per minute",
			[]string{
				"instanceid",
				"node_id",
			},
			nil,
		),
		redisShardingAvgRt: prometheus.NewDesc(
			prometheus.BuildFQName(kvstorenamespace, "redis", "sharding_avg_rt"),
			"Sharding avg rt us per minute",
			[]string{
				"instanceid",
				"node_id",
			},
			nil,
		),

		// Splitrw
		redisSplitrwMemoryUsage: prometheus.NewDesc(
			prometheus.BuildFQName(kvstorenamespace, "redis", "splitrw_memory_usage"),
			"splitrw Memory Usage per minute",
			[]string{
				"instanceid",
				"node_id",
			},
			nil,
		),
		redisSplitrwConnectionUsage: prometheus.NewDesc(
			prometheus.BuildFQName(kvstorenamespace, "redis", "splitrw_connection_usage"),
			"Sharding Connection Usage per minute",
			[]string{
				"instanceid",
				"node_id",
			},
			nil,
		),
		redisSplitrwCpuUsage: prometheus.NewDesc(
			prometheus.BuildFQName(kvstorenamespace, "redis", "splitrw_cpu_usage"),
			"Sharding Cpu Usage per minute",
			[]string{
				"instanceid",
				"node_id",
			},
			nil,
		),
	}
}

// Describe describes all the metrics exported by the cloudmonitor exporter.
// It implements prometheus.Collector.
func (e *RedisExporter) Describe(ch chan<- *prometheus.Desc) {
	// Standard
	ch <- e.redisStandardMemoryUsage
	ch <- e.redisStandardCpuUsage
	ch <- e.redisStandardConnectionUsage
	ch <- e.redisStandardAvgRt

	// Sharding
	ch <- e.redisShardingMemoryUsage
	ch <- e.redisShardingCpuUsage
	ch <- e.redisShardingConnectionUsage
	ch <- e.redisShardingAvgRt

	// Splitrw
	ch <- e.redisSplitrwMemoryUsage
	ch <- e.redisSplitrwCpuUsage
	ch <- e.redisSplitrwConnectionUsage
}

// Collect fetches the metrics from Aliyun cms
// It implements prometheus.Collector.
func (e *RedisExporter) Collect(ch chan<- prometheus.Metric) {
	redisDashboard := collector.NewRedisExporter(e.client)

	// Standard
	for _, point := range redisDashboard.RetrieveStandardMemoryUsage() {
		ch <- prometheus.MustNewConstMetric(
			e.redisStandardMemoryUsage,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
		)
	}
	for _, point := range redisDashboard.RetrieveStandardConnectionUsage() {
		ch <- prometheus.MustNewConstMetric(
			e.redisStandardConnectionUsage,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
		)
	}
	for _, point := range redisDashboard.RetrieveStandardCpuUsage() {
		ch <- prometheus.MustNewConstMetric(
			e.redisStandardCpuUsage,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
		)
	}
	for _, point := range redisDashboard.RetrieveStandardAvgRt() {
		ch <- prometheus.MustNewConstMetric(
			e.redisStandardAvgRt,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
		)
	}

	// Sharding
	for _, point := range redisDashboard.RetrieveShardingMemoryUsage() {
		ch <- prometheus.MustNewConstMetric(
			e.redisShardingMemoryUsage,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
			point.NodeId,
		)
	}
	for _, point := range redisDashboard.RetrieveShardingCpuUsage() {
		ch <- prometheus.MustNewConstMetric(
			e.redisShardingCpuUsage,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
			point.NodeId,
		)
	}
	for _, point := range redisDashboard.RetrieveShardingConnectionUsage() {
		ch <- prometheus.MustNewConstMetric(
			e.redisShardingConnectionUsage,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
			point.NodeId,
		)
	}
	for _, point := range redisDashboard.RetrieveShardingAvgRt() {
		ch <- prometheus.MustNewConstMetric(
			e.redisShardingAvgRt,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
			point.NodeId,
		)
	}
	// Splitrw
	for _, point := range redisDashboard.RetrieveSplitrwMemoryUsage() {
		ch <- prometheus.MustNewConstMetric(
			e.redisSplitrwMemoryUsage,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
			point.NodeId,
		)
	}
	for _, point := range redisDashboard.RetrieveSplitrwCpuUsage() {
		ch <- prometheus.MustNewConstMetric(
			e.redisSplitrwCpuUsage,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
			point.NodeId,
		)
	}
	for _, point := range redisDashboard.RetrieveSplitrwConnectionUsage() {
		ch <- prometheus.MustNewConstMetric(
			e.redisSplitrwConnectionUsage,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
			point.NodeId,
		)
	}
}
