// SWIG Specification for Go Bindings to Mosh

%module mosh

%include "../../mosh/src/terminal/parser.h"
%include "../../mosh/src/terminal/parseraction.h"
%include "../../mosh/src/terminal/parserstate.h"
%include "../../mosh/src/terminal/parserstatefamily.h"
%include "../../mosh/src/terminal/parsertransition.h"

%include "../../mosh/src/terminal/terminal.h"
%include "../../mosh/src/terminal/terminaldispatcher.h"
%include "../../mosh/src/terminal/terminaldisplay.h"
%include "../../mosh/src/terminal/terminalframebuffer.h"
%include "../../mosh/src/terminal/terminaluserinput.h"
