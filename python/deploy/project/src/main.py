#!/usr/bin/python3
import os

from formula import formula

project = os.environ.get("RIT_PROJECT")
version = os.environ.get("RIT_VERSION")
environment = os.environ.get("RIT_ENVIRONMENT")
type = os.environ.get("RIT_DEPLOYMENT_TYPE")

formula.run(project, version, environment, type)
