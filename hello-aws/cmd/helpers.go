package cmd

import (
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	gohomedir "github.com/mitchellh/go-homedir"
)

var homedir string

func init() {
	var err error
	homedir, err = gohomedir.Dir()
	if err != nil {
		panic(err)
	}
}

func newSession() *session.Session {
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String(region),
	}))
	return sess
}

func fromHomedir(elem ...string) string {
	return filepath.Join(homedir, filepath.Join(elem...))
}

func tsvFormat(elem ...string) string {
	return strings.Join(elem, "\t")
}
