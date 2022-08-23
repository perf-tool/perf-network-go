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
	"github.com/perf-tool/perf-network-go/metrics"
	"github.com/perf-tool/perf-network-go/util"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

func udpClientRun(clientConfig ClientConfig) {
	for i := 0; i < clientConfig.ConnNum; i++ {
		go func() {
			udpAddr, err := net.ResolveUDPAddr("udp4", clientConfig.addr())
			if err != nil {
				logrus.Error("resolve udp addr error: ", err)
				return
			}
			conn, err := net.Dial("udp4", udpAddr.String())
			if err != nil {
				logrus.Error("dial udp error: ", err)
				return
			}
			defer conn.Close()
			go func() {
				for {
					buffer := make([]byte, 2*clientConfig.PacketSize)
					n, err := conn.Read(buffer)
					if err != nil {
						logrus.Error("read udp message error: ", err)
					} else {
						metrics.UdpClientRecvBytesCount.Add(float64(n))
					}
				}
			}()
			ticker := time.NewTicker(time.Duration(clientConfig.TickPerConnMs) * time.Millisecond)
			for range ticker.C {
				message := util.RandBytes(clientConfig.PacketSize)
				startTime := time.Now()
				size, err := conn.Write(message)
				if err != nil {
					metrics.UdpClientSendFailCount.Inc()
					metrics.UdpClientConnSendFailCount.WithLabelValues(conn.LocalAddr().String(), conn.RemoteAddr().String()).Inc()
					logrus.Error("write udp message error: ", err)
					break
				} else {
					metrics.UdpClientSendBytesCount.Add(float64(size))
					cost := time.Since(startTime).Milliseconds()
					metrics.UdpClientSendSuccessCount.Inc()
					metrics.UdpClientSendSuccessLatency.Observe(float64(cost))
					metrics.UdpClientConnSendSuccessCount.WithLabelValues(conn.LocalAddr().String(), conn.RemoteAddr().String()).Inc()
					metrics.UdpClientConnSendSuccessLatency.WithLabelValues(conn.LocalAddr().String(), conn.RemoteAddr().String()).Observe(
						float64(cost))
				}
			}
		}()
	}
}
