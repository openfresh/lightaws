package ecr

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/spf13/cobra"

	"encoding/base64"
	"strings"
)

var (
	registryIds []string
	getLoginCmd = &cobra.Command{
		Use: "get-login",
		Long: `Log in to an Amazon ECR registry.

This  command  retrieves a token that is valid for a specified registry 
for 12 hours, and then it prints  a  docker  login  command  with  that
authorization  token.  You can execute the printed command to log in to
your registry with Docker. After you have logged in to  an  Amazon  ECR
registry with this command, you can use the Docker CLI to push and pull
images from that registry until the token expires.

NOTE:
    This command writes displays docker login commands  to  stdout  with
    authentication  credentials.  Your  credentials  could be visible by
    other users on your system in a process list display  or  a  command
    history. If you are not on a secure system, you should consider this
    risk and login interactively. For more information,  see  get-author
    ization-token.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			input := ecr.GetAuthorizationTokenInput{}
			if len(registryIds) > 0 {
				input.RegistryIds = aws.StringSlice(registryIds)
			}
			output, err := client.GetAuthorizationToken(&input)
			if err != nil {
				return err
			}

			if len(output.AuthorizationData) == 0 {
				return errors.New("get token is faield")
			}

			target := output.AuthorizationData[0]

			data, err := base64.StdEncoding.DecodeString(*target.AuthorizationToken)
			if err != nil {
				return err
			}

			encoded := string(data)
			tokens := strings.Split(encoded, ":")
			if len(tokens) != 2 {
				return fmt.Errorf("authorization token is invalid format. [%s]", encoded)
			}

			loginCommand := fmt.Sprintf("docker login -u %s -p %s -e none %s", tokens[0], tokens[1], *target.ProxyEndpoint)
			cmd.Println(loginCommand)
			return nil
		},
	}
)

func init() {
	getLoginCmd.Flags().StringSliceVar(&registryIds, "registry-ids", make([]string, 0), "")
}
