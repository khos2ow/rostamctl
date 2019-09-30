#!/usr/bin/env bash
#
# Copyright © 2019 Khosrow Moossavi.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

BUILD_DIR="${1:-bin}"
GOOS="${2:-"linux darwin windows freebsd"}"
GOARCH="${3:-"amd64 arm"}"
GOLDFLAGS="$4"

if [ -z "${GOLDFLAGS}" ]; then
    echo "Error: GOLDFLAGS is missing. e.g. ./build-all.sh <build_dir> <build_os_list> <build_arch_list> <build_ldflag>"
    exit 1
fi

PWD=$(cd $(dirname "$0") && pwd -P)
BUILD_DIR="${PWD}/../../${BUILD_DIR}"

CGO_ENABLED=0 gox \
    -verbose \
    -ldflags "${GOLDFLAGS}" \
    -gcflags=-trimpath=$(go env GOPATH) \
    -os="${GOOS}" \
    -arch="${GOARCH}" \
    -osarch="!darwin/arm" \
    -output="${BUILD_DIR}/{{.OS}}-{{.Arch}}/{{.Dir}}" ${PWD}/../../
