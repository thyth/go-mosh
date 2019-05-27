mosh.go:
	swig -go -cgo -c++ -intgosize 64 mosh.swigcxx

mosh: mosh.go
	c++ -arch x86_64 -std=c++11 -shared -o libmosh.so -I ../../mosh/ -I ../../mosh/src/crypto -I ../../mosh/src/frontend -I ../../mosh/src/network -I ../../mosh/src/protobufs -I ../../mosh/src/statesync -I ../../mosh/src/terminal -I ../../mosh/src/util -lncurses -lprotobuf mosh_wrap.cxx ../../mosh/src/crypto/base64.o ../../mosh/src/crypto/crypto.o ../../mosh/src/crypto/ocb.o ../../mosh/src/frontend/terminaloverlay.o ../../mosh/src/network/network.o ../../mosh/src/protobufs/hostinput.pb.o ../../mosh/src/statesync/completeterminal.o ../../mosh/src/terminal/parser.o ../../mosh/src/terminal/parseraction.o ../../mosh/src/terminal/parserstate.o ../../mosh/src/terminal/terminal.o ../../mosh/src/terminal/terminaldispatcher.o ../../mosh/src/terminal/terminaldisplay.o ../../mosh/src/terminal/terminaldisplayinit.o ../../mosh/src/terminal/terminalframebuffer.o ../../mosh/src/terminal/terminalfunctions.o ../../mosh/src/terminal/terminaluserinput.o ../../mosh/src/util/timestamp.o

clean:
	rm -f mosh.go mosh_wrap.cxx libmosh.so

all: mosh
