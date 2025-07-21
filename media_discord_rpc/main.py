import os
import time

from pypresence import ActivityType, Presence

from media_discord_rpc import subsonic

client_id = os.getenv("DISCORD_CLIENT_ID")
RPC = Presence(client_id)
RPC.connect()

while True:
    now_playing = subsonic.get_now_playing()
    print(now_playing)

    if now_playing:
        RPC.update(
            activity_type=ActivityType.LISTENING,
            details=now_playing["title"],
            state=now_playing["artist"],
            # end=int(300) + time.time(),
            small_image=now_playing["coverArt"],
        )
    else:
        print("Nothing is currently playing...")

    time.sleep(15)
