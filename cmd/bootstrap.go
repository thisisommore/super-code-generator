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
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/spf13/cobra"
)

// bootstrapCmd represents the bootstrap command
var bootstrapCmd = &cobra.Command{
	Aliases: []string{"b"},
	Use:     "bootstrap",
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		remoteUrl, _ := cmd.Flags().GetString("remote")

		GenerateProject("", remoteUrl)
	},
}

func GenerateProject(path string, remoteUrl string) {
	repo, _ := git.PlainClone(path, false, &git.CloneOptions{
		URL: "https://github.com/thisisommore/nodejs-mongodb-template.git",
	})
	if len(remoteUrl) > 0 {
		repo.DeleteRemote("origin")
		repo.CreateRemote(&config.RemoteConfig{
			Name: "origin",
			URLs: []string{remoteUrl},
		})
	}

}
func init() {
	rootCmd.AddCommand(bootstrapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bootstrapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	bootstrapCmd.Flags().StringP("remote", "r", "", "Remote to use for git")
}
