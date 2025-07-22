import os

import yaml

config_path = os.path.join(
    os.environ.get("HOME"), ".config", "media-discord-rpc", "config.yaml"
)

with open(config_path, "r") as file:
    app_config = yaml.safe_load(file)
