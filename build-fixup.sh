#!/usr/bin/env bash

# cd into the directory of this script
pushd . || exit 1
cd "$(dirname $0)" || exit 1

# fix-ups to mosh for building as a library
echo "Fixing up shared.h"
sed -i_ 's/#include "config.h"/#include "..\/include\/config.h"/' mosh/src/util/shared.h

# macOS (on case insensitive filesystems) chokes when building via go tools if this file is present --
# see https://github.com/mobile-shell/mosh/issues/1140 -- this was fixed within mosh itself for its own build, but not
# sure how to get it to apply to the build invoked for this library...
echo "Removing VERSION"
rm -f mosh/VERSION

popd
