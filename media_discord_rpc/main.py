import os
import time

from pypresence import ActivityType, Presence

from media_discord_rpc import jellyfin, subsonic

client_id = os.getenv("DISCORD_CLIENT_ID")
RPC = Presence(client_id)
RPC.connect()

while True:
    subsonic_now_playing = subsonic.get_now_playing()
    print(subsonic_now_playing)

    jellyfin_now_playing = jellyfin.jellyfin_get_now_playing()
    print(jellyfin_now_playing)

    now_playing = None
    activity_type = None
    if subsonic_now_playing:
        now_playing = subsonic_now_playing
        activity_type = ActivityType.LISTENING
    elif jellyfin_now_playing:
        now_playing = jellyfin_now_playing
        activity_type = ActivityType.WATCHING

    if now_playing:
        RPC.update(
            activity_type=activity_type,
            details=now_playing["details"],
            state=now_playing["state"],
            # end=int(300) + time.time(),
            small_image=now_playing["image"],
        )
    else:
        print("Nothing is currently playing...")

    time.sleep(30)
