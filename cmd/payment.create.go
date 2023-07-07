package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/stripe/stripe-cli/pkg/ansi"
)

var paymentMethod, vpa string
var amt int

func init() {
	paymentCreateCmd.Flags().IntVarP(&amt, "amount", "a", 0, "Amount of payment")
	paymentCreateCmd.MarkFlagRequired("amount")

	paymentCreateCmd.Flags().StringVarP(&paymentMethod, "method", "m", "upi", "Method of payment")

	paymentCreateCmd.Flags().StringVarP(&vpa, "vpa", "v", "", "VPA of payment")
}

var paymentCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create",
	Long:  `create`,
	Run: func(cmd *cobra.Command, args []string) {

		s := ansi.StartNewSpinner("creating upi payment ...", os.Stdout)

		url := "https://api-dark.razorpay.com/v1/payments/create/upi"
		method := "POST"
		payloadData := map[string]interface{}{
			"contact":           "9404237451",
			"email":             "manask.322@gmail.com",
			"amount":            amt,
			"method":            paymentMethod,
			"description":       "iosdsds",
			"force_terminal_id": "term_Kxgls6GTnd88bu",
			"bank":              "HDFC",
			"_": map[string]interface{}{
				"library": "cli",
			},
			"notes": map[string]interface{}{
				"transaction_id": "geddit-339B1HBHH53685-1",
				"txn_uuid":       "mozet9eUrrQb1ZXDdeN",
			},
			"currency": "INR",
		}

		if len(vpa) == 0 {
			payloadData["upi"] = map[string]string{
				"flow": "intent",
			}
		} else {
			payloadData["vpa"] = vpa
		}
		payload, _ := json.Marshal(payloadData)
		headers := map[string]string{
			"Content-Type": "application/json",
			// "Authorization": os.Getenv("BASIC_AUTH"),
		}

		req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
		req.SetBasicAuth("rzp_live_ruiXfILw0kpEXc", "X5i5mEBSVn82de7DoP3SUBRx")
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

		// fmt.Println("Response Body:", string(body))

		// resp, err := makeRequest(context.Background(), url, method, payloadData)

		// if err != nil {
		// 	log.Fatal(err)
		// }

		respBody, _ := ioutil.ReadAll(resp.Body)

		responseBody := map[string]interface{}{}

		_ = json.Unmarshal(respBody, &responseBody)

		val, err := json.MarshalIndent(responseBody, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		ansi.StopSpinner(s, "DONE!", os.Stdout)

		fmt.Println(string(val))
	},
}
