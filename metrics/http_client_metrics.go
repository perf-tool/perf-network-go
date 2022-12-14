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

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HttpClientRecvBytesCount = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: prometheus.BuildFQName(namespace, "http_client", "recv_bytes_total"),
		},
	)
	HttpClientSendBytesCount = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: prometheus.BuildFQName(namespace, "http_client", "send_bytes_total"),
		},
	)
	HttpClientSendSuccessCount = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: prometheus.BuildFQName(namespace, "http_client", "send_success_total"),
		},
	)
	HttpClientSendFailCount = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: prometheus.BuildFQName(namespace, "http_client", "send_fail_total"),
		},
	)
	HttpClientSendSuccessLatency = promauto.NewSummary(
		prometheus.SummaryOpts{
			Name:       prometheus.BuildFQName(namespace, "http_client", "send_latency_ms"),
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
	)
	HttpClientConnSendSuccessCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: prometheus.BuildFQName(namespace, "http_client_conn", "send_success_total")},
		[]string{"client_addr"},
	)
	HttpClientConnSendFailCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: prometheus.BuildFQName(namespace, "http_client_conn", "send_fail_total")},
		[]string{"client_addr"},
	)
	HttpClientConnSendSuccessLatency = promauto.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       prometheus.BuildFQName(namespace, "http_client_conn", "send_latency_ms"),
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001}},
		[]string{"client_addr"},
	)
)
