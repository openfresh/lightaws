package ecr

import (
	"errors"

	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/spf13/cobra"
)

var (
	repositoryName     string
	createReositoryCmd = &cobra.Command{
		Use:  "create-repository",
		Long: "Creates an image repository.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if repositoryName == "" {
				return errors.New("argument --repository-name is required")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			input := ecr.CreateRepositoryInput{
				RepositoryName: aws.String(repositoryName),
			}
			output, err := client.CreateRepository(&input)
			if err != nil {
				return err
			}

			res, err := json.Marshal(output)
			if err != nil {
				return err
			}

			cmd.Println(string(res))
			return nil
		},
	}
)

func init() {

	createReositoryCmd.Flags().StringVar(&repositoryName, "repository-name", "", `The name to use for the repository. 
        The repository name may be specified on its own (such as nginx-web-app ) 
        or it can be prepended with a namespace to group the repository into a category 
        (such as project-a/nginx-web-app )`)

}
