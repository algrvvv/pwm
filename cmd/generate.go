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
	"github.com/algrvvv/pwm/utils"
)

// generateCmd represents the generate command
var (
	passLen         int
	withoutUppers   bool
	withoutDigits   bool
	withoutSpecials bool
	clip            bool
	save            string

	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate new password",
		Example: `Usage: pwm generate --len 22 --clip --save generate_password
    In this example generated password has 22 symbols and save to database and clipboard`,
		Run: func(cmd *cobra.Command, args []string) {
			pass, err := utils.GeneratePassword(
				passLen,
				!withoutUppers,
				!withoutDigits,
				!withoutSpecials,
			)
			if err != nil {
				fmt.Println("failed to generate password: ", err)
				return
			}

			log.Infof("Generated password: %s\n", log.USF(pass))

			if save != "" {
				decValue, err := gpg.Encrypt(pass)
				if err != nil {
					fmt.Println("failed to encrypt data: ", err)
					return
				}

				note := storage.Note{Name: save, Value: decValue}
				if err = storageInstance.SaveNote(note); err != nil {
					fmt.Println("failed to save note: ", err)
					return
				}
				log.Infof("Generated password has been saved with name: %s\n", log.USF(save))
			}

			if clip {
				utils.Copy(pass)
				log.Infof("Generated password has been saved to clipboard\n")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntVarP(&passLen, "len", "l", 12, "password len. default: 12")
	generateCmd.Flags().
		BoolVarP(&withoutUppers, "without-uppers", "U", false, "dont use upper case symbols")
	generateCmd.Flags().
		BoolVarP(&withoutDigits, "without-digits", "D", false, "dont use digits")
	generateCmd.Flags().
		BoolVarP(&withoutSpecials, "without-specials", "S", false, "dont use special symbols")
	generateCmd.Flags().
		BoolVarP(&clip, "clip", "c", false, "save to clipboard")
	generateCmd.Flags().
		StringVarP(&save, "save", "s", "", "save note by name")
}
