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
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
)

func tcpServerRun(serverConfig ServerConfig) error {
	l, err := net.Listen("tcp", serverConfig.addr())
	if err != nil {
		return err
	}
	defer l.Close()
	logrus.Info("listen on ", serverConfig.addr())
	for {
		conn, err := l.Accept()
		if err != nil {
			logrus.Error("Error accepting ", err)
			break
		}
		go handleRequest(conn)
	}
	return nil
}

func handleRequest(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		reqLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		write, err := conn.Write(buf[:reqLen])
		if reqLen != write {
			logrus.Error("write not full success")
			break
		}
		if err != nil {
			logrus.Error("write data failed")
			break
		}
	}
}
