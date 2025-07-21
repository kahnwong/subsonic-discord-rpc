import os
import urllib.parse

import requests

BASE_URL = os.environ["BASE_URL"]
params = {
    "u": os.environ["USERNAME"],
    "t": os.environ["TOKEN"],
    "s": os.environ["SALT"],
    "v": "1.16.1",
    "c": "github-readme",
    "f": "json",
}


def _parse_track_info(track_info):
    return {
        "title": track_info["title"],
        "album": track_info["album"],
        "artist": track_info["artist"],
        "coverArt": track_info["coverArt"],
    }


def get_now_playing(params=params):
    r = requests.get(f"{BASE_URL}/rest/getNowPlaying", params=params).json()

    try:
        r = r["subsonic-response"]["nowPlaying"]["entry"][0]

        track_info = _parse_track_info(r)
        track_info["coverArt"] = (
            f"{BASE_URL}/rest/getCoverArt?id={track_info['coverArt']}&size=96&{urllib.parse.urlencode(params)}"
        )

        return track_info
    except KeyError:
        return None
