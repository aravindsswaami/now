/*
Copyright © 2021 Ashwin Raam <raamocyby@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"now/funcs"
	"strings"

	"github.com/spf13/cobra"
)

func doneEntry(tags []string, args []string) {
	entry := funcs.Done(strings.Trim(fmt.Sprint(args), "[]"), tags)
	funcs.AppendEntry(entry)
}

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "done command adds task completed entry.",
	Long: `done command adds task completed entry.

Example:

now done taskTitle
now done taskTitle --tags tag1,tag2
now done taskTitle -t tag1,tag2

`,
	Run: func(cmd *cobra.Command, args []string) {
		tags, _ := cmd.Flags().GetStringSlice("tags")
		if len(args) < 1 {
			record := funcs.ReadLastEntry(1)
			if record[0][3] == " started" {
				task := strings.Split(strings.Trim(record[0][1], " "), " ")
				newTags := strings.Split(strings.Trim(record[0][2], " "), " ")
				doneEntry(newTags, task)
			} else {
				fmt.Println("No open task is available to complete!")
				fmt.Println("Please create new task with the help of now doing command.")
			}
		} else {
			doneEntry(tags, args)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
	doneCmd.Flags().StringSliceP("tags", "t", []string{"General"}, "adds tags")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
