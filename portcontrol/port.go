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

   File: port.go
   Author: W. Max Lees <max.lees@gmail.com>
   Date: 06.16.2015
*/

package portcontrol

import (
	"fmt"
	"strconv"
)

type port string

var ports [255]bool

const offset = 7492

func Generate() (port, error) {
	for i := 0; i < 255; i++ {
		if !ports[i] {
			ports[i] = true
			return port(strconv.Itoa(i + offset)), nil
		}
	}

	return port(""), PortError("No more free jarvis ports")
}

func Free(aPort port) {
	intPort, err := strconv.Atoi(string(aPort))
	if err != nil {
		fmt.Printf("Cannot free port %v: %v\n", aPort, err)
	}
	ports[intPort-offset] = false
}
