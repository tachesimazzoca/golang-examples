package cmd

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

var s3Cmd = &cobra.Command{
	Use: "s3",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var s3BucketsCmd = &cobra.Command{
	Use:   "buckets",
	Short: "Display a list of all s3 buckets",
	Run: func(cmd *cobra.Command, args []string) {
		sess := newSession()
		svc := s3.New(sess)
		out, err := svc.ListBuckets(nil)
		if err != nil {
			panic(err)
		}
		fmt.Println(tsvFormat("Name", "CreationDate"))
		for _, x := range out.Buckets {
			fmt.Println(tsvFormat(*x.Name, x.CreationDate.Format(time.RFC3339)))
		}
	},
}

func init() {
	rootCmd.AddCommand(s3Cmd)
	s3Cmd.AddCommand(s3BucketsCmd)
}
