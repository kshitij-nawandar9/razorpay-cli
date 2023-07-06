package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/stripe/stripe-cli/pkg/ansi"
)

const URI = "/v1/customers"

var customerCmd = &cobra.Command{
	Use:   "customer",
	Short: "Fetch all customers",
	Long:  `Fetch all customers`,
	Run: func(cmd *cobra.Command, args []string) {
		method := "GET"
		payload := []byte(``)

		s := ansi.StartNewSpinner("Loading Customers ...", os.Stdout)

		resp, err := makeRequest(context.TODO(), URI, method, payload, os.Getenv(OsUsername), os.Getenv(OsSecret))

		if err != nil {
			log.Fatal(err)
		}

		val, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		ansi.StopSpinner(s, "", os.Stdout)

		fmt.Println(string(val))

	},
}
