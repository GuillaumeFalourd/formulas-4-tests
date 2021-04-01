#!/usr/bin/python3
import os

from formula import formula

path = os.environ.get("RIT_INPUT_PATH")

formula.Run(path)
