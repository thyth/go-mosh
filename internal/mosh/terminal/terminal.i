%module terminal

%apply unsigned long long { uint64_t }
%include <std_string.i>
%include <std_list.i>
%include <std_vector.i>

// ignore all classes, except for Terminal::{Complete, Display, Framebuffer}
%rename($ignore, %$isclass) "";
%rename("%s") Complete;
%rename("%s") Display;
%rename("%s") Framebuffer;

%rename (equals) Terminal::Complete::operator==;

%rename (equals) Terminal::Cell::operator==;
%ignore Terminal::Cell::operator!=;
%rename (equals) Terminal::DrawState::operator==;
%rename (assign) Terminal::Framebuffer::operator=;
%rename (equals) Terminal::Framebuffer::operator==;
%rename (equals) Terminal::Renditions::operator==;
%rename (equals) Terminal::Row::operator==;
%warnfilter(516) get_renditions; // a pair of declarations overloaded only by const pointer types -- ignore them
%ignore Terminal::Framebuffer::get_rows;
%ignore Terminal::Framebuffer::get_row;
%ignore Terminal::Framebuffer::get_mutable_row;
%warnfilter(315) shared_ptr;
%warnfilter(315) make_shared;

%include "../../../mosh/src/terminal/terminalframebuffer.h"
%include "../../../mosh/src/terminal/terminaldisplay.h"
%include "../../../mosh/src/statesync/completeterminal.h"

%{
    #include "../../../mosh/src/terminal/terminalframebuffer.h"
    #include "../../../mosh/src/terminal/terminaldisplay.h"
    #include "../../../mosh/src/statesync/completeterminal.h"
%}
