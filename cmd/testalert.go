// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"log"
	"os"

	"github.com/spf13/cobra"
)

// testalertCmd represents the testalert command
var testalertCmd = &cobra.Command{
	Use:   "testalert",
	Short: "send a test alert to test auth credentials",
	Long:  `Use the authentication credentials in the config file to send a test alert`,
	Run: func(cmd *cobra.Command, args []string) {
		Config.Containers = []Container{
			Container{
				Name: "fd42c70222be1d96224ffeb28416d4b61ffa431c0aa97818cf5ef67e9317a7d8",
			},
		}

		err := Config.Validate()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		Start(&Config)
	},
}

func init() {
	RootCmd.AddCommand(testalertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testalertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testalertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
