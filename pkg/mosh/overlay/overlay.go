/*
 * go-mosh: mosh SWIG wrapper for Golang
 * Copyright 2019-2023 Daniel Selifonov
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

package overlay

import (
	internals "gitlab.hive.thyth.com/chronostruct/go-mosh/internal/mosh"
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

func (pe *PredictionEngine) LocalFrameSent(epoch uint64) {
	pe.wrapped.Set_local_frame_sent(epoch)
}

func (pe *PredictionEngine) LocalFrameAcked(epoch uint64) {
	pe.wrapped.Set_local_frame_acked(epoch)
}

func (pe *PredictionEngine) LocalFrameLateAcked(epoch uint64) {
	pe.wrapped.Set_local_frame_late_acked(epoch)
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
