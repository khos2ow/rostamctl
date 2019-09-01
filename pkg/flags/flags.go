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

// Package flags contains general utility of the rostamctl flags
package flags

import (
	"github.com/khos2ow/rostamctl/pkg/output"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// GlobalFlags for the rostamctl command
type GlobalFlags struct {
	LogLevel      string
	OutputColored bool
	OutputFormat  string
}

// Normalize checks and normalizes input flags and falls back to default values when needed
func (gf *GlobalFlags) Normalize(cmd *cobra.Command, args []string) error {
	if err := gf.parseLogLevel(cmd, args); err != nil {
		return err
	}
	if err := gf.parseOutputFormat(cmd, args); err != nil {
		return err
	}
	return nil
}

func (gf *GlobalFlags) parseLogLevel(cmd *cobra.Command, args []string) error {
	level := DefaultLogLevel
	parsed, err := logrus.ParseLevel(gf.LogLevel)
	if err != nil {
		logrus.Warnf("Invalid log level '%s', defaulting to '%s'", gf.LogLevel, level)
	} else {
		level = parsed
	}
	logrus.SetLevel(level)
	return nil
}

func (gf *GlobalFlags) parseOutputFormat(cmd *cobra.Command, args []string) error {
	if !output.Has(gf.OutputFormat) {
		logrus.Warnf("Invalid output format '%s', defaulting to '%s'", gf.OutputFormat, DefaultOutputFormat)
		gf.OutputFormat = DefaultOutputFormat
	}
	return nil
}
