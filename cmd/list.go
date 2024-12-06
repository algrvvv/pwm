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
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"

	"github.com/algrvvv/pwm/log"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Show all your notes",
	Example: "pwm list",
	Run: func(_ *cobra.Command, args []string) {
		notes, err := storageInstance.GetNotes()
		if err != nil {
			cobra.CheckErr(err)
		}

		var search string
		if len(args) == 1 {
			search = args[0]
		}

		// реализация через использование команды less
		var result strings.Builder
		result.WriteString(log.Sinfo("List of your notes:\n\n"))
		for i, note := range notes {
			if search == "" || (search != "" && strings.Contains(note.Name, search)) {
				result.WriteString(fmt.Sprintf("\t%s. %s\n", log.USF(i+1), note.Name))
			}
		}

		cmd := exec.Command("less")
		cmd.Stdin = bytes.NewBufferString(result.String())
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			cobra.CheckErr(err)
		}

		// реализация через собственный простой аналог less
		// const pageSize = 2
		// scanner := bufio.NewScanner(os.Stdin)
		//
		// for i := 0; i < len(notes); i += pageSize {
		// 	e := i + pageSize
		// 	if e > len(notes) {
		// 		e = len(notes)
		// 	}
		//
		// 	for _, note := range notes[i:e] {
		// 		fmt.Println(note.Name)
		// 	}
		//
		// 	fmt.Print("\r")
		// 	scanner.Scan()
		// 	fmt.Print("\r")
		// 	if scanner.Text() == "q" {
		// 		return
		// 	}
		// }
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
