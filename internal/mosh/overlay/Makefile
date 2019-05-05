# this is a pseudo-Makefile intended to capture all of the include and linking arguments eventually used by cgo

overlay.go:
	swig -go -cgo -c++ -intgosize 64 overlay.i

overlay: overlay.go
	c++ -arch x86_64 -std=c++11 -I ../../../mosh/ -I ../../../mosh/src/crypto -I ../../../mosh/src/frontend -I ../../../mosh/src/network -I ../../../mosh/src/protobufs -I ../../../mosh/src/statesync -I ../../../mosh/src/terminal -I ../../../mosh/src/util overlay_wrap.cxx

clean:
	rm overlay.go overlay_wrap.cxx

all: overlay
