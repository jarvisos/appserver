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
	"github.com/jarvisos/app/nlpclient"
	"github.com/jarvisos/appserver/portcontrol"
	"net"
	"net/rpc"
	"os/exec"
	"strings"
	"time"
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

	fmt.Printf("Listening on port 7491\n")

	// Listen for calls
	rpc.Accept(l)

}

func (app *AppServer) DirectCall(call string, result *[]byte) error {
	fmt.Printf("Direct Call: %v\n", call)

	port, err := portcontrol.Generate()
	if err != nil {
		fmt.Printf("Could not generate a port: %v\n", err)
		return err
	}

	parts := strings.Split(strings.Trim(call, " \n"), " ")
	command := exec.Command(parts[0], "-p="+string(port))
	error := command.Start()
	if error != nil {
		fmt.Printf("Error starting '%s -p=%s': %v\n", parts[0], string(port), error)
		portcontrol.Free(port)
		return error
	}

	time.Sleep(time.Second)
	time.Sleep(time.Second)

	clientApp, err := nlpclient.NewClient("localhost:"+string(port), time.Second)
	if err != nil {
		fmt.Printf("Error connecting to app: %v\n", err)
		portcontrol.Free(port)
		return err
	}

	result, err = clientApp.Call("test")
	if err != nil {
		fmt.Printf("Error calling function for client app: %v\n", err)
		return err
	}

	str := []byte("Success")
	result = &str

	portcontrol.Free(port)

	return nil
}
