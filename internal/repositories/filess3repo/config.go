package filess3repo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

func ConnectAws() *session.Session {
	MyRegion := os.Getenv("AWS_REGION")
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
		},
	)
	if err != nil {
		panic(err)
	}
	return sess
}
