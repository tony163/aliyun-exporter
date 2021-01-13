package collector

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

type RedisExporter struct {
	project Project
}

func NewRedisExporter(c *cms.Client) *RedisExporter {
	return &RedisExporter{
		project: Project{
			client:    c,
			Namespace: "acs_kvstore",
		},
	}
}

// Standard
func (db *RedisExporter) RetrieveStandardMemoryUsage() []datapoint {
	return retrieve("StandardMemoryUsage", db.project)
}

func (db *RedisExporter) RetrieveStandardConnectionUsage() []datapoint {
	return retrieve("StandardConnectionUsage", db.project)
}

func (db *RedisExporter) RetrieveStandardCpuUsage() []datapoint {
	return retrieve("StandardCpuUsage", db.project)
}

func (db *RedisExporter) RetrieveStandardAvgRt() []datapoint {
	return retrieve("StandardAvgRt", db.project)
}

// Sharding
func (db *RedisExporter) RetrieveShardingMemoryUsage() []datapoint {
	return retrieve("ShardingMemoryUsage", db.project)
}

func (db *RedisExporter) RetrieveShardingConnectionUsage() []datapoint {
	return retrieve("ShardingConnectionUsage", db.project)
}

func (db *RedisExporter) RetrieveShardingCpuUsage() []datapoint {
	return retrieve("ShardingCpuUsage", db.project)
}

func (db *RedisExporter) RetrieveShardingAvgRt() []datapoint {
	return retrieve("ShardingAvgRt", db.project)
}

// Splitrw
func (db *RedisExporter) RetrieveSplitrwMemoryUsage() []datapoint {
	return retrieve("SplitrwMemoryUsage", db.project)
}

func (db *RedisExporter) RetrieveSplitrwConnectionUsage() []datapoint {
	return retrieve("SplitrwConnectionUsage", db.project)
}

func (db *RedisExporter) RetrieveSplitrwCpuUsage() []datapoint {
	return retrieve("SplitrwCpuUsage", db.project)
}
