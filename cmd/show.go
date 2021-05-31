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

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show commands displays entires.",
	Long: `show command displays entires.

Example:

now show (Shows the last entry)
now show --count n (Shows last n entries of the day)
now show -c n (Shows last n entries of the day)

`,
	Run: func(cmd *cobra.Command, args []string) {
		n, _ := cmd.Flags().GetInt("count")
		var records [][]string
		if n > 0 {
			records = funcs.ReadLastEntry(n)
		} else {
			records = funcs.ReadLastEntry(1)
		}
		for i := 0; i < 20; i++ {
			fmt.Print("-")
		}
		fmt.Println()
		for _, record := range records {
			fmt.Println(strings.Join(record, ">"))
		}
		for i := 0; i < 20; i++ {
			fmt.Print("-")
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().IntP("count", "c", -1, "entry display count")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
