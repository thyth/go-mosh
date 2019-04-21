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

// TODO get_fb (callee vs caller owner?)

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

// TODO wrap Terminal::Complete

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

// TODO wrap Terminal::Framebuffer

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
