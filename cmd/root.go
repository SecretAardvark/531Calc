/*
Copyright © 2020 Chad Tennent SecretAardvark@protonmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/SecretAardvark/531calc/lift"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	//Weight  is the amount of weight you lifted.
	Weight float32
	//Reps is the number of reps you did.
	Reps    float32
	cfgFile string
	lifts   []lift.Lift
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "531calc",
	Short: "531calc takes your lifting numbers and calculates a full 5/3/1 cycle.",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		Lift := lift.Lift{Name: os.Args[1]}
		Lift.GetOneRep(Weight, Reps)
		Lift.GetTM()
		Lift.GetCycle()
		fmt.Println(Lift)

		jsonfile, _ := os.Open("531.json")
		defer jsonfile.Close()

		bytes, err := ioutil.ReadAll(jsonfile)
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal(bytes, &lifts)

		lifts = append(lifts, Lift)

		file, _ := json.MarshalIndent(lifts, "", " ")
		_ = ioutil.WriteFile("531.json", file, 0644)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().Float32VarP(&Weight, "weight", "w", 0, "The amount of weight on the bar (required).")
	rootCmd.PersistentFlags().Float32VarP(&Reps, "reps", "r", 0, "The amount of reps you did (required).")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.531calc.yaml)")

	rootCmd.MarkPersistentFlagRequired("weight")
	rootCmd.MarkPersistentFlagRequired("reps")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".531calc" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".531calc")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func findOneRM(weight, reps float32) float32 {
	return weight*reps*float32(.0333) + weight
}
