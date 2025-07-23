package cmd

import (
	"fmt"
	"time"

	"github.com/kahnwong/media-discord-rpc/discord"
	"github.com/kahnwong/media-discord-rpc/integrations"
	"github.com/kahnwong/rich-go/client"
	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

var watchingCmd = &cobra.Command{
	Use:   "watching",
	Short: "Display watching activity",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			nowPlaying, _ := integrations.JellyfinGetNowPlaying()
			fmt.Println(nowPlaying)

			if nowPlaying.Title != "" {
				discord.SetActivity(client.ActivityTypes.Listening, nowPlaying.Episode, nowPlaying.Title, nowPlaying.CoverArt)
			} else {
				log.Info().Msg("Nothing is currently playing...")
			}

			time.Sleep(15 * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(watchingCmd)
}
