/*
   Copyright 2015 W. Max Lees

   This file is part of jarvisos.

   Jarvisos is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   Jarvisos is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with jarvisos.  If not, see <http://www.gnu.org/licenses/>.

   File: app.go
   Author: W. Max Lees <max.lees@gmail.com>
   Date: 06.14.2015
*/

package nlpclient

import (
	"net"
	"net/rpc"
	"time"
)

type AppServerClient struct {
	connection *rpc.Client
}

func (c *AppServerClient) DirectCall(call string) (*[]byte, error) {
	var result *[]byte
	err := c.connection.Call("AppServer.DirectCall", call, &result)
	return result, err
}

func NewClient(dsn string, timeout time.Duration) (*AppServerClient, error) {
	connection, err := net.DialTimeout("tcp", dsn, timeout)
	if err != nil {
		return nil, err
	}
	return &AppServerClient{connection: rpc.NewClient(connection)}, nil
}
