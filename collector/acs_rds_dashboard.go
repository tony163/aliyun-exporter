package collector

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

// RDSDashboard represents the dashboard of RDS
type RDSDashboard struct {
	project Project
}

// NewRDSDashboard returns a project respresents acs_rds_dashboard
func NewRDSDashboard(c *cms.Client) *RDSDashboard {
	return &RDSDashboard{
		project: Project{
			client:    c,
			Namespace: "acs_rds_dashboard",
		},
	}
}

func (db *RDSDashboard) RetrieveCPUUsage() []datapoint {
	return retrieve("CpuUsage", db.project)
}

func (db *RDSDashboard) RetrieveConnectionUsage() []datapoint {
	return retrieve("ConnectionUsage", db.project)
}

func (db *RDSDashboard) RetrieveMemoryUsage() []datapoint {
	return retrieve("MemoryUsage", db.project)
}

func (db *RDSDashboard) RetrieveDataDelay() []datapoint {
	return retrieve("DataDelay", db.project)
}

func (db *RDSDashboard) RetrieveDiskUsage() []datapoint {
	return retrieve("DiskUsage", db.project)
}

func (db *RDSDashboard) RetrieveIOPSUsage() []datapoint {
	return retrieve("IOPSUsage", db.project)
}
