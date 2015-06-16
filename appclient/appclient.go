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

   File: appclient.go
   Author: W. Max Lees <max.lees@gmail.com>
   Date: 06.15.2015
*/

package appclient

import (
	"net"
	"net/rpc"
	"time"
)

type AppClient struct {
	connection *rpc.Client
}

func (c *AppClient) Call(call string) (*[]byte, error) {
	var result *[]byte
	err := c.connection.Call("App.Call", call, &result)
	return result, err
}

func (c *AppClient) Who(full bool) (*[]byte, error) {
	var result *[]byte
	err := c.connection.Call("App.Who", full, &result)
	return result, err
}

func NewClient(dsn string, timeout time.Duration) (*AppClient, error) {
	connection, err := net.DialTimeout("tcp", dsn, timeout)
	if err != nil {
		return nil, err
	}
	return &AppClient{connection: rpc.NewClient(connection)}, nil
}
