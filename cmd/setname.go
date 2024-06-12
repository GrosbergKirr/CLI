package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"mycli/internal/consts"
	"mycli/internal/grpc_client"
	"os"
)

var param consts.Consts
var NameChange = &cobra.Command{
	Use:   "setname",
	Short: "Set host mame. Use -n [newname] to set New host name.\n Use -a [address] to set server address.\n Use -p [password] to set server user password.",
	Long: "This command set name on service which you want to set host mame. Use -n [newname] to set New host name.\n " +
		"Use -a [address] to set server address.\n Use -p [password] to set server user password.",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		c, err := grpc_client.NewClient(param.Address)
		if err != nil {
			fmt.Printf("cant create client: %s", err)
			os.Exit(1)
		}
		r, err := c.ChangeHostName(ctx, param.HostName, param.Address, param.Password)
		if err != nil {
			fmt.Printf("cant change name: %s", err)
		}
		fmt.Println(r)

	},
}

func init() {
	rootCmd.AddCommand(NameChange)

	NameChange.Flags().StringVarP(&param.HostName, "name", "n", "RandomName", "Name to set")
	NameChange.Flags().StringVarP(&param.Address, "address", "a", "localhost:9090", "Address of server")
	NameChange.Flags().StringVarP(&param.Password, "password", "p", "21071999", "Sudo password")

}
