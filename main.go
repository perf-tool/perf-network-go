package main

import (
	"github.com/perf-tool/perf-network-go/perfn"
	"os"
)

func main() {
	var config perfn.Config
	envProtocolType := os.Getenv("PROTOCOL_TYPE")
	if envProtocolType == "udp" {
		config.ProtocolType = perfn.ProtocolTypeUdp
	} else if envProtocolType == "tcp" {
		config.ProtocolType = perfn.ProtocolTypeTcp
	} else if envProtocolType == "http" {
		config.ProtocolType = perfn.ProtocolTypeHttp
	}
	envCommType := os.Getenv("COMM_TYPE")
	if envCommType == "client" {
		config.CommType = perfn.CommTypeClient
	} else if envCommType == "server" {
		config.CommType = perfn.CommTypeServer
	}
	err := perfn.Run(config)
	if err != nil {
		panic(err)
	}
}
