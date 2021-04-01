#!/usr/bin/python3
import os
import json

def Run(numberOne, numberTwo, operation):
    input_json = json.dumps({'number_one': numberOne, 'number_two': numberTwo})

    if operation == "sum":
        stdin_cmd = f"echo '{input_json}' | rit python math sum numbers --stdin"

    elif operation == "multiplication":
        stdin_cmd = f"echo '{input_json}' |  rit python math multiply numbers --stdin"

    else:
        print("Unexpected operation")

    os.system(f'{stdin_cmd}')
