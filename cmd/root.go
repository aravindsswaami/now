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
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "now",
	Short: "Now is time marking cli-tool.",
	Long: `Now is a simple cli-tool for time marking that enables ease of tracking how the day went. 
Intuitive commands that allows you mark the start time and end time of a task.

Example:

now doing taskTitle --tags tag1,tag2
now done taskTitle --tags tag1,tag2

It also supports tags to group different task.
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()

		if err != nil {
			log.Panic(err)
		}

		_, err = os.Stat(home + "/nowData")

		if os.IsNotExist(err) {
			errDir := os.MkdirAll(home+"/nowData", 0755)
			if errDir != nil {
				log.Fatal(err)
			}

		}
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.now.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// if cfgFile != "" {
	// 	// Use config file from the flag.
	// 	viper.SetConfigFile(cfgFile)
	// } else {
	// 	// Find home directory.
	// 	home, err := homedir.Dir()
	// 	cobra.CheckErr(err)

	// 	// Search config in home directory with name ".now" (without extension).
	// 	viper.AddConfigPath(home)
	// 	viper.SetConfigName(".now")
	// }

	// viper.AutomaticEnv() // read in environment variables that match

	// // If a config file is found, read it in.
	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	// }
}
