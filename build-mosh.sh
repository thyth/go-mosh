#!/usr/bin/env bash

# change into the mosh directory (relative to this script)
pushd . || exit 1
cd "$(dirname $0)/mosh" || exit 1

./autogen.sh
./configure
make
popd
