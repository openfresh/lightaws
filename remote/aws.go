package remote

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewAWSSession(region string) *session.Session {
	ses := session.New()
	cred := credentials.NewChainCredentials([]credentials.Provider{
		&credentials.EnvProvider{},
		&credentials.SharedCredentialsProvider{},
		&ec2rolecreds.EC2RoleProvider{
			Client: ec2metadata.New(ses),
		},
	})

	ses.Config.Credentials = cred
	ses.Config.WithMaxRetries(aws.UseServiceDefaultRetries).WithRegion(region)
	return ses
}
