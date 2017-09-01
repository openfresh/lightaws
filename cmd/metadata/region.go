package metadata

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	regionCmd = &cobra.Command{
		Use:   "region",
		Short: "get region which EC2 instance belongs",
		RunE: func(cmd *cobra.Command, args []string) error {

			region, err := client.Region()
			if err != nil {
				return err
			}

			fmt.Println(region)
			return nil
		},
	}
)
