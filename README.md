## Aliyun Cloud Monitor Exporter

通过阿里云CMS API获取metrics数据，需配合Prometheus使用。



### 构建

```
go get -d
go build
```



### 运行

```
./aliyun-exporter -id access_id -secret access_secret -region cn-hangzhou
```



### 参考

[云服务主要监控项](https://www.alibabacloud.com/help/zh/doc-detail/28619.htm?spm=a2c63.p38356.b99.119.1c4f38f4vw6wdV#title-ykw-yt8-lus)

[API概览](https://help.aliyun.com/document_detail/28617.html?spm=a2c4g.11186623.6.710.401d7084DesZNh)

[alibaba-cloud-sdk-go](https://github.com/aliyun/alibaba-cloud-sdk-go)

[OpenAPI Explorer](https://api.aliyun.com/#/?product=Cms&version=2019-01-01&api=DescribeMetricMetaList&params={}&tab=DEMO&lang=GO)

[nevill/cloudmonitor_exporter](https://github.com/nevill/cloudmonitor_exporter)

