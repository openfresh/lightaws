package metadata

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	instanceTypeCmd = &cobra.Command{
		Use:   "instance-type",
		Short: "get own EC2 instance type",
		RunE: func(cmd *cobra.Command, args []string) error {

			instanceType, err := client.GetMetadata("instance-type")
			if err != nil {
				return err
			}

			fmt.Println(instanceType)
			return nil
		},
	}
)
