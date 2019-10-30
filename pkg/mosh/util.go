/*
 * go-mosh: mosh SWIG wrapper for Golang
 * Copyright 2019 Daniel Selifonov
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package mosh

import (
	internals "gitlab.hive.thyth.com/chronostruct/go-mosh/internal/mosh"

	"time"
)

func Timestamp(t time.Time) uint64 {
	return uint64(t.UnixNano() / int64(time.Millisecond/time.Nanosecond))
}

func GetVersion() string {
	return internals.BUILD_VERSION
}
