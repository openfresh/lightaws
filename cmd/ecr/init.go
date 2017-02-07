package ecr

import (
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/openfresh/lightaws/remote"
	"github.com/spf13/cobra"
)

var (
	EcrCmd = &cobra.Command{
		Use:  "ecr",
		Long: `Amazon EC2 Container Registry (Amazon ECR) is a managed AWS Docker registry service. Customers can use the familiar Docker CLI to push, pull, and manage images. Amazon ECR provides a secure, scalable, and reliable registry. Amazon ECR supports private Docker repositories with resource-based permissions using AWS IAM so that specific users or Amazon EC2 instances can access repositories and images. Developers can use the Docker CLI to author and manage images.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			client = ecr.New(remote.NewAWSSession(region))
			return nil
		},
	}
	client *ecr.ECR
	region string
)

func init() {
	EcrCmd.AddCommand(getLoginCmd)
	EcrCmd.AddCommand(createReositoryCmd)
	EcrCmd.PersistentFlags().StringVarP(&region, "region", "r", "", "")
}
