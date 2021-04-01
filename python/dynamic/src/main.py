#!/usr/bin/python3
import os

from formula import formula

dynamic = os.environ.get("RIT_DYNAMIC")

formula.run(dynamic)
