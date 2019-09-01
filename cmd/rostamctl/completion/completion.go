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

// Package completion implements the `completion` command
package completion

import (
	"github.com/khos2ow/rostamctl/cmd/rostamctl/completion/bash"
	"github.com/khos2ow/rostamctl/cmd/rostamctl/completion/zsh"
	"github.com/khos2ow/rostamctl/pkg/cli"
	"github.com/khos2ow/rostamctl/pkg/util"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for shell completion
func NewCommand(cli *cli.Wrapper) *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "completion",
		Short: "Output completion code for the specified shell (bash or zsh)",
		Long: util.LongDescription(`
            Outputs rostamctl shell completion for the given shell (bash or zsh)
            This depends on the bash-completion binary.  Example installation instructions:

            # for bash users
                $ rostamctl completion bash > ~/.rostamctl-completion
                $ source ~/.rostamctl-completion

            # for zsh users
                % rostamctl completion zsh > /usr/local/share/zsh/site-functions/_rostamctl
                % autoload -U compinit && compinit

            Additionally, you may want to output the completion to a file and source in your .bashrc
            Note for zsh users: [1] zsh completions are only supported in versions of zsh >= 5.2
        `),
	}

	cmd.AddCommand(zsh.NewCommand(cli))
	cmd.AddCommand(bash.NewCommand(cli))

	return cmd
}
