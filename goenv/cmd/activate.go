// Copyright © 2018 Moises P. Sena <moisespsena@gmail.com>
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

package cmd

import (
	"github.com/moisespsena-go/goenv"
	"github.com/spf13/cobra"
)

var activateCmd = &cobra.Command{
	Use:   "activate NAME",
	Short: "Activate the virtualenv with NAME.",
	Long: `Activate the virtualenv with NAME.
Examples:
  $ eval $(goenv activate teste)
  $ eval $(goenv -d ~/my-goenv activate teste)
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		env, err := goenv.NewGoEnvCmd(db, true)
		if err != nil {
			return err
		}
		return env.ActivateCode(args[0])
	},
}

func init() {
	rootCmd.AddCommand(activateCmd)
}
