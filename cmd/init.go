/*
Copyright © 2024 algrvvv <alexandrgr25@gmail.com>

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
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	GPGUser          string `yaml:"gpg_user"`
	ClipboardTimeout int    `yaml:"clipboard_timeout"`
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init configiguration",
	Long:  `Init configuration and database`,
	Run: func(cmd *cobra.Command, args []string) {
		config := Config{}
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("Enter gpg default user: ")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("failed to read input: ", err)
				continue
			}

			config.GPGUser = input[:len(input)-1]

			fmt.Print("Enter clipboard timeout in minutes: ")
			input, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println("failed to read input: ", err)
				continue
			}

			timeout, err := strconv.Atoi(input[:len(input)-1])
			if err != nil {
				fmt.Println("failed to parse your timeout: ", err)
				continue
			}

			config.ClipboardTimeout = timeout
			break
		}

		fmt.Println("config: ", config)

		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		configDir := filepath.Join(home, ".pwm")
		err = os.Mkdir(configDir, 0o700)
		if err != nil && !errors.Is(err, os.ErrExist) {
			fmt.Println("failed to create config dir: ", err)
			return
		}

		configPath := filepath.Join(configDir, "config.yml")
		fmt.Println("got path: " + configPath)

		viper.SetConfigFile(configPath)
		viper.Set("gpg_user", config.GPGUser)
		viper.Set("clipboard_timeout", config.ClipboardTimeout)
		_ = viper.WriteConfig()

		dbpath := filepath.Join(configDir, "database.db")
		f, err := os.OpenFile(dbpath, os.O_RDWR|os.O_CREATE, 0o600)
		if err != nil {
			fmt.Println("failed to create database: ", err)
			return
		}
		cobra.CheckErr(f.Close())
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
