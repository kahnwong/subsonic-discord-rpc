package integrations

import (
	"context"
	"fmt"

	jellyfin "github.com/sj14/jellyfin-go/api"
)

var client *jellyfin.APIClient

type JellyfinNowPlaying struct {
	Title    string
	Episode  string
	CoverArt string
}

func JellyfinGetNowPlaying() (JellyfinNowPlaying, error) {
	sessions, _, err := client.SessionAPI.GetSessions(context.Background()).Execute()
	if err != nil {
		return JellyfinNowPlaying{}, err
	}

	var r JellyfinNowPlaying
	if len(sessions) > 0 {
		nowPlaying := sessions[0].NowPlayingItem.Get()
		imageTagsPrimary := nowPlaying.ImageTags["Primary"]

		r = JellyfinNowPlaying{
			Title:    *nowPlaying.Name.Get(),
			Episode:  "",
			CoverArt: fmt.Sprintf("%s/Items/%s/Images/Primary?fillHeight=100&tag=%s", AppConfig.JellyfinApiEndpoint, *nowPlaying.Id, imageTagsPrimary),
		}
	}

	return r, err
}

func init() {
	config := &jellyfin.Configuration{
		Servers: jellyfin.ServerConfigurations{{URL: AppConfig.JellyfinApiEndpoint}},
		DefaultHeader: map[string]string{
			"Authorization": fmt.Sprintf("MediaBrowser Token=\"%s\"", "736a20a8e3ba430bb75a4ad9009e6ef9"),
		},
	}

	client = jellyfin.NewAPIClient(config)
}
