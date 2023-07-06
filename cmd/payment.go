package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var paymentCmd = &cobra.Command{
	Use:   "payment",
	Short: "payment",
	Long:  `payment`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Payment command")
	},
}
