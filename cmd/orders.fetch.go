package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var orderId string

func init() {
	orderFetchCmd.Flags().StringVarP(&orderId, "id", "i", "order_MAcMQWQ5eQoFiC", "Order Id of a payment")
}

var orderFetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch Orders by Id",
	Long:  "Fetch Orders by Id",
	Run: func(cmd *cobra.Command, args []string) {
		method := "GET"
		payload := []byte(``)

		resp, err := makeRequest(context.TODO(), ordersURI+"/"+orderId, method, payload)

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
