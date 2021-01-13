package collector

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

type CloudMonitorOss struct {
	project Project
}

func NewCloudMonitorOss(c *cms.Client) *CloudMonitorOss {
	return &CloudMonitorOss{
		project: Project{
			client:    c,
			Namespace: "acs_oss",
		},
	}
}

func (db *CloudMonitorOss) RetrieveUserAvailability() []datapoint {
	return retrieve("UserAvailability", db.project)
}

func (db *CloudMonitorOss) RetrieveUserRequestValidRate() []datapoint {
	return retrieve("UserRequestValidRate", db.project)
}

func (db *CloudMonitorOss) RetrieveUserTotalRequestCount() []datapoint {
	return retrieve("UserTotalRequestCount", db.project)
}
