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

   File: appserver.go
   Author: W. Max Lees <max.lees@gmail.com>
   Date: 06.13.2015
*/

package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type AppServer struct{}

func main() {
	appServer := AppServer{}
	rpc.Register(&appServer)

	// Listen on the specified port
	l, err := net.Listen("tcp", "localhost:7491")
	if err != nil {
		fmt.Printf("Error initializing app server on port 7491: %v\n", err)
		return
	}

	// Listen for calls
	rpc.Accept(l)

}

func (app *AppServer) DirectCall(call string, result *[]byte) error {
	fmt.Printf("Direct Call: %v\n", call)

	str := []byte("Success")
	result = &str

	return nil
}
