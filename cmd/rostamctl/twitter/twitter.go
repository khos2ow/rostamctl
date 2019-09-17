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

// Package twitter implements the `twitter` command
package twitter

import (
	"github.com/khos2ow/rostamctl/cmd/rostamctl/twitter/list"
	"github.com/khos2ow/rostamctl/pkg/cli"
	"github.com/khos2ow/rostamctl/pkg/util"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for twitter actions
func NewCommand(cli *cli.Wrapper) *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "twitter",
		Short: "Manage activity of Twitter accounts",
		Long: util.LongDescription(`
            Manage activity of Twitter accounts
        `),
	}

	cmd.AddCommand(list.NewCommand(cli))

	return cmd
}
