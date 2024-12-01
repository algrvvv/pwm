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
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.design/x/clipboard"

	"github.com/algrvvv/pwm/database"
	"github.com/algrvvv/pwm/storage"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile         string
	storageInstance storage.Storage

	rootCmd = &cobra.Command{
		Use:   "pwm",
		Short: "password manager",
		Long:  `pwm - Password manager`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if cmd.Name() == "init" {
				return nil
			}

			if err := database.Open(); err != nil {
				return err
			}

			storageInstance = storage.NewSqliteStorage()
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			if database.C != nil {
				return database.Close()
			}
			return nil
		},
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initClipboard)

	rootCmd.PersistentFlags().
		StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pwm.yml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		path := filepath.Join(home, ".pwm")
		viper.AddConfigPath(path)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		cobra.CheckErr(err)
	}
}

func initClipboard() {
	if err := clipboard.Init(); err != nil {
		panic("failed to init clipboard: " + err.Error())
	}
}
