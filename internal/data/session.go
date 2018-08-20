package data

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewSession() (*session.Session, error) {
	session, err := session.NewSession(&aws.Config{
    	Region: aws.String("us-west-2")},
	)
	return session, err
}
