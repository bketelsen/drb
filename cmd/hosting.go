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
	"fmt"

	"github.com/devrel-blox/drb/hosting"
	_ "github.com/devrel-blox/drb/hosting/azure"
	_ "github.com/devrel-blox/drb/hosting/netlify"
	_ "github.com/devrel-blox/drb/hosting/vercel"
	"github.com/spf13/cobra"
)

var (
	provider string
)

// hostingCmd represents the hosting command
var hostingCmd = &cobra.Command{
	Use:   "hosting",
	Short: "Generate the necessary boiler plate to host content",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("hosting called")
	},
}
var cmdList = &cobra.Command{
	Use:   "list",
	Short: "List available providers",
	Long:  `List available hosting providers.`,

	Run: func(cmd *cobra.Command, args []string) {
		list := hosting.Providers()
		for _, p := range list {
			fmt.Printf("%s:\t %s\n", p.Name(), p.Description())
		}
	},
}
var cmdInstall = &cobra.Command{
	Use:   "install",
	Short: "Install hosting support for a provider",
	Long:  `Install hosting support for a provider`,

	Run: func(cmd *cobra.Command, args []string) {
		p := hosting.GetProvider(provider)
		if p == nil {
			err := errors.New("unknown provider")
			cobra.CheckErr(err)
		}
		fmt.Println("Installing support for", p.Name())
		p.Install()
	},
}

func init() {
	hostingCmd.AddCommand(cmdList)
	hostingCmd.AddCommand(cmdInstall)
	rootCmd.AddCommand(hostingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cmdInstall.Flags().StringVarP(&provider, "provider", "p", "vercel", "hosting provider to target")
	cobra.CheckErr(cmdInstall.MarkFlagRequired("provider"))
}
