#!/usr/bin/python3
import os

from formula import formula

image_path = os.environ.get("RIT_IMAGE_PATH")
real_width = os.environ.get("RIT_REAL_WIDTH")

formula.run(image_path, real_width)
