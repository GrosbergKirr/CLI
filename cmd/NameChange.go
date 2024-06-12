package cmd

import (
	"Yadro_CLI/internal/grpc_client"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var name, addr, password string

var NameChange = &cobra.Command{
	Use:   "setname",
	Short: "Greet someone",
	Long:  `This command greets the person specified by the name flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		c, err := grpc_client.NewClient(addr)
		if err != nil {
			fmt.Printf("cant create client: %s", err)
			os.Exit(1)
		}
		r, err := c.ChangeHostName(ctx, name, addr, password)
		if err != nil {
			fmt.Printf("cant change name: %s", err)
		}
		fmt.Println(r)

	},
}

func init() {
	rootCmd.AddCommand(NameChange)

	// Define flags and configuration settings.
	NameChange.Flags().StringVarP(&name, "name", "n", "RandomName", "Name to set")
	NameChange.Flags().StringVarP(&addr, "address", "a", "localhost:9090", "Address of server")
	NameChange.Flags().StringVarP(&password, "password", "p", "21071999", "Sudo password")
}
