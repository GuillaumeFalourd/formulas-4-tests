#!/usr/bin/python3
from colored import fg, attr
import time
from progress.bar import ChargingBar

def run(project, version, environment, type):

    download('1️⃣  SEARCHING FOR PROJECT',100)
    print("🔎%s \033[1m{} PROJECT ENCOUNTERED\033[0m %s\n".format(project.upper()) % (fg(3), attr(0)))

    download('2️⃣  CHECKING VERSION',125)
    print("📦%s \033[1m{} VERSION AVAILABLE\033[0m %s\n".format(version) % (fg(3), attr(0)))

    download('3️⃣  PREPARING ENVIRONMENT',150)
    print("📋%s \033[1m{} ENVIRONMENT READY\033[0m %s\n".format(environment) % (fg(3), attr(0)))

    download('4️⃣  CONFIGURING DEPLOYMENT',175)
    print("🛠%s  \033[1m{} DEPLOYMENT CONFIGURED\033[0m %s\n".format(type) % (fg(3), attr(0)))

    download('5️⃣  DEPLOYING PROJECT',200)
    print("✅%s \033[1m{} PROJECT SUCCESSFULLY DEPLOYED\033[0m %s".format(project.upper()) % (fg(2), attr(0)))

def download(message, value):
    bar = ChargingBar(message, max=value)
    for i in range(value):
        time.sleep(0.010)
        bar.next()
    bar.finish()