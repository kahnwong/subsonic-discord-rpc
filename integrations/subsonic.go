package integrations

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/carlmjohnson/requests"
	"github.com/google/go-querystring/query"
	"github.com/rs/zerolog/log"
)

// auth
type SubsonicAuth struct {
	Username string `url:"u"`
	Token    string `url:"t"`
	Salt     string `url:"s"`
	Version  string `url:"v"`
	Client   string `url:"c"`
	Format   string `url:"f"`
}

var authParams url.Values

// structs
type NowPlaying struct {
	SubsonicResponse struct {
		NowPlaying struct {
			Entry []struct {
				Title    string `json:"title"`
				Album    string `json:"album"`
				Artist   string `json:"artist"`
				CoverArt string `json:"coverArt"`
			} `json:"entry"`
		} `json:"nowPlaying"`
	} `json:"subsonic-response"`
}

type NowPlayingParsed struct {
	Title    string
	Artist   string
	CoverArt string
}

func subsonicGetNowPlaying() NowPlaying {
	var response NowPlaying
	err := requests.
		URL(AppConfig.SubsonicApiEndpoint).
		Method(http.MethodGet).
		Path("rest/getNowPlaying").
		Params(authParams).
		ToJSON(&response).
		Fetch(context.Background())

	if err != nil {
		log.Error().Msg("Failed to get NowPlaying")
	}

	return response
}

func SubsonicNowPlaying() NowPlayingParsed {
	nowPlaying := subsonicGetNowPlaying()

	if len(nowPlaying.SubsonicResponse.NowPlaying.Entry) > 0 {
		coverArt := nowPlaying.SubsonicResponse.NowPlaying.Entry[0].CoverArt

		return NowPlayingParsed{
			Title:    nowPlaying.SubsonicResponse.NowPlaying.Entry[0].Title,
			Artist:   nowPlaying.SubsonicResponse.NowPlaying.Entry[0].Artist,
			CoverArt: fmt.Sprintf("%s/rest/getCoverArt?id=%s&size=120&%s", AppConfig.SubsonicApiEndpoint, coverArt, authParams.Encode()),
		}
	} else {
		return NowPlayingParsed{}
	}
}

func init() {
	var authValues = SubsonicAuth{
		Username: AppConfig.SubsonicUsername,
		Token:    AppConfig.SubsonicToken,
		Salt:     AppConfig.SubsonicSalt,
		Version:  "1.16.1",
		Client:   "media-discord-rpc",
		Format:   "json",
	}

	authParams, _ = query.Values(authValues)
}
