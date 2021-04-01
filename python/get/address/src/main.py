#!/usr/bin/python3
import os

from formula import formula

cep = os.environ.get("CEP")
formula.Run(cep)
