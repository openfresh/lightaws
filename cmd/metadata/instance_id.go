package metadata

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	instanceIdCmd = &cobra.Command{
		Use:   "instance-id",
		Short: "get own EC2 instance ID",
		RunE: func(cmd *cobra.Command, args []string) error {

			instanceID, err := client.GetMetadata("instance-id")
			if err != nil {
				return err
			}

			fmt.Println(instanceID)
			return nil
		},
	}
)
