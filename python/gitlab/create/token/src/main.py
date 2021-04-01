#!/usr/bin/python3
import os

from formula import formula

api_token = os.environ.get("RIT_GITLAB_TOKEN")
user_id = os.environ.get("RIT_GITLAB_USER")
token_name = os.environ.get("RIT_TOKEN_NAME")

formula.run(api_token, user_id, token_name)
