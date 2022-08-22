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

func tcpClientRun(clientConfig ClientConfig) {
	for i := 0; i < clientConfig.ConnNum; i++ {
		go func() {
			tcpAddr, err := net.ResolveTCPAddr("tcp4", clientConfig.addr())
			if err != nil {
				logrus.Error("resolve tcp addr error: ", err)
				return
			}
			conn, err := net.Dial("tcp4", tcpAddr.String())
			if err != nil {
				logrus.Error("dial tcp error: ", err)
				return
			}
			defer conn.Close()
			ticker := time.NewTicker(time.Duration(clientConfig.TickPerConnMs) * time.Millisecond)
			for range ticker.C {
				message := util.RandBytes(clientConfig.PacketSize)
				startTime := time.Now()
				_, err := conn.Write(message)
				if err != nil {
					metrics.TcpClientSendFailCount.WithLabelValues(conn.LocalAddr().String(), conn.RemoteAddr().String()).Inc()
					logrus.Error("write tcp message error: ", err)
					break
				} else {
					metrics.TcpClientSendSuccessCount.WithLabelValues(conn.LocalAddr().String(), conn.RemoteAddr().String()).Inc()
					metrics.TcpClientSendSuccessLatency.WithLabelValues(conn.LocalAddr().String(), conn.RemoteAddr().String()).Observe(
						float64(time.Since(startTime).Milliseconds()))
				}
			}
		}()
	}
}
