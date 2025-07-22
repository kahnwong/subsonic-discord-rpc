from jellyfin_apiclient_python import JellyfinClient

from media_discord_rpc import app_config

client = JellyfinClient()
client.config.data["auth.ssl"] = True
client.config.data["app.name"] = "media-discord-rpc"
client.config.data["app.version"] = "0.0.1"


def jellyfin_get_now_playing():
    client.authenticate(
        {
            "Servers": [
                {
                    "AccessToken": app_config["JELLYFIN_API_KEY"],
                    "address": app_config["JELLYFIN_ENDPOINT"],
                }
            ]
        },
        discover=False,
    )

    sessions_r = client.jellyfin.sessions()
    try:
        session_id = sessions_r[0]["Id"]
        now_playing = client.jellyfin.get_now_playing(session_id)

        title = now_playing["Name"]
        image = f"{app_config['JELLYFIN_ENDPOINT']}/Items/{now_playing['PlayState']['MediaSourceId']}/Images/Primary?fillHeight=100&tag={now_playing['ImageTags']['Primary']}"
        r = {"details": title, "state": None, "image": image}

        if now_playing["Type"] == "Episode":
            r["details"] = now_playing["SeriesName"]
            r["state"] = (
                f"S{now_playing['ParentIndexNumber']}E{now_playing['IndexNumber']}"
            )
            r["image"] = (
                f"{app_config['JELLYFIN_ENDPOINT']}/Items/{now_playing['ParentThumbItemId']}/Images/Primary?fillHeight=100&tag={now_playing['SeriesPrimaryImageTag']}"
            )

        return r
    except:  # noqa
        return None
