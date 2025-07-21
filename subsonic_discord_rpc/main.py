import os
import time

from pypresence import ActivityType, Presence

client_id = os.getenv("CLIENT_ID")
RPC = Presence(client_id)
RPC.connect()

RPC.update(
    activity_type=ActivityType.LISTENING,
    details="Aspiral",
    state="Epica",
    # end=int(300) + time.time(),
    # small_image="",
)

while True:
    time.sleep(15)
