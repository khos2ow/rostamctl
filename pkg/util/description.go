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

// Package util contains general utility of the rostamctl
package util

import (
	"fmt"
	"strings"

	"github.com/lithammer/dedent"
)

// LongDescription formats long multi-line description and removes
// the left empty space from the lines
func LongDescription(a interface{}) string {
	return strings.TrimLeft(dedent.Dedent(fmt.Sprint(a)), "\n")
}
