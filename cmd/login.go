package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type loginCmd struct {
	cmd *cobra.Command

	username string
	secret   string
}

func newLoginCmd() *loginCmd {
	lc := &loginCmd{}
	lc.cmd = &cobra.Command{
		Use:   "login to the system",
		Short: "login to razorpay cli",
		Long:  `login to razorpay cli`,
		RunE:  lc.runTriggerCmd,
	}

	lc.cmd.Flags().StringVar(&lc.username, "username", "", "username for APIs")
	lc.cmd.Flags().StringVar(&lc.secret, "secret", "", "secret for APIs")

	return lc
}

func (lc *loginCmd) runTriggerCmd(cmd *cobra.Command, args []string) error {

	fmt.Println(lc.username)
	fmt.Println(lc.secret)

	if err := os.Setenv(OsUsername, lc.username); err != nil {
		os.Getenv(OsUsername)
		log.Fatal(err)
		return err
	}
	if err := os.Setenv(OsSecret, lc.secret); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
