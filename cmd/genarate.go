/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"strings"
	"super-code-gen/pkg/generate"

	"github.com/spf13/cobra"
)

// gCmd represents the g command
var gCmd = &cobra.Command{
	Use:   "g",
	Short: "Generates route with controllers",
	Long: `
Generates route with controllers,
For e.g. super-code-gen g user/updateUser`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("no actions supplied")
		}
		for _, arg := range args {
			if contains := strings.Contains(arg, "/"); !contains {
				return errors.New("invalid actions format found")

			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		vStatus, _ := cmd.Flags().GetBool("validate")
		vBody, _ := cmd.Flags().GetBool("body")
		name := strings.Split(args[0], "/")[0]
		paths := make([]string, 0)
		for _, arg := range args {
			paths = append(paths, strings.Split(arg, "/")[1])
		}
		genCode(name, paths, vStatus, vBody)

	},
}

func genCode(name string, routes []string, validate bool, body bool) {
	generate.GenRoute(name, routes, validate)
	generate.GenController(name, routes, body)
}

func init() {
	rootCmd.AddCommand(gCmd)
	gCmd.Flags().BoolP("validate", "v", false, "Import body from express-validator for validation")
	gCmd.Flags().BoolP("body", "b", false, "Add body interface and create body const with type support in controller")
}
