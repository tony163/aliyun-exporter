package exporter

import (
	"aliyun-exporter/collector"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	slbnamespace = "aliyun"
)

type SlbCloudmonitorExporter struct {
	client *cms.Client

	// slb dashbaord
	activeConnection *prometheus.Desc
	packetRX         *prometheus.Desc
	packetTX         *prometheus.Desc
	trafficRX        *prometheus.Desc
	trafficTX        *prometheus.Desc
	newConnection    *prometheus.Desc
	maxConnection    *prometheus.Desc
	dropConnection   *prometheus.Desc
	dropPacketRX     *prometheus.Desc
	dropPacketTX     *prometheus.Desc
	dropTrafficRX    *prometheus.Desc
	dropTrafficTX    *prometheus.Desc
	qps              *prometheus.Desc
	rt               *prometheus.Desc
	statusCode5xx    *prometheus.Desc
	upstreamCode4xx  *prometheus.Desc
	upstreamCode5xx  *prometheus.Desc
	upstreamRt       *prometheus.Desc
}

// NewExporter instantiate an CloudmonitorExport
func SlbExporter(c *cms.Client) *SlbCloudmonitorExporter {
	return &SlbCloudmonitorExporter{
		client: c,

		// slb dashboard
		activeConnection: prometheus.NewDesc(
			prometheus.BuildFQName(slbnamespace, "slb", "active_connection"),
			"Number of active connections per minute",
			[]string{
				"instanceid",
				"port",
				"vip",
			},
			nil,
		),

		packetRX: prometheus.NewDesc(
			prometheus.BuildFQName(slbnamespace, "slb", "packet_rx_average"),
			"Average packets received per second",
			[]string{
				"instanceid",
				"port",
				"vip",
			},
			nil,
		),

		packetTX: prometheus.NewDesc(
			prometheus.BuildFQName(slbnamespace, "slb", "packet_tx_average"),
			"Average packets sent per second",
			[]string{
				"instanceid",
				"port",
				"vip",
			},
			nil,
		),

		trafficRX: prometheus.NewDesc(
			prometheus.BuildFQName(slbnamespace, "slb", "traffic_rx_average"),
			"Average traffic received per second",
			[]string{
				"instanceid",
				"port",
				"vip",
			},
			nil,
		),

		trafficTX: prometheus.NewDesc(
			prometheus.BuildFQName(slbnamespace, "slb", "traffic_tx_average"),
			"Average traffic sent per second",
			[]string{
				"instanceid",
				"port",
				"vip",
			},
			nil,
		),

		newConnection: prometheus.NewDesc(
			prometheus.BuildFQName(slbnamespace, "slb", "new_connection_average"),
			"Average number of new connections created per second",
			[]string{
				"instanceid",
				"port",
				"vip",
			},
			nil,
		),
	}
}

// Describe describes all the metrics exported by the cloudmonitor exporter.
// It implements prometheus.Collector.
func (e *SlbCloudmonitorExporter) Describe(ch chan<- *prometheus.Desc) {
	// slb dashboard
	ch <- e.activeConnection
	ch <- e.packetRX
	ch <- e.packetTX
	ch <- e.trafficRX
	ch <- e.trafficTX
	ch <- e.newConnection
}

// Collect fetches the metrics from Aliyun cms
// It implements prometheus.Collector.
func (e *SlbCloudmonitorExporter) Collect(ch chan<- prometheus.Metric) {
	slbDashboard := collector.NewSLBDashboard(e.client)

	// slb
	for _, point := range slbDashboard.RetrieveActiveConnection() {
		ch <- prometheus.MustNewConstMetric(
			e.activeConnection,
			prometheus.GaugeValue,
			float64(point.Maximum),
			point.InstanceId,
			point.Port,
			point.Vip,
		)
	}

	for _, point := range slbDashboard.RetrievePacketRX() {
		ch <- prometheus.MustNewConstMetric(
			e.packetRX,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
			point.Port,
			point.Vip,
		)
	}

	for _, point := range slbDashboard.RetrievePacketTX() {
		ch <- prometheus.MustNewConstMetric(
			e.packetTX,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
			point.Port,
			point.Vip,
		)
	}

	for _, point := range slbDashboard.RetrieveTrafficRX() {
		ch <- prometheus.MustNewConstMetric(
			e.trafficRX,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
			point.Port,
			point.Vip,
		)
	}

	for _, point := range slbDashboard.RetrieveTrafficTX() {
		ch <- prometheus.MustNewConstMetric(
			e.trafficTX,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
			point.Port,
			point.Vip,
		)
	}

	for _, point := range slbDashboard.RetrieveNewConnection() {
		ch <- prometheus.MustNewConstMetric(
			e.newConnection,
			prometheus.GaugeValue,
			float64(point.Average),
			point.InstanceId,
			point.Port,
			point.Vip,
		)
	}
}
