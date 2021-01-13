package main

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

func newCmsClient() *cms.Client {
	cmsClient, err := cms.NewClientWithAccessKey(
		config.regionId,
		config.accessKeyId,
		config.accessKeySecret,
	)

	if err != nil {
		panic(err)
	}

	return cmsClient
}
