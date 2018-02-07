package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/clearbit/go-ddb"
)

func NewConnection()  *dynamodb.DynamoDB{
	c := credentials.NewEnvCredentials()

	config := aws.Config{Credentials:c,Region:aws.String("us-east-1"),LogLevel:aws.LogLevel(aws.LogDebugWithHTTPBody)}
	sess, err := session.NewSession(&config)
	if err != nil {
		panic(err)
	}
	return dynamodb.New(sess)
}

func NewScanner() *ddb.Scanner{
	config := ddb.Config{Svc:NewConnection()}
	return ddb.NewScanner(config)
}

