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

import "fmt"

type ProtocolType int

const (
	ProtocolTypeUdp ProtocolType = iota
	ProtocolTypeTcp
	ProtocolTypeHttp
)

type CommType int

const (
	CommTypeClient CommType = iota
	CommTypeServer
)

type Config struct {
	ProtocolType             ProtocolType
	CommType                 CommType
	PrometheusMetricsDisable bool
}

type CommonConfig struct {
	Host string
	Port int
}

type ClientConfig struct {
	CommonConfig  CommonConfig
	ConnNum       int
	TickPerConnMs int
	PacketSize    int
}

func (c *ClientConfig) addr() string {
	return fmt.Sprintf("%s:%d", c.CommonConfig.Host, c.CommonConfig.Port)
}

type ServerConfig struct {
	CommonConfig CommonConfig
}

func (s *ServerConfig) addr() string {
	return fmt.Sprintf("%s:%d", s.CommonConfig.Host, s.CommonConfig.Port)
}
