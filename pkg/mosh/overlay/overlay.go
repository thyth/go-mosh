package overlay

import (
	internals "gitlab.hive.thyth.com/chronostruct/go-mosh/internal/mosh"
	util "gitlab.hive.thyth.com/chronostruct/go-mosh/pkg/mosh"
	"gitlab.hive.thyth.com/chronostruct/go-mosh/pkg/mosh/terminal"

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

func (pe *PredictionEngine) Apply(fb *terminal.Framebuffer) {
	pe.wrapped.Apply(fb.Wrapped)
	runtime.KeepAlive(fb)
}

func (pe *PredictionEngine) NewUserByte(c byte, fb *terminal.Framebuffer) {
	pe.wrapped.New_user_byte(c, fb.Wrapped)
	runtime.KeepAlive(fb)
}

func (pe *PredictionEngine) Cull(fb *terminal.Framebuffer) {
	pe.wrapped.Cull(fb.Wrapped)
	runtime.KeepAlive(fb)
}

func (pe *PredictionEngine) Reset() {
	pe.wrapped.Reset()
}

func (pe *PredictionEngine) LocalFrameSent(t time.Time) {
	pe.wrapped.Set_local_frame_sent(util.Timestamp(t))
}

func (pe *PredictionEngine) LocalFrameAcked(t time.Time) {
	pe.wrapped.Set_local_frame_acked(util.Timestamp(t))
}

func (pe *PredictionEngine) LocalFrameLateAcked(t time.Time) {
	pe.wrapped.Set_local_frame_late_acked(util.Timestamp(t))
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
