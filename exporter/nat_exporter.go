package exporter

import (
	"aliyun-exporter/collector"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	natnamespace = "aliyun"
)

type NatCloudmonitorExporter struct {
	client *cms.Client

	// nat gateway
	netTxRate        *prometheus.Desc
	netTxRatePercent *prometheus.Desc
	snatConnections  *prometheus.Desc
}

// NatNewExporter instantiate an CloudmonitorExport
func NatExporter(c *cms.Client) *NatCloudmonitorExporter {
	return &NatCloudmonitorExporter{
		client: c,

		// nat gateway
		netTxRate: prometheus.NewDesc(
			prometheus.BuildFQName(natnamespace, "net_tx_rate", "bytes"),
			"Outbound bandwith of gateway in bits/s",
			[]string{
				"instanceid",
			},
			nil,
		),
		netTxRatePercent: prometheus.NewDesc(
			prometheus.BuildFQName(natnamespace, "net_tx_rate", "percent"),
			"Outbound bandwith of gateway used in percentage",
			[]string{
				"instanceid",
			},
			nil,
		),
		snatConnections: prometheus.NewDesc(
			prometheus.BuildFQName(natnamespace, "snat", "connections"),
			"Max number of snat connections per minute",
			[]string{
				"instanceid",
			},
			nil,
		),
	}
}

// Describe describes all the metrics exported by the cloudmonitor exporter.
// It implements prometheus.Collector.
func (e *NatCloudmonitorExporter) Describe(ch chan<- *prometheus.Desc) {
	// nat gateway
	ch <- e.netTxRate
	ch <- e.netTxRatePercent
	ch <- e.snatConnections
}

// Collect fetches the metrics from Aliyun cms
// It implements prometheus.Collector.
func (e *NatCloudmonitorExporter) Collect(ch chan<- prometheus.Metric) {
	natGateway := collector.NewNatGateway(e.client)

	// nat
	for _, point := range natGateway.RetrieveNetTxRate() {
		ch <- prometheus.MustNewConstMetric(
			e.netTxRate,
			prometheus.GaugeValue,
			float64(point.Value),
			point.InstanceId,
		)
	}

	for _, point := range natGateway.RetrieveNetTxRatePercent() {
		ch <- prometheus.MustNewConstMetric(
			e.netTxRatePercent,
			prometheus.GaugeValue,
			float64(point.Value),
			point.InstanceId,
		)
	}

	for _, point := range natGateway.RetrieveSnatConn() {
		ch <- prometheus.MustNewConstMetric(
			e.snatConnections,
			prometheus.GaugeValue,
			float64(point.Maximum),
			point.InstanceId,
		)
	}
}
