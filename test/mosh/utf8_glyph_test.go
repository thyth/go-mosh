/*
 * go-mosh: mosh SWIG wrapper for Golang
 * Copyright 2026 Daniel Selifonov
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
	"testing"

	_ "gitlab.hive.thyth.com/chronostruct/go-mosh/pkg/mosh"
	"gitlab.hive.thyth.com/chronostruct/go-mosh/pkg/mosh/terminal"
)

// TestAsciiGlyphRoundtrip sanity checks that the terminal emulator accepts a simple ASCII stdout, and we can validate
// the contents of cell 0, 0 match that ASCII character.
func TestAsciiGlyphRoundtrip(t *testing.T) {
	a := "a"
	emulator := terminal.MakeComplete(50, 50)
	emulator.Perform(a)
	fb := emulator.GetFramebuffer()
	cell := fb.GetCell(0, 0)
	grapheme := cell.PrintGrapheme()
	if grapheme != a {
		t.Errorf("Unexpected cell contents: '%v'\n\tdebug: %s", grapheme, cell.DebugContents())
	}
}

// TestUtf8GlyphRoundtrip validates that we can round-trip UTF-8 glyphs outside the ASCII range.
func TestUtf8GlyphRoundtrip(t *testing.T) {
	unicodeSmilingFace := "\U0001F60A"
	emulator := terminal.MakeComplete(50, 50)
	emulator.Perform(unicodeSmilingFace)
	fb := emulator.GetFramebuffer()
	cell := fb.GetCell(0, 0)
	grapheme := cell.PrintGrapheme()
	if grapheme != unicodeSmilingFace {
		t.Errorf("Unexpected cell contents: '%v'\n\tdebug: %s", grapheme, cell.DebugContents())
	}
}
