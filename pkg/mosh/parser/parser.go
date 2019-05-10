package parser

import (
	internals "../../../internal/mosh/parser"
	"runtime"
)

// Action covers all wrapper types for supported Parser::Action sub-classes
type Action interface {
	Intern() internals.Action
	Name() string
	Ignore() bool
}

type Resize struct {
	wrapped internals.Resize
}

func (r *Resize) Intern() internals.Action {
	return r.wrapped
}

func (r *Resize) Name() string {
	return r.wrapped.Name()
}

func (r *Resize) Ignore() bool {
	return r.wrapped.Ignore()
}

func MakeResize(width, height int64) *Resize {
	wrapped := internals.NewResize(width, height)
	r := &Resize{
		wrapped: wrapped,
	}
	runtime.SetFinalizer(r, func(r *Resize) {
		internals.DeleteResize(r.wrapped)
	})
	return r
}

type UserByte struct {
	wrapped internals.UserByte
}

func (ub *UserByte) Intern() internals.Action {
	return ub.wrapped
}

func (ub *UserByte) Name() string {
	return ub.wrapped.Name()
}

func (ub *UserByte) Ignore() bool {
	return ub.wrapped.Ignore()
}

func MakeUserByte(b int) *UserByte {
	wrapped := internals.NewUserByte(b)
	ub := &UserByte{
		wrapped: wrapped,
	}
	runtime.SetFinalizer(ub, func(ub *UserByte) {
		internals.DeleteUserByte(ub.wrapped)
	})
	return ub
}
