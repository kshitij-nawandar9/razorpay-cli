package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const ordersURI = "/v1/orders"

var ordersCmd = &cobra.Command{
	Use:   "order",
	Short: "Payments order",
	Long:  "Payments order",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Payment Order command")
	},
}
