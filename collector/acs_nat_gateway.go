package collector

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

// NatGateway represents nat gateway dashboard
type NatGateway struct {
	project Project
}

// NewNatGateway returns a project respresents acs_nat_gateway
func NewNatGateway(c *cms.Client) *NatGateway {
	return &NatGateway{
		project: Project{
			client:    c,
			Namespace: "acs_nat_gateway",
		},
	}
}

func (db *NatGateway) RetrieveNetTxRate() []datapoint {
	return retrieve("net_tx.rate", db.project)
}

func (db *NatGateway) RetrieveNetrxRate() []datapoint {
	return retrieve("net_rx.rate", db.project)
}

func (db *NatGateway) RetrieveNetTxRatePercent() []datapoint {
	return retrieve("net_tx.ratePercent", db.project)
}

func (db *NatGateway) RetrieveSnatConn() []datapoint {
	return retrieve("SnatConnection", db.project)
}
