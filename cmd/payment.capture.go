package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var captureAmount int
var paymentID string

func init() {
	paymentCaptureCmd.Flags().IntVarP(&captureAmount, "amount", "a", 100, "Amount of payment")
	paymentCaptureCmd.Flags().StringVarP(&paymentID, "payment_id", "p", "pay_MAmNSwTmksYlQY", "Payment Id")
	paymentCaptureCmd.MarkFlagRequired("amount")
	paymentCaptureCmd.MarkFlagRequired("paymentId")
}

var paymentCaptureCmd = &cobra.Command{
	Use:   "capture",
	Short: "Capture the payment",
	Long:  `Capture the payment`,
	Run: func(cmd *cobra.Command, args []string) {
		method := "POST"
		payloadData := map[string]interface{}{
			"amount":   captureAmount,
			"currency": "INR",
		}

		fmt.Println(payloadData)

		resp, err := makeRequest(context.TODO(), paymentURI+"/"+paymentID+"/"+"capture", method, payloadData, os.Getenv(OsUsername), os.Getenv(OsSecret))

		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		val, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(val))
	},
}
