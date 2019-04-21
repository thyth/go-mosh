package mosh

import (
	internals "../../internal/mosh"

	"runtime"
	"time"
)

type DisplayPreference int

const (
	PredictAlways DisplayPreference = iota
	PredictNever
	PredictAdaptive
	PredictExperimental
)

func (dp DisplayPreference) internal() internals.OverlayPredictionEngineDisplayPreference {
	switch dp {
	case PredictAlways:
		return internals.PredictionEngineAlways
	case PredictNever:
		return internals.PredictionEngineNever
	case PredictAdaptive:
		return internals.PredictionEngineAdaptive
	case PredictExperimental:
		return internals.PredictionEngineExperimental
	default:
		return internals.PredictionEngineNever
	}
}

type PredictionEngine struct {
	wrapped internals.PredictionEngine
}

func MakePredictionEngine() *PredictionEngine {
	wrapped := internals.NewPredictionEngine()
	pe := &PredictionEngine{
		wrapped: wrapped,
	}
	runtime.SetFinalizer(pe, func(pe *PredictionEngine) {
		internals.DeletePredictionEngine(pe.wrapped)
	})
	return pe
}

func (pe *PredictionEngine) SetDisplayPreference(dp DisplayPreference) {
	pe.wrapped.Set_display_preference(dp.internal())
}

func (pe *PredictionEngine) SetPredictOverwrite(overwrite bool) {
	pe.wrapped.Set_predict_overwrite(overwrite)
}

func (pe *PredictionEngine) Apply(fb *Framebuffer) {
	pe.wrapped.Apply(fb.wrapped)
	runtime.KeepAlive(fb)
}

func (pe *PredictionEngine) NewUserByte(c byte, fb *Framebuffer) {
	pe.wrapped.New_user_byte(c, fb.wrapped)
	runtime.KeepAlive(fb)
}

func (pe *PredictionEngine) Cull(fb *Framebuffer) {
	pe.wrapped.Cull(fb.wrapped)
	runtime.KeepAlive(fb)
}

func (pe *PredictionEngine) Reset() {
	pe.wrapped.Reset()
}

func timestamp(t time.Time) uint64 {
	return uint64(t.UnixNano() / int64(time.Millisecond/time.Nanosecond))
}

func (pe *PredictionEngine) LocalFrameSent(t time.Time) {
	pe.wrapped.Set_local_frame_sent(timestamp(t))
}

func (pe *PredictionEngine) LocalFrameAcked(t time.Time) {
	pe.wrapped.Set_local_frame_acked(timestamp(t))
}

func (pe *PredictionEngine) LocalFrameLateAcked(t time.Time) {
	pe.wrapped.Set_local_frame_late_acked(timestamp(t))
}

func (pe *PredictionEngine) SetSendInterval(duration time.Duration) {
	milliseconds := duration.Nanoseconds() / int64(time.Millisecond/time.Nanosecond)
	if milliseconds < 0 {
		milliseconds = 0
	}
	pe.wrapped.Set_send_interval(uint(milliseconds))
}

func (pe *PredictionEngine) WaitTime() time.Duration {
	milliseconds := pe.wrapped.Wait_time()
	return time.Millisecond * time.Duration(milliseconds)
}
