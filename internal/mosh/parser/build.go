package parser

// Per https://github.com/golang/go/issues/6696 -- use a build.go file to specify import "C"/cgo pragmas for SWIG
// generated wrappers. To build this project, there are fairly involved additional includes/linking requirements.

// #cgo CXXFLAGS: -std=c++11 -I ${SRCDIR}/../../../mosh/ -I ${SRCDIR}/../../../mosh/src/crypto -I ${SRCDIR}/../../../mosh/src/frontend -I ${SRCDIR}/../../../mosh/src/network -I ${SRCDIR}/../../../mosh/src/protobufs -I ${SRCDIR}/../../../mosh/src/statesync -I ${SRCDIR}/../../../mosh/src/terminal -I ${SRCDIR}/../../../mosh/src/util
// #cgo LDFLAGS: ${SRCDIR}/../../../mosh/src/terminal/parser.o ${SRCDIR}/../../../mosh/src/terminal/parseraction.o ${SRCDIR}/../../../mosh/src/terminal/parserstate.o ${SRCDIR}/../../../mosh/src/terminal/terminaldispatcher.o ${SRCDIR}/../../../mosh/src/terminal/terminal.o ${SRCDIR}/../../../mosh/src/terminal/terminalframebuffer.o ${SRCDIR}/../../../mosh/src/terminal/terminaluserinput.o ${SRCDIR}/../../../mosh/src/terminal/terminalfunctions.o
import "C"
