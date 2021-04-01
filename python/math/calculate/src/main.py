#!/usr/bin/python3
import os

from formula import formula

numberOne = os.environ.get("NUMBER_ONE")
numberTwo = os.environ.get("NUMBER_TWO")
operation = os.environ.get("OPERATION")

formula.Run(numberOne, numberTwo, operation)
