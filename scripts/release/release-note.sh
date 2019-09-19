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

PWD=$(cd $(dirname "$0") && pwd -P)
CLOSEST_VERSION=$(git describe --tags --abbrev=0)

# Install git-chglog binary
export PATH=$PATH:$(go env GOPATH)/bin
make git-chglog

# Generate Changelog
git-chglog --config ${PWD}/../../scripts/chglog/config-release-note.yml --output ${PWD}/../../CURRENT-RELEASE-CHANGELOG.md ${CLOSEST_VERSION}

cat ${PWD}/../../CURRENT-RELEASE-CHANGELOG.md