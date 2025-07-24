package discord

import (
	"github.com/kahnwong/media-discord-rpc/config"
	"github.com/kahnwong/rich-go/client"
	"github.com/rs/zerolog/log"
)

type DiscordApp struct {
	Subsonic string
	Jellyfin string
	Intellij string
}

var DiscordApps = DiscordApp{
	Subsonic: config.AppConfig.DiscordAppIDSubsonic,
	Jellyfin: config.AppConfig.DiscordAppIDJellyfin,
	Intellij: config.AppConfig.DiscordAppIDIntellij,
}

func SetActivity(app string, activityType int, state string, details string, largeImage string) {
	err := client.Login(app)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to login to Discord")
	}

	err = client.SetActivity(client.Activity{
		ActivityType: activityType,
		State:        state,
		Details:      details,
		LargeImage:   largeImage,
	})

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to set activity")
	}
}
