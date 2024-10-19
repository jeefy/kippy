package config

import (
	"github.com/jeefy/kippy/pkg/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	Debug          = "debug"
	Email          = "email"
	DiscordWebhook = "discordWebhook"
	SlackWebhook   = "slackWebhook"
	GenericWebhook = "genericWebhook"
	SendgridAPIKey = "sendgridAPIKey"
)

var Sinks []types.KippySink

func LoadConfig(cmd *cobra.Command) {
	viper.SetDefault(Debug, false)

	viper.BindEnv(Email, "EMAIL")
	viper.BindEnv(DiscordWebhook, "DISCORD_WEBHOOK")
	viper.BindEnv(SlackWebhook, "SLACK_WEBHOOK")
	viper.BindEnv(GenericWebhook, "GENERIC_WEBHOOK")

	Sinks = []types.KippySink{
		{Type: "discord", Config: viper.GetString("discordWebhook")},
		{Type: "slack", Config: viper.GetString("slackWebhook")},
		{Type: "webhook", Config: viper.GetString("genericWebhook")},
		{Type: "email", Config: viper.GetString("email")},
	}
}
