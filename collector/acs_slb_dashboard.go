package collector

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

// SLBDashboard represents the dashboard of SLB
type SLBDashboard struct {
	project Project
}

// NewSLBDashboard returns a project respresents acs_slb_dashboard
func NewSLBDashboard(c *cms.Client) *SLBDashboard {
	return &SLBDashboard{
		project: Project{
			client:    c,
			Namespace: "acs_slb_dashboard",
		},
	}
}

// below are Layer-4 metrics

func (db *SLBDashboard) RetrieveActiveConnection() []datapoint {
	return retrieve("ActiveConnection", db.project)
}

func (db *SLBDashboard) RetrievePacketTX() []datapoint {
	return retrieve("PacketTX", db.project)
}

func (db *SLBDashboard) RetrievePacketRX() []datapoint {
	return retrieve("PacketRX", db.project)
}

func (db *SLBDashboard) RetrieveTrafficRX() []datapoint {
	return retrieve("TrafficRXNew", db.project)
}

func (db *SLBDashboard) RetrieveTrafficTX() []datapoint {
	return retrieve("TrafficTXNew", db.project)
}

func (db *SLBDashboard) RetrieveNewConnection() []datapoint {
	return retrieve("NewConnection", db.project)
}

func (db *SLBDashboard) RetrieveMaxConnection() []datapoint {
	return retrieve("MaxConnection", db.project)
}

func (db *SLBDashboard) RetrieveDropConnection() []datapoint {
	return retrieve("DropConnection", db.project)
}

func (db *SLBDashboard) RetrieveDropPacketRX() []datapoint {
	return retrieve("DropPacketRX", db.project)
}

func (db *SLBDashboard) RetrieveDropPacketTX() []datapoint {
	return retrieve("DropPacketTX", db.project)
}

func (db *SLBDashboard) RetrieveDropTrafficRX() []datapoint {
	return retrieve("DropTrafficRX", db.project)
}

func (db *SLBDashboard) RetrieveDropTrafficTX() []datapoint {
	return retrieve("DropTrafficTX", db.project)
}

// below are Layer-7 metrics

func (db *SLBDashboard) RetrieveQps() []datapoint {
	return retrieve("Qps", db.project)
}

func (db *SLBDashboard) RetrieveRt() []datapoint {
	return retrieve("Rt", db.project)
}

func (db *SLBDashboard) RetrieveStatusCode5xx() []datapoint {
	return retrieve("StatusCode5xx", db.project)
}

func (db *SLBDashboard) RetrieveUpstreamCode4xx() []datapoint {
	return retrieve("UpstreamCode4xx", db.project)
}

func (db *SLBDashboard) RetrieveUpstreamCode5xx() []datapoint {
	return retrieve("UpstreamCode5xx", db.project)
}

func (db *SLBDashboard) RetrieveUpstreamRt() []datapoint {
	return retrieve("UpstreamRt", db.project)
}
