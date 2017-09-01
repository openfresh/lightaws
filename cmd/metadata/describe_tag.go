package metadata

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	describeTagCmd = &cobra.Command{
		Use:   "describe-tag",
		Short: "get own tag",
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
