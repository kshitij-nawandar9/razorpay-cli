package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var webhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "webhook testing",
	Long:  `webhook testing`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Webhook command")
	},
}
