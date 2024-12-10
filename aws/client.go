package aws

import (
	"os"
)

type Credentials struct {
	AccessKey string
	SecretKey string
}

func GetAWSCredentials() Credentials {
	return Credentials{
		AccessKey: os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}
}
