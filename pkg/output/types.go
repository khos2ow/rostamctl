// Copyright © 2019 Khosrow Moossavi.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package output

import (
	"strings"
)

var outputFormats = []string{"json", "yaml"}

// Get returns available output formats
func Get() []string {
	return outputFormats
}

// Has returns true if the list of available output formats
// contains the provided string, false if not
func Has(format string) bool {
	for _, f := range outputFormats {
		if format == f {
			return true
		}
	}
	return false
}

// FormatStrings returns a string representing all output formats.
// this is useful for help text / flag info
func FormatStrings() string {
	var b strings.Builder
	b.WriteString("[")
	for i, format := range outputFormats {
		b.WriteString(format)
		if i+1 != len(outputFormats) {
			b.WriteString(", ")
		}
	}
	b.WriteString("]")
	return b.String()
}
