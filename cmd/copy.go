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
	"database/sql"
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/algrvvv/pwm/gpg"
	"github.com/algrvvv/pwm/log"
	"github.com/algrvvv/pwm/utils"
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:               "copy",
	Short:             "Get and save to clipboard note by name",
	Example:           "Usage: pwm copy note_name",
	ValidArgsFunction: GetValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("invalid params; use pwm help copy")
			return
		}
		name := args[0]
		note, err := storageInstance.GetNoteByName(name)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Infof("Note by name %s not found\n", log.USF(name))
				return
			}

			cobra.CheckErr(err)
		}

		decryptedValue := note.Value
		if note.UsePassword {
			decryptedValue, err = gpg.Decrypt(note.Value)
			if err != nil {
				fmt.Printf("failed to decrypt data: %v\n", err)
				return
			}
		}

		utils.Copy(decryptedValue)
		log.Infof("note with name %s founded and save to clipboard\n", log.USF(name))
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// copyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// copyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
