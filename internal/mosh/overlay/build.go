package overlay

// Per https://github.com/golang/go/issues/6696 -- use a build.go file to specify import "C"/cgo pragmas for SWIG
// generated wrappers. To build this project, there are fairly involved additional includes/linking requirements.

// #cgo CXXFLAGS: -I ${SRCDIR}/../../../mosh/ -I ${SRCDIR}/../../../mosh/src/crypto -I ${SRCDIR}/../../../mosh/src/frontend -I ${SRCDIR}/../../../mosh/src/network -I ${SRCDIR}/../../../mosh/src/protobufs -I ${SRCDIR}/../../../mosh/src/statesync -I ${SRCDIR}/../../../mosh/src/terminal -I ${SRCDIR}/../../../mosh/src/util
// #cgo LDFLAGS: ${SRCDIR}/../../../mosh/src/terminal/parser.o ${SRCDIR}/../../../mosh/src/terminal/parseraction.o ${SRCDIR}/../../../mosh/src/terminal/parserstate.o ${SRCDIR}/../../../mosh/src/frontend/terminaloverlay.o ${SRCDIR}/../../../mosh/src/terminal/terminaldispatcher.o ${SRCDIR}/../../../mosh/src/terminal/terminalframebuffer.o ${SRCDIR}/../../../mosh/src/terminal/terminal.o ${SRCDIR}/../../../mosh/src/terminal/terminaluserinput.o ${SRCDIR}/../../../mosh/src/terminal/terminalfunctions.o ${SRCDIR}/../../../mosh/src/network/network.o ${SRCDIR}/../../../mosh/src/util/timestamp.o ${SRCDIR}/../../../mosh/src/crypto/crypto.o ${SRCDIR}/../../../mosh/src/crypto/base64.o ${SRCDIR}/../../../mosh/src/crypto/ocb.o
import "C"
