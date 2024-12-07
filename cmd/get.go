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

// getCmd represents the get command
var (
	clipFlag bool
	getCmd   = &cobra.Command{
		Use:               "get",
		Short:             "Command for get your pass by name",
		Long:              `Usage: pwm get note_name`,
		ValidArgsFunction: GetValidArgs,
		Run: func(_ *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("invalid params; use pwm help get")
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

			decryptedValue, err := gpg.Decrypt(note.Value)
			if err != nil {
				fmt.Printf("failed to decrypt data: %v\n", err)
				return
			}

			if clipFlag {
				utils.Copy(decryptedValue)
			}

			log.Infof("note with name %s founded\n", log.USF(name))
			fmt.Printf("%s: %s\n", note.Name, decryptedValue)
		},
	}
)

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVarP(&clipFlag, "clip", "c", false, "also save to clipboard")
}
