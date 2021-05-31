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

func doingEntry(cmd *cobra.Command, args []string) {
	tags, _ := cmd.Flags().GetStringSlice("tags")
	entry := funcs.Doing(strings.Trim(fmt.Sprint(args), "[]"), tags)
	funcs.AppendEntry(entry)
}

// doingCmd represents the doing command
var doingCmd = &cobra.Command{
	Use:   "doing",
	Short: "doing command adds task started entry.",
	Long: `doing command adds task started entry.

Example:

now doing taskTitle
now doing taskTitle --tags tag1,tag2
now doing taskTitle -t tag1,tag2

`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			doingEntry(cmd, args)
		} else {
			fmt.Println("Enter task name after doing command.")
		}

	},
}

func init() {
	rootCmd.AddCommand(doingCmd)
	doingCmd.Flags().StringSliceP("tags", "t", []string{"General"}, "adds tags")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
