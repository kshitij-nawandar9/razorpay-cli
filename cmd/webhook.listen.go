package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

var webhookListenCmd = &cobra.Command{
	Use:   "listen",
	Short: "webhookListen testing",
	Long:  `webhookListen testing`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		tun, err := ngrok.Listen(ctx,
			config.HTTPEndpoint(),
			ngrok.WithAuthtokenFromEnv(),
		)
		if err != nil {
			log.Println(err)
		}

		log.Println("tunnel created:", tun.URL())

		err = http.Serve(tun, http.HandlerFunc(handler))
		if err != nil {
			log.Println(err)
		}
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("webhook received")
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Request Body:", string(body))
	fmt.Fprintln(w, "<h1>Hello from ngrok-go.</h1>")
}
