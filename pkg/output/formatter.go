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
	"encoding/json"
	"fmt"

	"github.com/tidwall/pretty"
	yaml "gopkg.in/yaml.v2"
)

// Formatter is used to format retrieved object to selected
// format and print it out on STDOUT and also colorized it
// if the flag is set.
type Formatter struct {
	builder *Builder
}

// Format prints the representation of input 'object' to
// STDOUT based on the requested 'format' (JSON or YAML).
// The output will also be colorized if flag is set.
func (f *Formatter) Format(object interface{}) error {
	builder := f.builder
	if builder.format == "json" {
		return f.toJSON(object, builder)
	} else if builder.format == "yaml" {
		return f.toYAML(object, builder)
	}
	return nil
}

// toJSON prints the JSON representation of input 'object'
// to STDOUT. The output will be colorized if flag is set.
func (f *Formatter) toJSON(object interface{}, builder *Builder) error {
	jsoned, err := json.Marshal(object)
	if err != nil {
		return err
	}
	jsoned = pretty.Pretty(jsoned)
	if builder.colored {
		jsoned = pretty.Color(jsoned, nil)
	}
	_, err = fmt.Printf("%s", jsoned)
	return err
}

// toYAML prints the YAML representation of input 'object'
// to STDOUT. The output will be colorized if flag is set.
func (f *Formatter) toYAML(object interface{}, builder *Builder) error {
	yamled, err := yaml.Marshal(object)
	if err != nil {
		return err
	}
	// TODO colorize yaml
	// if builder.colored {
	// yamled = Color(yamled, nil)
	// }
	_, err = fmt.Printf("%s", yamled)
	return err
}
