/*
 * go-mosh: mosh SWIG wrapper for Golang
 * Copyright 2025 Daniel Selifonov
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

/*
#include <locale.h>
static void init_utf8_locale(void) {
	setlocale(LC_ALL, "en_US.UTF-8");
}
*/
import "C"

func init() {
	// Go does not use libc system locale, thus defaults to the equivalent of `setlocale(LC_ALL, "C")`. The ordinary
	// startup lifecycles of both mosh-client and mosh-server do result in a `setlocale(LC_ALL, "")` to initialize a
	// UTF-8 enabled locale.
	//
	// See: https://github.com/mobile-shell/mosh/blob/1105d481bb9143dad43adf768f58da7b029fd39c/src/frontend/mosh-client.cc#L193
	//
	// We need to do the same since "C" locale results in the conversion loss of non-ASCII UTF-8 glyphs.
	C.init_utf8_locale()
}
