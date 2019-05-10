%module parser

%apply unsigned int { wchar_t }
%include <std_string.i>

// ignore all classes, except for Parser::{Action, Resize, UserByte}
%rename($ignore, %$isclass) "";
%rename("%s") Action;
%rename("%s") Resize;
%rename("%s") UserByte;

%rename (equals) Parser::Resize::operator==;
%rename (equals) Parser::UserByte::operator==;

%include "../../../mosh/src/terminal/parseraction.h"
%{
    #include "../../../mosh/src/terminal/parseraction.h"
%}
