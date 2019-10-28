/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"agenda/service"
	"github.com/spf13/cobra"
)


// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register -n [username] -p [password] -e [email] -t [phone]",
	Short: "user register",
	Long: `Command like:register -n heihei -p 000000 -e 161248@qq.com -t 16513`,
	Run: func(cmd *cobra.Command, args []string) {
		tmp_n, _ := cmd.Flags().GetString("name")
		tmp_p, _ := cmd.Flags().GetString("password")
		tmp_e, _ := cmd.Flags().GetString("email")
		tmp_t, _ := cmd.Flags().GetString("phone")
		if service.GetFlag() == true {
			service.RegisterUser(tmp_n, tmp_p,tmp_e,tmp_t)
		} else {
			fmt.Println("Please log out first")
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	registerCmd.Flags().StringP("name", "n", "", "user name")
	registerCmd.Flags().StringP("password", "p", "", "user password")
	registerCmd.Flags().StringP("email", "e", "", "user email")
	registerCmd.Flags().StringP("phone", "t", "", "user phone")
}
