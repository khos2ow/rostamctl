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

// Package list implements the `twitter list` command
package list

import (
	"github.com/khos2ow/rostamctl/pkg/api/twitter"
	"github.com/khos2ow/rostamctl/pkg/cli"
	"github.com/khos2ow/rostamctl/pkg/output"
	"github.com/khos2ow/rostamctl/pkg/util"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for twitter list
func NewCommand(cli *cli.Wrapper) *cobra.Command {
	cmd := &cobra.Command{
		Args:    cobra.NoArgs,
		Aliases: []string{"ls"},
		Use:     "list",
		Short:   "List blocked accounts",
		Long:    util.LongDescription(`List blocked Twitter accounts.`),
		RunE: func(cmd *cobra.Command, args []string) error {
			list, err := twitter.NewClient().List()
			if err != nil {
				return err
			}
			return cli.OutputBuilder.Build(func(formatter *output.Formatter) error {
				return formatter.Format(list)
			})
		},
	}

	return cmd
}
