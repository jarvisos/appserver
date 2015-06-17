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

   File: porterror.go
   Author: W. Max Lees <max.lees@gmail.com>
   Date: 06.16.2015
*/

package portcontrol

import (
	"fmt"
	"time"
)

type portError struct {
	When time.Time
	What string
}

func (e portError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func PortError(err string) portError {
	return portError{
		time.Now(),
		err,
	}
}
