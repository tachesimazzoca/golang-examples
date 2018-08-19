package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "hello-cobra",
	Version: "0.0.0",
	Short:   "Shows how to implement a CLI tool using github.com/spf13/cobra",
	Long:    "Some long description here.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello cobra")
		fmt.Println(args)
	},
}

var config string

var verbose bool = false

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&config, "config", "c", "", "config file")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "d", false, "verbose mode")
}

func initConfig() {
	debug("cobra.OnInitialize")
	debug("config: " + config)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
