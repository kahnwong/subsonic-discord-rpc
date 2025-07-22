import urllib.parse

import requests

from media_discord_rpc import app_config

SUBSONIC_API_ENDPOINT = app_config["SUBSONIC_API_ENDPOINT"]
params = {
    "u": app_config["SUBSONIC_USERNAME"],
    "t": app_config["SUBSONIC_TOKEN"],
    "s": app_config["SUBSONIC_SALT"],
    "v": "1.16.1",
    "c": "media-discord-rpc",
    "f": "json",
}


def _parse_track_info(track_info):
    return {
        "details": track_info["title"],
        # "album": track_info["album"],
        "state": track_info["artist"],
        "image": track_info["coverArt"],
    }


def get_now_playing(params=params):
    r = requests.get(
        f"{SUBSONIC_API_ENDPOINT}/rest/getNowPlaying", params=params
    ).json()

    try:
        r = r["subsonic-response"]["nowPlaying"]["entry"][0]

        track_info = _parse_track_info(r)
        track_info["image"] = (
            f"{SUBSONIC_API_ENDPOINT}/rest/getCoverArt?id={track_info['image']}&size=96&{urllib.parse.urlencode(params)}"
        )

        return track_info
    except KeyError:
        return None
