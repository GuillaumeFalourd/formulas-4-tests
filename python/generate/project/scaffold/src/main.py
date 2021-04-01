#!/usr/bin/python3
import os

from formula import formula

name = os.environ.get("RIT_PROJECT_NAME")
language = os.environ.get("RIT_PROJECT_LANGUAGE")
architecture = os.environ.get("RIT_PROJECT_ARCHITECTURE")
dependencies = os.environ.get("RIT_PROJECT_DEPENDENCIES")

formula.run(name, language, architecture, dependencies)
