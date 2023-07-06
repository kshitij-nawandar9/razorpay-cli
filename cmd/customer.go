package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

const URI = "/v1/customers"

var customerCmd = &cobra.Command{
	Use:   "customer",
	Short: "Fetch all customers",
	Long:  `Fetch all customers`,
	Run: func(cmd *cobra.Command, args []string) {
		method := "GET"
		payload := []byte(``)

		resp, err := makeRequest(context.TODO(), URI, method, payload)

		if err != nil {
			log.Fatal(err)
		}

		val, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(val))

	},
}
