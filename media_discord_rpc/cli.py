import time

import click
from pypresence import ActivityType, Presence

from media_discord_rpc import app_config, jellyfin, subsonic

client_id = app_config["DISCORD_CLIENT_ID"]
RPC = Presence(client_id)
RPC.connect()


@click.group()
@click.version_option()
def cli():
    ""


@cli.command(name="media")
def media():
    "Display listening/watching activity"
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


@cli.command(name="writing")
def writing():
    "Display writing activity"
    while True:
        RPC.update(
            activity_type=ActivityType.PLAYING,
            details="Writing a blog post",
            state=None,
            small_image="https://github.com/kahnwong/dashboard-icons/blob/master/rpc/tea.png?raw=true",
        )

        print("Writing a blog post...")
        time.sleep(30)
