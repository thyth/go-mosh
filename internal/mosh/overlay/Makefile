# this is a pseudo-Makefile intended to capture all of the include and linking arguments eventually used by cgo

overlay.go:
	swig -go -cgo -c++ -intgosize 64 overlay.i

overlay: overlay.go
	c++ -arch x86_64 -std=c++11 -shared -o liboverlay.so -I ../../../mosh/ -I ../../../mosh/src/crypto -I ../../../mosh/src/frontend -I ../../../mosh/src/network -I ../../../mosh/src/protobufs -I ../../../mosh/src/statesync -I ../../../mosh/src/terminal -I ../../../mosh/src/util overlay_wrap.cxx ../../../mosh/src/terminal/parser.o ../../../mosh/src/terminal/parseraction.o ../../../mosh/src/terminal/parserstate.o ../../../mosh/src/frontend/terminaloverlay.o ../../../mosh/src/terminal/terminaldispatcher.o ../../../mosh/src/terminal/terminalframebuffer.o ../../../mosh/src/terminal/terminal.o ../../../mosh/src/terminal/terminaluserinput.o ../../../mosh/src/terminal/terminalfunctions.o ../../../mosh/src/network/network.o ../../../mosh/src/util/timestamp.o ../../../mosh/src/crypto/crypto.o ../../../mosh/src/crypto/base64.o ../../../mosh/src/crypto/ocb.o

clean:
	rm overlay.go overlay_wrap.cxx

all: overlay
