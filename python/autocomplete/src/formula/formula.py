#!/usr/bin/python3
from colored import fg, attr
from distutils.util import strtobool


def Run(path):
    print("Hello World!")
    print(f"{fg(2)}The informed path is: {path}.{attr(0)}")
