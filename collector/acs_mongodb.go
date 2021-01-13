package collector

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

type CloudMonitorMongoDB struct {
	project Project
}

func NewCloudMonitorMongoDB(c *cms.Client) *CloudMonitorMongoDB {
	return &CloudMonitorMongoDB{
		project: Project{
			client:    c,
			Namespace: "acs_mongodb",
		},
	}
}

func (db *CloudMonitorMongoDB) RetrieveCPUUtilization() []datapoint {
	return retrieve("CPUUtilization", db.project)
}

func (db *CloudMonitorMongoDB) RetrieveMemoryUtilization() []datapoint {
	return retrieve("MemoryUtilization", db.project)
}

func (db *CloudMonitorMongoDB) RetrieveDiskUtilization() []datapoint {
	return retrieve("DiskUtilization", db.project)
}

func (db *CloudMonitorMongoDB) RetrieveIOPSUtilization() []datapoint {
	return retrieve("IOPSUtilization", db.project)
}

func (db *CloudMonitorMongoDB) RetrieveConnectionUtilization() []datapoint {
	return retrieve("ConnectionUtilization", db.project)
}
