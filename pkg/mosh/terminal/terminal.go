/*
 * go-mosh: mosh SWIG wrapper for Golang
 * Copyright 2019-2022 Daniel Selifonov
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

package terminal

import (
	internals "gitlab.hive.thyth.com/chronostruct/go-mosh/internal/mosh"
	util "gitlab.hive.thyth.com/chronostruct/go-mosh/pkg/mosh"
	"gitlab.hive.thyth.com/chronostruct/go-mosh/pkg/mosh/parser"

	"runtime"
	"time"
)

type Complete struct {
	wrapped internals.Complete
}

func MakeComplete(width, height int) *Complete {
	wrapped := internals.NewComplete(int64(width), int64(height))
	emu := &Complete{
		wrapped: wrapped,
	}
	runtime.SetFinalizer(emu, func(emu *Complete) {
		internals.DeleteComplete(emu.wrapped)
	})
	return emu
}

func (emu *Complete) Act(action parser.Action) string {
	effect := emu.wrapped.Act(action.Intern())
	runtime.KeepAlive(action)
	return effect
}

func (emu *Complete) Perform(s string) string {
	return emu.wrapped.Act(s)
}

func (emu *Complete) GetFramebuffer() *Framebuffer {
	// this returns a const reference to an internal framebuffer instance --
	// this instance is created and collected independently of this reference; library consumers should be careful not
	// to hold onto these references for longer than necessary, in case the parent reference is finalized
	internFb := emu.wrapped.Get_fb()
	fb := &Framebuffer{
		Wrapped: internFb,
	}
	return fb
}

func (emu *Complete) ResetInput() {
	emu.wrapped.Reset_input()
}

func (emu *Complete) GetEchoAck() time.Time {
	// create a Go time.Time from a millisecond epoch time -- to millisecond accuracy
	milliseconds := int64(emu.wrapped.Get_echo_ack())
	seconds := milliseconds / int64(time.Second/time.Millisecond)
	nanos := (milliseconds % int64(time.Second/time.Millisecond)) * int64(time.Millisecond/time.Nanosecond)
	return time.Unix(seconds, nanos)
}

func (emu *Complete) SetEchoAck(t time.Time) bool {
	return emu.wrapped.Set_echo_ack(util.Timestamp(t))
}

func (emu *Complete) RegisterInputFrame(n uint64, now time.Time) {
	emu.wrapped.Register_input_frame(n, util.Timestamp(now))
}

func (emu *Complete) WaitTime(now time.Time) time.Duration {
	// the integer returned here seems to be a duration (in seconds)
	return time.Second * time.Duration(emu.wrapped.Wait_time(util.Timestamp(now)))
}

func (emu *Complete) Subtract(other *Complete) {
	emu.wrapped.Subtract(other.wrapped)
	runtime.KeepAlive(other)
}

func (emu *Complete) DiffFrom(existing *Complete) string {
	diff := emu.wrapped.Diff_from(existing.wrapped)
	runtime.KeepAlive(existing)
	return diff
}

func (emu *Complete) InitDiff() string {
	return emu.wrapped.Init_diff()
}

func (emu *Complete) ApplyString(diff string) {
	emu.wrapped.Apply_string(diff)
}

func (emu *Complete) Equals(other *Complete) bool {
	eq := emu.wrapped.Equals(other.wrapped)
	runtime.KeepAlive(other)
	return eq
}

func (emu *Complete) Compare(other *Complete) bool {
	cmp := emu.wrapped.Compare(other.wrapped)
	runtime.KeepAlive(other)
	return cmp
}

type Framebuffer struct {
	Wrapped internals.Framebuffer
}

func MakeFramebuffer(width, height int) *Framebuffer {
	wrapped := internals.NewFramebuffer(width, height)
	fb := &Framebuffer{
		Wrapped: wrapped,
	}
	runtime.SetFinalizer(fb, func(fb *Framebuffer) {
		internals.DeleteFramebuffer(fb.Wrapped)
	})
	return fb
}

func CopyFramebuffer(other *Framebuffer) *Framebuffer {
	wrapped := internals.NewFramebuffer(other.Wrapped)
	runtime.KeepAlive(other)
	fb := &Framebuffer{
		Wrapped: wrapped,
	}
	runtime.SetFinalizer(fb, func(fb *Framebuffer) {
		internals.DeleteFramebuffer(fb.Wrapped)
	})
	return fb
}

// note: skipping reference assignment operator `Framebuffer &operator=( const Framebuffer &other );`
// note: skipping access to DrawState field reference
// note: skipping get_rows getter, since this is returning smart pointer types

func (fb *Framebuffer) Scroll(n int) {
	fb.Wrapped.Scroll(n)
}

func (fb *Framebuffer) MoveRowsAutoscroll(rows int) {
	fb.Wrapped.Move_rows_autoscroll(rows)
}

// note: skipping `Row *get_row( int row )`
// note: skipping `Cell *get_cell( int row = -1, int col = -1 )`
// note: skipping `Row *get_mutable_row( int row )`
// note: skipping `Cell *get_mutable_cell( int row = -1, int col = -1 )`
// note: skipping `Cell *get_combining_cell( void )`
// note: skipping `void apply_renditions_to_cell( Cell *cell )`

func (fb *Framebuffer) InsertLine(beforeRow, count int) {
	fb.Wrapped.Insert_line(beforeRow, count)
}

func (fb *Framebuffer) DeleteLine(row, count int) {
	fb.Wrapped.Delete_line(row, count)
}

func (fb *Framebuffer) InsertCell(row, col int) {
	fb.Wrapped.Insert_cell(row, col)
}

func (fb *Framebuffer) DeleteCell(row, col int) {
	fb.Wrapped.Delete_cell(row, col)
}

func (fb *Framebuffer) Reset() {
	fb.Wrapped.Reset()
}

func (fb *Framebuffer) SoftReset() {
	fb.Wrapped.Soft_reset()
}

func (fb *Framebuffer) InitializeTitle() {
	fb.Wrapped.Set_title_initialized()
}

func (fb *Framebuffer) IsTitleInitialized() bool {
	return fb.Wrapped.Is_title_initialized()
}

// note: skipping the following due to use of std::vector<wchar_t>
// - void set_icon_name( const title_type &s )
// - void set_window_title( const title_type &s )
// - void set_clipboard( const title_type &s )
// - const title_type & get_icon_name( void )
// - const title_type & get_window_title( void )
// - const title_type & get_clipboard( void )
// - void prefix_window_title( const title_type &s )

func (fb *Framebuffer) Resize(width, height int) {
	fb.Wrapped.Resize(width, height)
}

// note: skipping the following due to exposure of Terminal::Cell and Terminal::Row
// - void reset_cell( Cell *c )
// - void reset_row( Row *r )

func (fb *Framebuffer) RingBell() {
	fb.Wrapped.Ring_bell()
}

func (fb *Framebuffer) BellCount() uint {
	return fb.Wrapped.Get_bell_count()
}

func (fb *Framebuffer) Equals(other *Framebuffer) bool {
	equals := fb.Wrapped.Equals(other.Wrapped)
	runtime.KeepAlive(other)
	return equals
}

type Display struct {
	wrapped internals.Display
}

func MakeDisplay(useEnvironment bool) *Display {
	wrapped := internals.NewDisplay(useEnvironment)
	disp := &Display{
		wrapped: wrapped,
	}
	runtime.SetFinalizer(disp, func(disp *Display) {
		internals.DeleteDisplay(disp.wrapped)
	})
	return disp
}

func (disp *Display) Open() string {
	return disp.wrapped.Open()
}

func (disp *Display) Close() string {
	return disp.wrapped.Close()
}

func (disp *Display) NewFrame(initialized bool, last, f *Framebuffer) string {
	frame := disp.wrapped.New_frame(initialized, last.Wrapped, f.Wrapped)
	runtime.KeepAlive(last)
	runtime.KeepAlive(f)
	return frame
}
