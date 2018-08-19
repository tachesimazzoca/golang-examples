package cmd

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Base64 encoder and decoder",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Usage())
	},
}

var base64EncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "converts string into base64 characters",
	Long:  "",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(base64.StdEncoding.EncodeToString([]byte(args[0])))
	},
}

var base64DecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "converts base64 characters into quoted string",
	Long:  "",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if data, err := base64.StdEncoding.DecodeString(args[0]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Printf("%q", data)
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.AddCommand(base64EncodeCmd)
	base64Cmd.AddCommand(base64DecodeCmd)
}
