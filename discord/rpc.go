package discord

import (
	"os"

	"github.com/kahnwong/rich-go/client"
	"github.com/rs/zerolog/log"
)

func SetActivity(activityType int, state string, details string, smallImage string) {
	err := client.Login(os.Getenv("DISCORD_APP_ID"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to login to Discord")
	}

	err = client.SetActivity(client.Activity{
		ActivityType: activityType,
		State:        state,
		Details:      details,
		SmallImage:   smallImage,
	})

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to set activity")
	}
}
