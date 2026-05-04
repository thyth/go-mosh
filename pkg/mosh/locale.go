/*
 * go-mosh: mosh SWIG wrapper for Golang
 * Copyright 2025-2026 Daniel Selifonov
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
#include <langinfo.h>
#include <strings.h>

static int codeset_is_utf8(void) {
    const char *cs = nl_langinfo(CODESET);
    return cs != 0 &&
        (strcasecmp(cs, "UTF-8") == 0 ||
         strcasecmp(cs, "UTF8") == 0);
}

static const char* init_utf8_locale(void) {
    const char *r;
    r = setlocale(LC_CTYPE, "");
    if (r != 0 && codeset_is_utf8()) return r;
    r = setlocale(LC_CTYPE, "C.UTF-8");
    if (r != 0 && codeset_is_utf8()) return r;
    r = setlocale(LC_CTYPE, "C.utf8");
    if (r != 0 && codeset_is_utf8()) return r;
    r = setlocale(LC_CTYPE, "en_US.UTF-8");
    if (r != 0 && codeset_is_utf8()) return r;
    r = setlocale(LC_CTYPE, "UTF-8");
    if (r != 0 && codeset_is_utf8()) return r;

    return NULL;
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
	//
	// We try a few different ways to discover a UTF-8 capable locale, in case "" itself fails to produce a UTF-8
	// enabled locale. We explicitly panic on initialization instead of proceeding with silent UTF-8 glyph drops.
	if locale := C.init_utf8_locale(); locale == nil {
		panic("failed to initialize a UTF-8 LC_CTYPE for embedded mosh")
	}
}
