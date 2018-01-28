package metadata

import (
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/openfresh/lightaws/remote"
	"github.com/spf13/cobra"
)

var (
	MetadataCmd = &cobra.Command{
		Use:  "metadata",
		Long: `EC2 metadata`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			client = ec2metadata.New(remote.NewAWSSession(region))
			return nil
		},
	}
	client *ec2metadata.EC2Metadata
	region string
)

func init() {
	MetadataCmd.AddCommand(regionCmd)
	MetadataCmd.AddCommand(instanceIdCmd)
	MetadataCmd.AddCommand(instanceTypeCmd)
}
