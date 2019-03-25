// SWIG Specification for Go Bindings to Mosh
//
// This only exposes a subset of the Mosh library as necessary for downstream consumers of this library.
//
// Specifically, allowing use of the following C++ classes:
// - Overlay::PredictionEngine
// - Parser::Action
// - Parser::Resize
// - Parser::UserByte
// - Terminal::Complete
// - Terminal::Display
// - Terminal::Framebuffer

%module mosh

// C++ Standard Library Includes
%include <std_deque.i>
%include <std_list.i>
%include <std_string.i>
%include <std_vector.i>

// %include <std_wstring.i>
// Does not exist for SWIG Go, but only used in Overlay::{NotificationEngine, TitleEngine, OverlayManager}, which we
// don't particularly care about, so suppressing the warning.
%warnfilter(315) wstring;

// --- namespace Terminal ---
%rename (equal) Terminal::Emulator::operator==;
%include "../../mosh/src/terminal/terminal.h"
%{
    #include <wchar.h>
    #include <stdio.h>
    #include <vector>
    #include <deque>

    #include "parseraction.h"
    #include "terminalframebuffer.h"
    #include "terminaldispatcher.h"
    #include "terminaluserinput.h"
    #include "terminaldisplay.h"
%}

// --- namespace Crypto ---
%rename (assign) Crypto::Session::operator=;
%include "../../mosh/src/crypto/crypto.h"
%{
    #include "ae.h"
    #include <string>
    #include <string.h>
    #include <stdint.h>
    #include <stdlib.h>
    #include <exception>
%}

// --- namespace Network ---
%include "../../mosh/src/network/network.h"
%{
    #include <stdint.h>
    #include <deque>
    #include <sys/socket.h>
    #include <netinet/in.h>
    #include <string>
    #include <math.h>
    #include <vector>
    #include <assert.h>
    #include <exception>
    #include <string.h>

    #include "crypto.h"
%}

// --- Overlay::PredictionEngine ---
%include "../../mosh/src/frontend/terminaloverlay.h"
%{
    #include "terminalframebuffer.h"
    #include "network.h"
    #include "transportsender.h"
    #include "parser.h"

    #include <vector>
    #include <limits.h>
%}

// --- Parser::{Action, Resize, UserByte} ---
%rename (equal) Parser::Resize::operator==;
%rename (equal) Parser::UserByte::operator==;
%include "../../mosh/src/terminal/parseraction.h"
%{
    #include <string>
    #include <vector>

    #include "shared.h"
%}

// --- Terminal::Complete ---
%rename (equal) Terminal::Complete::operator==;
%include "../../mosh/src/statesync/completeterminal.h"
%{
    #include <list>
    #include <stdint.h>

    #include "parser.h"
    #include "terminal.h"
%}

// --- Terminal::Display ---
%include "../../mosh/src/terminal/terminaldisplay.h"
%{
    #include "terminalframebuffer.h"
%}

// --- Terminal::Framebuffer ---
%rename (equal) Terminal::Cell::operator==;
%rename (notEqual) Terminal::Cell::operator!=;
%rename (equal) Terminal::DrawState::operator==;
%rename (assign) Terminal::Framebuffer::operator=;
%rename (equal) Terminal::Framebuffer::operator==;
%rename (equal) Terminal::Renditions::operator==;
%rename (equal) Terminal::Row::operator==;
%warnfilter(516) get_renditions; // a pair of declarations overloaded only by const pointer types -- ignore them
//%include "../../mosh/src/util/shared.h" // for shared::{make_shared, shared_ptr}
// Some SWIG backends support C++ smart pointers (see: http://www.swig.org/Doc3.0/Library.html#Library_std_shared_ptr),
// but Golang backend does not. These smart pointers are mostly internal implementation details for Mosh, but they are
// exposed in a few places. Renaming the exposure points with "unsafe" to alert to unintentional uses and suppressing
// warning messages. If the SWIG backend for Go gains support for smart pointers in the future, revisit this.
%rename (__unsafe__get_rows) Terminal::Framebuffer::get_rows;
%rename (__unsafe__get_row) Terminal::Framebuffer::get_row;
%rename (__unsafe__get_mutable_row) Terminal::Framebuffer::get_mutable_row;
%warnfilter(315) shared_ptr;
%warnfilter(315) make_shared;
%include "../../mosh/src/terminal/terminalframebuffer.h"
%{
    #include <assert.h>
    #include <limits.h>
    #include <stdint.h>

    #include <vector>
    #include <deque>
    #include <string>
    #include <list>

    #include "shared.h"
%}
