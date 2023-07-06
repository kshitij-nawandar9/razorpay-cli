/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "razorpay-cli",
	Short: "command line interface for razorpay APIs",
	Long: `Command Line Interface to check how razorpay API works. It also have capabilities to test webhook integration and 
	testing the webhook consumption. Details of all the APIs can be found at https://razorpay.com/docs/api`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(customerCmd)

	rootCmd.AddCommand(webhookCmd)
	webhookCmd.AddCommand(webhookListenCmd)

	rootCmd.AddCommand(paymentCmd)
	paymentCmd.AddCommand(paymentCreateCmd)
}
