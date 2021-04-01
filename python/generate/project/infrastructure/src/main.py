#!/usr/bin/python3
import os

from formula import formula

name = os.environ.get("RIT_PROJECT_NAME")
pipeline_tool = os.environ.get("RIT_PIPELINE_TOOL")
commit_tool = os.environ.get("RIT_COMMIT_TOOL")
build_tool = os.environ.get("RIT_BUILD_TOOL")
test_tool = os.environ.get("RIT_TEST_TOOL")
deploy_tool = os.environ.get("RIT_DEPLOY_TOOL")
operation_tool = os.environ.get("RIT_OPERATION_TOOL")

formula.run(name, pipeline_tool, commit_tool, build_tool, test_tool, deploy_tool, operation_tool)
