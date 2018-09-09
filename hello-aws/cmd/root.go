package cmd

import (
	"fmt"

	awsCreds "github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/spf13/cobra"
)

var region string

var credentials string

var profile string

var creds *awsCreds.Credentials

var rootCmd = &cobra.Command{
	Use:     "hello-aws",
	Version: "0.0.0",
	Short:   "Shows how to use AWS Go SDK",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("region: ", region)
		fmt.Println("credentials: ", credentials)
		fmt.Println("profile: ", profile)
	},
}

func init() {

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&region,
		"region", "r", "us-east-1", "one of the AWS regions")
	rootCmd.PersistentFlags().StringVarP(&credentials,
		"credentials", "c", fromHomedir(".aws", "credentials"),
		"path to shared .aws/credentials")
	rootCmd.PersistentFlags().StringVarP(&profile,
		"profile", "p", "default", "one of the profiles in shared credentials")
}

func initConfig() {
	creds = awsCreds.NewSharedCredentials(credentials, profile)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
