package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var receipt string
var amount int

func init() {
	orderCreateCmd.Flags().StringVarP(&receipt, "receipt", "r", "Receipt no. 1", "Receipt of an order")
	orderCreateCmd.Flags().IntVarP(&amount, "amt", "a", 100, "Amount of an order")
}

var orderCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create order",
	Long:  "Create order",
	Run: func(cmd *cobra.Command, args []string) {
		method := "POST"
		payload := map[string]interface{}{
			"amount":   amount,
			"currency": "INR",
			"receipt":  receipt,
			"notes": map[string]interface{}{
				"notes_key_1": "Order for payment testing",
				"notes_key_2": "Order for payment testing",
			},
		}

		resp, err := makeRequest(context.TODO(), ordersURI, method, payload, os.Getenv(OsUsername), os.Getenv(OsSecret))

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
