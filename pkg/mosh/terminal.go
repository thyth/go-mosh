package mosh

import (
	internals "../../internal/mosh"
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
	runtime.SetFinalizer(emu, func (emu *Complete) {
		internals.DeleteComplete(emu.wrapped)
	})
	return emu
}

func (emu *Complete) Act(action Action) string {
	effect := emu.wrapped.Act(action.intern())
	runtime.KeepAlive(action)
	return effect
}

func (emu *Complete) Perform(s string) string {
	return emu.wrapped.Act(s)
}

func (emu *Complete) GetFramebuffer() *Framebuffer {
	internFb := emu.wrapped.Get_fb()
	fb := &Framebuffer{
		wrapped: internFb,
	}
	// TODO callee vs caller owner? i.e. should this Framebuffer be tracked as part of the *Complete instead?
	runtime.SetFinalizer(fb, func (fb *Framebuffer) {
		internals.DeleteFramebuffer(fb.wrapped)
	})
	return fb
}

func (emu *Complete) ResetInput() {
	emu.wrapped.Reset_input()
}

func (emu *Complete) GetEchoAck() time.Time {
	// create a Go time.Time from a millisecond epoch time -- to millisecond accuracy
	milliseconds := int64(emu.wrapped.Get_echo_ack())
	seconds := milliseconds / int64(time.Second / time.Millisecond)
	nanos := (milliseconds % int64(time.Second / time.Millisecond)) * int64(time.Millisecond / time.Nanosecond)
	return time.Unix(seconds, nanos)
}

func (emu *Complete) SetEchoAck(t time.Time) bool {
	return emu.wrapped.Set_echo_ack(timestamp(t))
}

func (emu *Complete) RegisterInputFrame(n uint64, now time.Time) {
	emu.wrapped.Register_input_frame(n, timestamp(now))
}

func (emu *Complete) WaitTime(now time.Time) time.Duration {
	// the integer returned here seems to be a duration (in seconds)
	return time.Second * time.Duration(emu.wrapped.Wait_time(timestamp(now)))
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
	eq := emu.wrapped.Equal(other.wrapped)
	runtime.KeepAlive(other)
	return eq
}

func (emu *Complete) Compare(other *Complete) bool {
	cmp := emu.wrapped.Compare(other.wrapped)
	runtime.KeepAlive(other)
	return cmp
}

type Framebuffer struct {
	wrapped internals.Framebuffer
}

func MakeFramebuffer(width, height int) *Framebuffer {
	wrapped := internals.NewFramebuffer(width, height)
	fb := &Framebuffer{
		wrapped: wrapped,
	}
	runtime.SetFinalizer(fb, func (fb *Framebuffer) {
		internals.DeleteFramebuffer(fb.wrapped)
	})
	return fb
}

func CopyFramebuffer(other *Framebuffer) *Framebuffer {
	wrapped := internals.NewFramebuffer(other.wrapped)
	runtime.KeepAlive(other)
	fb := &Framebuffer{
		wrapped: wrapped,
	}
	runtime.SetFinalizer(fb, func (fb *Framebuffer) {
		internals.DeleteFramebuffer(fb.wrapped)
	})
	return fb
}

// note: skipping reference assignment operator `Framebuffer &operator=( const Framebuffer &other );`
// note: skipping access to DrawState field reference
// note: skipping get_rows getter, since this is returning smart pointer types

func (fb *Framebuffer) Scroll(n int) {
	fb.wrapped.Scroll(n)
}

func (fb *Framebuffer) MoveRowsAutoscroll(rows int) {
	fb.wrapped.Move_rows_autoscroll(rows)
}

// note: skipping `Row *get_row( int row )`
// note: skipping `Cell *get_cell( int row = -1, int col = -1 )`
// note: skipping `Row *get_mutable_row( int row )`
// note: skipping `Cell *get_mutable_cell( int row = -1, int col = -1 )`
// note: skipping `Cell *get_combining_cell( void )`
// note: skipping `void apply_renditions_to_cell( Cell *cell )`

func (fb *Framebuffer) InsertLine(beforeRow, count int) {
	fb.wrapped.Insert_line(beforeRow, count)
}

func (fb *Framebuffer) DeleteLine(row, count int) {
	fb.wrapped.Delete_line(row, count)
}

func (fb *Framebuffer) InsertCell(row, col int) {
	fb.wrapped.Insert_cell(row, col)
}

func (fb *Framebuffer) DeleteCell(row, col int) {
	fb.wrapped.Delete_cell(row, col)
}

func (fb *Framebuffer) Reset() {
	fb.wrapped.Reset()
}

func (fb *Framebuffer) SoftReset() {
	fb.wrapped.Soft_reset()
}

func (fb *Framebuffer) InitializeTitle() {
	fb.wrapped.Set_title_initialized()
}

func (fb *Framebuffer) IsTitleInitialized() bool {
	return fb.wrapped.Is_title_initialized()
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
	fb.wrapped.Resize(width, height)
}

// note: skipping the following due to exposure of Terminal::Cell and Terminal::Row
// - void reset_cell( Cell *c )
// - void reset_row( Row *r )

func (fb *Framebuffer) RingBell() {
	fb.wrapped.Ring_bell()
}

func (fb *Framebuffer) BellCount() uint {
	return fb.wrapped.Get_bell_count()
}

func (fb *Framebuffer) Equals(other *Framebuffer) bool {
	equals := fb.wrapped.Equal(other.wrapped)
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
	frame := disp.wrapped.New_frame(initialized, last.wrapped, f.wrapped)
	runtime.KeepAlive(last)
	runtime.KeepAlive(f)
	return frame
}
