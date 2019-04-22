package mosh

import (
	internals "../../internal/mosh"
)

type Action interface {
	intern() internals.Action
}

// TODO wrap Parser::Action

type Resize struct {
	wrapped internals.Resize
}

func (r *Resize) intern() internals.Action {
	return r.wrapped
}

// TODO wrap Parser::Resize

type UserByte struct {
	wrapped internals.UserByte
}

func (ub *UserByte) intern() internals.Action {
	return ub.wrapped
}

// TODO wrap Parser::UserByte
