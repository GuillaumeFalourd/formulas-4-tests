#!/usr/bin/python3
import os

from formula import formula

input1 = os.environ.get("NUMBER_ONE")
input2 = os.environ.get("NUMBER_TWO")
formula.Run(input1, input2)
