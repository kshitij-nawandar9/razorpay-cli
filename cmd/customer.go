package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var customerCmd = &cobra.Command{
	Use:   "customer",
	Short: "Fetch all customers",
	Long:  `Fetch all customers`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://api.razorpay.com/v1/customers"
		method := "GET"
		payload := []byte(``)
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
