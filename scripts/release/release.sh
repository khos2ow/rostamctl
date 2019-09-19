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
set -o pipefail

NEW_VERSION=$1
PUSH=$2
CURRENT_VERSION=$3
MAKEFILE=$4

if [ -z "${NEW_VERSION}" ]; then
    if [ -z "${MAKEFILE}" ]; then
        echo "Error: VERSION is missing. e.g. ./release.sh <version> <push>"
    else
        echo "Error: missing value for 'version'. e.g. 'make release version=x.y.z'"
    fi
    exit 1
fi

if [ -z "${PUSH}" ]; then
    echo "Error: PUSH is missing. e.g. ./release.sh <version> <push>"
    exit 1
fi

if [ -z "${CURRENT_VERSION}" ]; then
    CURRENT_VERSION=$(git describe --tags --exact-match 2>/dev/null || git describe --tags 2>/dev/null || echo "v0.0.1-$(COMMIT_HASH)")
fi

if [ "v${NEW_VERSION}" = "${CURRENT_VERSION}" ]; then
    echo "Error: provided version (v${version}) exists."
    exit 1
fi

PWD=$(cd $(dirname "$0") && pwd -P)
CLOSEST_VERSION=$(git describe --tags --abbrev=0)

# Bump the released version in README and version.go
sed -i -E 's|'${CLOSEST_VERSION}'|v'${NEW_VERSION}'|g' README.md
sed -i -E 's|'${CLOSEST_VERSION}'-alpha|v'${NEW_VERSION}'|g' cmd/rostamctl/version/version.go

# Commit changes
git add README.md cmd/rostamctl/version/version.go
git commit -m "Release version v${NEW_VERSION}"

if [ "${PUSH}" == "true" ]; then
    git push origin master
fi

# Tag the release
git tag --annotate --message "v${NEW_VERSION} Release" "v${NEW_VERSION}"

if [ "${PUSH}" == "true" ]; then
    git push origin v${NEW_VERSION}
fi

# Generate Changelog
make --no-print-directory -f ${PWD}/../../Makefile changelog push="${PUSH}"

# Bump the next version in version.go
NEXT_VERSION=$(echo "${NEW_VERSION}" | sed 's/^v//' | awk -F'[ .]' '{print $1"."$2+1".0"}')
sed -i -E 's|'${NEW_VERSION}'|'${NEXT_VERSION}'-alpha|g' cmd/rostamctl/version/version.go

# Commit changes
git add cmd/rostamctl/version/version.go
git commit -m "Bump version to ${NEXT_VERSION}-alpha"

if [ "${PUSH}" == "true" ]; then
    git push origin master
fi
