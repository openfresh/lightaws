package ec2

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/openfresh/lightaws/remote"
	"github.com/spf13/cobra"
)

var (
	EC2Cmd = &cobra.Command{
		Use:  "ec2",
		Long: `EC2`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			client = ec2.New(remote.NewAWSSession(region))
			return nil
		},
	}
	client     *ec2.EC2
	instanceID string
	region     string
	tagName    string
)

func init() {
	EC2Cmd.AddCommand(describeNameTagCmd)
	EC2Cmd.PersistentFlags().StringVarP(&instanceID, "instance-id", "i", "", "")
	EC2Cmd.PersistentFlags().StringVarP(&region, "region", "r", "", "")
}
