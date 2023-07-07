package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/stripe/stripe-cli/pkg/ansi"
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

		s := ansi.StartNewSpinner("fetching payment ...", os.Stdout)
		time.Sleep(1 * time.Second)
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
		ansi.StopSpinner(s, "DONE!", os.Stdout)
		fmt.Println(string(val))
	},
}
