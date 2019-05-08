# this is a pseudo-Makefile intended to capture all of the include and linking arguments eventually used by cgo

terminal.go:
	swig -go -cgo -c++ -intgosize 64 terminal.i

terminal: terminal.go
	c++ -arch x86_64 -std=c++11 -shared -o libterminal.so -I ../../../mosh/ -I ../../../mosh/src/terminal -I ../../../mosh/src/util -lncurses -lprotobuf terminal_wrap.cxx ../../../mosh/src/terminal/parser.o ../../../mosh/src/terminal/parseraction.o ../../../mosh/src/terminal/parserstate.o ../../../mosh/src/terminal/terminalframebuffer.o ../../../mosh/src/terminal/terminaldisplay.o ../../../mosh/src/terminal/terminaldisplayinit.o ../../../mosh/src/statesync/completeterminal.o ../../../mosh/src/terminal/terminal.o ../../../mosh/src/terminal/terminaldispatcher.o ../../../mosh/src/terminal/terminaluserinput.o ../../../mosh/src/terminal/terminalfunctions.o ../../../mosh/src/protobufs/hostinput.pb.o

clean:
	rm terminal.go terminal_wrap.cxx

all: terminal
