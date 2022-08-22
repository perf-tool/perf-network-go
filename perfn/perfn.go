// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package perfn

import (
	"github.com/perf-tool/perf-network-go/util"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
)

func Run(config Config) error {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":20008", nil)
	if err != nil {
		return err
	}
	switch config.ProtocolType {
	case ProtocolTypeUdp:
		if config.CommType == CommTypeClient {
			logrus.Info("start udp client")
			udpClientRun(getClientConfig())
		}
	case ProtocolTypeTcp:
		if config.CommType == CommTypeClient {
			logrus.Info("start tcp client")
			tcpClientRun(getClientConfig())
		}
	case ProtocolTypeHttp:
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		<-interrupt
	}
}

func getClientConfig() ClientConfig {
	return ClientConfig{
		Host:          util.GetEnvStr("CLIENT_HOST", "localhost"),
		Port:          util.GetEnvInt("CLIENT_PORT", 5678),
		ConnNum:       util.GetEnvInt("CLIENT_CONN_NUM", 10),
		TickPerConnMs: util.GetEnvInt("CLIENT_TICK_PER_CONN_MS", 1000),
		PacketSize:    util.GetEnvInt("CLIENT_PACKET_SIZE", 1024),
	}
}
