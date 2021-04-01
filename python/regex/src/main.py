#!/usr/bin/python3
import os

from formula import formula

age = os.environ.get("RIT_AGE")

formula.run(age)
