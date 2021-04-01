#!/usr/bin/python3
import os

from formula import formula

multi = os.environ.get("RIT_INPUT_MULTISELECT")

formula.run(multi)
