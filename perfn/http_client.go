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
	"bytes"
	"fmt"
	"github.com/perf-tool/perf-network-go/metrics"
	"github.com/perf-tool/perf-network-go/util"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

func httpClientRun(clientConfig ClientConfig) {
	for i := 0; i < clientConfig.ConnNum; i++ {
		path := util.GetEnvStr("CLIENT_HTTP_PATH", "perf")
		url := fmt.Sprintf("http://%s/%s", clientConfig.addr(), path)
		go func() {
			ticker := time.NewTicker(time.Duration(clientConfig.TickPerConnMs) * time.Millisecond)
			for range ticker.C {
				message := util.RandBytes(clientConfig.PacketSize)
				startTime := time.Now()
				resp, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBuffer(message))
				if err != nil {
					metrics.HttpClientSendFailCount.Inc()
					metrics.HttpClientConnSendFailCount.WithLabelValues(clientConfig.addr()).Inc()
					logrus.Error("send http request message error: ", err)
					break
				} else {
					metrics.HttpClientSendBytesCount.Add(float64(len(message)))
					cost := time.Since(startTime).Milliseconds()
					metrics.HttpClientSendSuccessCount.Inc()
					metrics.HttpClientSendSuccessLatency.Observe(float64(cost))
					metrics.HttpClientConnSendSuccessCount.WithLabelValues(clientConfig.addr()).Inc()
					metrics.HttpClientConnSendSuccessLatency.WithLabelValues(clientConfig.addr()).Observe(
						float64(cost))
					_, err := io.ReadAll(resp.Body)
					if err != nil {
						logrus.Error("read http body error: ", err)
						_ = resp.Body.Close()
						break
					} else {
						_ = resp.Body.Close()
					}
				}
			}
		}()
	}
}
