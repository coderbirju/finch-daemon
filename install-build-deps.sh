#!/usr/bin/env bash

#   Copyright The Soci Snapshotter Authors.

#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at

#       http://www.apache.org/licenses/LICENSE-2.0

#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

# A script to install the build dependencies needed to build SOCI in a
# Ubuntu 20.04 container.
#
# Usage: bash install-build-dependencies.sh

apt -y update && apt install -y make \
    gcc \
    linux-libc-dev \
    libseccomp-dev \
    pkg-config \
    git