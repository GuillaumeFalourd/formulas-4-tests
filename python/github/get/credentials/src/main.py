#!/usr/bin/python3
import os

from formula import formula

username = os.environ.get("GIT_USER")
token = os.environ.get("GIT_TOKEN")
email = os.environ.get("GIT_EMAIL")
formula.Run(username, token, email)
