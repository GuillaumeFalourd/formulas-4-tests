#!/usr/bin/python3
import os

from formula import formula

input1 = int(os.environ.get("NUMBER_ONE"))
input2 = int(os.environ.get("NUMBER_TWO"))
formula.Run(input1, input2)
