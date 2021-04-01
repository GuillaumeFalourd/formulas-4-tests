#!/usr/bin/python3
import os

from formula import formula_horusec
from formula import formula_superlinter

def run():
    current_pwd = os.environ.get("CURRENT_PWD")
    os.chdir(current_pwd)
    if not os.path.exists(".github"):
        os.makedirs(".github")
        os.chdir('.github')
    else:
        os.chdir('.github')
    
    if not os.path.exists("workflows"):
        os.makedirs("workflows")
        os.chdir('workflows')
    else:
        os.chdir('workflows')
    
    print("Creating Worflows...")
    formula_horusec.run()
    formula_superlinter.run()
