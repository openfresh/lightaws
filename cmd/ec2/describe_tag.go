package ec2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
)

var (
	describeNameTagCmd = &cobra.Command{
		Use:   "describe-name-tag",
		Short: "get EC2 tag value",
		RunE: func(cmd *cobra.Command, args []string) error {

			params := &ec2.DescribeInstancesInput{
				InstanceIds: aws.StringSlice([]string{instanceID}),
			}

			res, err := client.DescribeInstances(params)
			if err != nil {
				return err
			}

			if len(res.Reservations) == 0 {
				return fmt.Errorf("describe-instances %s is failed.", instanceID)
			}

			for _, i := range res.Reservations[0].Instances {
				var nameTag string
				for _, t := range i.Tags {
					if *t.Key == "Name" {
						nameTag = *t.Value
						break
					}
				}
				fmt.Println(nameTag)
			}

			return nil
		},
	}
)
