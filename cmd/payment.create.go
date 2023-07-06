package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var method, vpa string
var amt int

func init() {
	paymentCreateCmd.Flags().IntVarP(&amt, "amount", "a", 0, "Amount of payment")
	paymentCreateCmd.MarkFlagRequired("amount")

	paymentCreateCmd.Flags().StringVarP(&method, "method", "m", "upi", "Method of payment")

	paymentCreateCmd.Flags().StringVarP(&vpa, "vpa", "v", "", "VPA of payment")
	paymentCreateCmd.MarkFlagRequired("vpa")
}

var paymentCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create",
	Long:  `create`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://api-dark.razorpay.com/v1/payments/create/ajax"
		method := "POST"
		payloadData := map[string]interface{}{
			"contact":           "9404237451",
			"email":             "manask.322@gmail.com",
			"amount":            amt,
			"method":            method,
			"vpa":               vpa,
			"description":       "iosdsds",
			"force_terminal_id": "term_Kxgls6GTnd88bu",
			"bank":              "HDFC",
			"_": map[string]interface{}{
				"library": "cli",
			},
			"upi": map[string]interface{}{
				"expiry_time": 10,
			},
			"notes": map[string]interface{}{
				"transaction_id": "geddit-339B1HBHH53685-1",
				"txn_uuid":       "mozet9eUrrQb1ZXDdeN",
			},
			"currency": "INR",
		}
		payload, _ := json.Marshal(payloadData)
		headers := map[string]string{
			"Content-Type":  "application/json",
			"Authorization": os.Getenv("BASIC_AUTH"),
		}

		req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		for key, value := range headers {
			req.Header.Set(key, value)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		fmt.Println("Response Body:", string(body))
	},
}
