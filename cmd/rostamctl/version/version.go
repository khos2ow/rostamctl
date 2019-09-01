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

// Package version implements the `version` command
package version

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/khos2ow/rostamctl/pkg/cli"
	"github.com/spf13/cobra"
)

// current version
const current = "v0.0.1-dev"

// Provisioned by ldflags
var (
	version    string
	commitHash string
	buildDate  string
)

// NewCommand returns a new cobra.Command for version
func NewCommand(cli *cli.Wrapper) *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "version",
		Short: "Print the rostamctl CLI version",
		Long:  "Print the rostamctl CLI version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(fmt.Sprintf("rostamctl version %s", Version()))
			return nil
		},
	}

	return cmd
}

// Version return the full version of the binary including commit hash and build date
func Version() string {
	if version == "" {
		version = current
	}
	if commitHash != "" && !strings.HasSuffix(version, commitHash) {
		version += " " + commitHash
	}
	if buildDate == "" {
		buildDate = time.Now().Format(time.RFC3339)
	}

	return fmt.Sprintf("%s %s/%s BuildDate: %s", version, runtime.GOOS, runtime.GOARCH, buildDate)
}
