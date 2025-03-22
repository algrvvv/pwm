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
	"fmt"

	"github.com/spf13/cobra"

	"github.com/algrvvv/pwm/gpg"
	"github.com/algrvvv/pwm/log"
	"github.com/algrvvv/pwm/storage"
)

// storeCmd represents the store command
var (
	withoutPassword bool
	storeCmd        = &cobra.Command{
		Use:     "store",
		Short:   "Add new note",
		Example: `Usage: pwm store name_note "some_string". some_string can be password or another note`,
		Run: func(_ *cobra.Command, args []string) {
			if len(args) != 2 {
				fmt.Println("invalid args; use pwm help store")
				return
			}

			name, value := args[0], args[1]

			encryptedValue := value
			if !withoutPassword {
				var err error
				encryptedValue, err = gpg.Encrypt(value)
				if err != nil {
					cobra.CheckErr(err)
				}
			}

			note := storage.Note{Name: name, Value: encryptedValue, UsePassword: !withoutPassword}
			if err := storageInstance.SaveNote(note); err != nil {
				cobra.CheckErr(err)
			}
			// un := color.New(color.Underline).Sprint(name)
			log.Infof("note with name %s saved\n", log.USF(name))
		},
	}
)

func init() {
	rootCmd.AddCommand(storeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	storeCmd.Flags().BoolVarP(&withoutPassword, "without-password", "W", false, "dont use password for encrypt")
}
