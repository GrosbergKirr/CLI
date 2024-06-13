/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"mycli/internal/consts"
	"mycli/internal/grpc_client"
	"os"
)

var dnsparam consts.Consts

var DNSSet = &cobra.Command{
	Use:   "dnsset",
	Short: "Set dns mame.\n Use -n [DNS/DNS/DNS...] to set DNSList.\n Use -a [address] to set server address.\n Use -p [password] to set server user password.",
	Long:  "This command set DNSList you want. Use -n [DNS/DNS/DNS...] to set DNSList.\n Use -a [address] to set server address.\n Use -p [password] to set server user password.",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		c, err := grpc_client.NewClient(dnsparam.Address)
		if err != nil {
			fmt.Printf("cant create client: %s", err)
			os.Exit(1)
		}
		r, err := c.DNSChange(ctx, dnsparam.DNSList, dnsparam.Address, dnsparam.Password)
		if err != nil {
			fmt.Printf("cant change DNSList: %s", err)
		}
		fmt.Println(r)

	},
}

func init() {
	rootCmd.AddCommand(DNSSet)

	DNSSet.Flags().StringVarP(&dnsparam.DNSList, "dnsname", "n", "DNSList", "DNS to set")
	DNSSet.Flags().StringVarP(&dnsparam.Address, "dnsaddress", "a", "localhost:9090", "Address of server")
	DNSSet.Flags().StringVarP(&dnsparam.Password, "dnspassword", "p", "21071999", "Sudo password")
}
