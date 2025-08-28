#!/usr/bin/env bash

set -o errexit

# change into the mosh directory (relative to this script)
pushd .
cd "$(dirname "$0")/mosh"

# build mosh
./autogen.sh
# Note, the configure script has 7 flags controlling static linking of libstdc++, libgcc, utempter, zlib, curses, crypto
# and (most importantly) protobuf (plus the following 8th meta-flag to turn on all 7). Despite the larger size, we want
# to statically link the mosh dependencies so that consumers of this go-mosh library do not need to deal with them.
./configure --enable-static-libraries
make
popd
