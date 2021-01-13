package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"aliyun-exporter/exporter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var config struct {
	accessKeyId     string
	accessKeySecret string
	regionId        string
	host            string
	port            int
	service         string
}

func main() {
	flag.StringVar(&(config.accessKeyId), "id", os.Getenv("ACCESS_KEY_ID"), "阿里云AccessKey ID")
	flag.StringVar(&(config.accessKeySecret), "secret", os.Getenv("ACCESS_KEY_SECRET"), "阿里云AccessKey Secret")
	flag.StringVar(&(config.regionId), "region", os.Getenv("REGIONID"), "阿里云Region ID")
	flag.StringVar(&(config.host), "host", "0.0.0.0", "服务监听地址")
	flag.IntVar(&(config.port), "port", 9180, "服务监听端口")
	flag.StringVar(&(config.service), "service", "kvstore,nat,oss,rds,slb,mongodb", "输出Metrics的服务，默认为全部")
	flag.Parse()

	serviceArr := strings.Split(config.service, ",")

	for _, ae := range serviceArr {
		switch ae {
		case "kvstore":
			kvstore := exporter.KvstoreExporter(newCmsClient())
			prometheus.MustRegister(kvstore)
		case "nat":
			nat := exporter.NatExporter(newCmsClient())
			prometheus.MustRegister(nat)
		case "oss":
			oss := exporter.OssExporter(newCmsClient())
			prometheus.MustRegister(oss)
		case "rds":
			rds := exporter.RdsExporter(newCmsClient())
			prometheus.MustRegister(rds)
		case "slb":
			slb := exporter.SlbExporter(newCmsClient())
			prometheus.MustRegister(slb)
		case "mongodb":
			mongo := exporter.MongoDBExporter(newCmsClient())
			prometheus.MustRegister(mongo)
		default:
			log.Println("暂不支持该服务，请根据提示选择服务。")
		}
	}

	// 启动服务
	listenAddress := net.JoinHostPort(config.host, strconv.Itoa(config.port))
	log.Println("Running on", listenAddress)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
