%module overlay

%apply unsigned long long { uint64_t }
%include <std_deque.i>
%include <std_list.i>
%include <std_vector.i>

%warnfilter(315) wstring;

// ignore all classes, except for the PredictionEngine
%rename($ignore, %$isclass) "";
%rename("%s") PredictionEngine;

// create a dummy designator of the Framebuffer in the Terminal namespace, so the correct C++ wrapper is generated
namespace Terminal {
    class Framebuffer { };
}

// suppress warning about the Network namespace. All uses of that namespace are in ignored classes.
namespace Network { }

%include "../../../mosh/src/frontend/terminaloverlay.h"
%{
    #include "../../../mosh/src/frontend/terminaloverlay.h"
%}
