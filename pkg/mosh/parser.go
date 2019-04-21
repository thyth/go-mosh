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

// TODO wrap Parser::Resize

type UserByte struct {
	wrapped internals.UserByte
}

// TODO wrap Parser::UserByte
