package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var paymentId string

func init() {
	paymentFetchCmd.Flags().StringVarP(&paymentId, "id", "i", "pay_MAVhcpLPpG00kd", "Payment Id")
}

var paymentFetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch Payment by Id",
	Long:  "Fetch Payment by Id",
	Run: func(cmd *cobra.Command, args []string) {
		method := "GET"
		payload := []byte(``)

		resp, err := makeRequest(context.TODO(), paymentURI+"/"+paymentId, method, payload, os.Getenv(OsUsername), os.Getenv(OsSecret))

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
