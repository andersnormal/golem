package cmd

import (
	"net/http"

	"github.com/andersnormal/golem/pkg/xcode"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

var (
	botName string
)

func init() {
	CreateBot.Flags().StringVar(&botName, "name", "", "bot name")
}

var CreateBot = &cobra.Command{
	Use:   "create-bot",
	Short: "Creates a new bot",
	Run: func(cmd *cobra.Command, args []string) {
		httpClient := &http.Client{}

		xcodeClient := xcode.New(httpClient)

		req := &xcode.CreateBotRequest{
			Name:          botName,
			Configuration: &xcode.Configuration{},
		}

		success, failure, err := xcodeClient.Bot.Create(req)
		if err != nil {
			spew.Dump(failure, err)

			return
		}

		spew.Dump(success)
	},
}
