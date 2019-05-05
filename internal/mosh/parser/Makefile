# this is a pseudo-Makefile intended to capture all of the include and linking arguments eventually used by cgo

parser.go:
	swig -go -cgo -c++ -intgosize 64 parser.i

parser: parser.go
	c++ -arch x86_64 -std=c++11 -I ../../../mosh/ -I ../../../mosh/src/crypto -I ../../../mosh/src/frontend -I ../../../mosh/src/network -I ../../../mosh/src/protobufs -I ../../../mosh/src/statesync -I ../../../mosh/src/terminal -I ../../../mosh/src/util parser_wrap.cxx ../../../mosh/src/terminal/parser.o ../../../mosh/src/terminal/parseraction.o ../../../mosh/src/terminal/parserstate.o ../../../mosh/src/terminal/terminaldispatcher.o ../../../mosh/src/terminal/terminal.o ../../../mosh/src/terminal/terminalframebuffer.o ../../../mosh/src/terminal/terminaluserinput.o ../../../mosh/src/terminal/terminalfunctions.o

clean:
	rm parser.go parser_wrap.cxx

all: parser
