package cmd

import (
	"time"

	"github.com/kahnwong/media-discord-rpc/discord"
	"github.com/kahnwong/rich-go/client"
	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

var writingCmd = &cobra.Command{
	Use:   "writing",
	Short: "Display writing activity",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			discord.SetActivity(discord.DiscordApps.Intellij, client.ActivityTypes.Playing, "Writing a blog post", "", "")

			log.Info().Msg("Writing...")
			time.Sleep(60 * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(writingCmd)
}
