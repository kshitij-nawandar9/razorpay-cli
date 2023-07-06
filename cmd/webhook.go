package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var webhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "merchant webhook",
	Long:  `merchant webhook`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Webhook command")
	},
}
