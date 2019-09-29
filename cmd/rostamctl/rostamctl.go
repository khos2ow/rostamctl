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

// Package rostamctl implements root command
package rostamctl

import (
	"os"

	"github.com/khos2ow/rostamctl/cmd/rostamctl/completion"
	"github.com/khos2ow/rostamctl/cmd/rostamctl/twitter"
	"github.com/khos2ow/rostamctl/cmd/rostamctl/version"
	"github.com/khos2ow/rostamctl/pkg/cli"
	"github.com/khos2ow/rostamctl/pkg/flags"
	"github.com/khos2ow/rostamctl/pkg/output"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	logutil "sigs.k8s.io/kind/pkg/log"
)

// NewCommand returns a new cobra.Command implementing the root command for rostamctl
func NewCommand() *cobra.Command {
	cli := &cli.Wrapper{}
	flg := &flags.GlobalFlags{}
	cmd := &cobra.Command{
		Args:         cobra.NoArgs,
		Use:          "rostamctl",
		Short:        "rostamctl manages authentication, configurations and interactions with the RostamBot APIs.",
		Long:         "rostamctl manages authentication, configurations and interactions with the RostamBot APIs.",
		SilenceUsage: true,
		Version:      version.Version(),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := flg.Normalize(cmd, args); err != nil {
				return err
			}
			cli.GlobalFlags = flg
			cli.OutputBuilder = output.NewBuilder(flg.OutputFormat)
			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&flg.OutputFormat, "output", flags.DefaultOutputFormat, "output format "+output.FormatStrings())
	cmd.PersistentFlags().StringVar(&flg.LogLevel, "loglevel", flags.DefaultLogLevel.String(), "log level "+logutil.LevelsString())

	cmd.AddCommand(completion.NewCommand(cli))
	cmd.AddCommand(twitter.NewCommand(cli))
	cmd.AddCommand(version.NewCommand(cli))

	return cmd
}

// Run runs the `rostamctl` root command
func Run() error {
	return NewCommand().Execute()
}

// Main wraps Run and sets the log formatter
func Main() {
	// let's explicitly set stdout
	logrus.SetOutput(os.Stdout)

	// this formatter is the default, but the timestamps output aren't
	// particularly useful, they're relative to the command start
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
		// we force colors because this only forces over the isTerminal check
		// and this will not be accurately checkable later on when we wrap
		// the logger output with our logutil.StatusFriendlyWriter
		ForceColors: logutil.IsTerminal(logrus.StandardLogger().Out),
	})

	if err := Run(); err != nil {
		os.Exit(1)
	}
}
