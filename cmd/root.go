/*
Copyright © 2020 Jay Baker <jay@jay-baker.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "ctxman",
	Short: "ctxman helps you manage contexts and environments",
	Run: func(cmd *cobra.Command, args []string) {

		switch argsCount := len(args); argsCount {
		case 0:
			fmt.Println("please specify an environment and optionally a namespace")
			os.Exit(1)
		case 1:
			fmt.Println("switching to env: ", args[0])
			err := Swap(args[0], "")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case 2:
			fmt.Println("switching to env: ", args[0])
			fmt.Println("and namespace: ", args[1])
		case 3:
			fmt.Println("too many arguments!")
			os.Exit(1)

		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("$HOME/")
	viper.AddConfigPath("$HOME/.config/ctxman/")
	viper.AddConfigPath("/etc/ctxman/")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("unable to read config file")
	}
}
