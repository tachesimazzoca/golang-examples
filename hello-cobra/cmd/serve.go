package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		var host string
		var port int
		var err error
		if host, err = cmd.Flags().GetString("host"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if port, err = cmd.Flags().GetInt("port"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("host: %s, port: %d", host, port)
		fmt.Println()
	},
}

func init() {
	serveCmd.Flags().StringP("host", "i", "localhost", "hostname")
	serveCmd.Flags().IntP("port", "p", 8080, "port number")
	rootCmd.AddCommand(serveCmd)
}
