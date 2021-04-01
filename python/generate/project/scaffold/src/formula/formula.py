#!/usr/bin/python3
from colored import fg, attr
import time
from progress.bar import ChargingBar

def run(name, language, architecture, dependencies):
    download('1️⃣  CREATING DIRECTORIES',100)
    print("%s🗂  \033[1m{} ARCHITECTURE DIRECTORIES CREATED\033[0m %s\n".format(architecture) % (fg(3), attr(0)))

    download('2️⃣  GENERATING PROJECT CLASSES',125)
    print("%s📝 \033[1mPROJECT CLASSES IN {} GENERATED\033[0m %s\n".format(language) % (fg(3), attr(0)))

    download('3️⃣  CONFIGURING PROPERTIES',150)
    print("%s📊 \033[1mPROJECT PROPERTIES CONFIGURED\033[0m %s\n" % (fg(3), attr(0)))

    if dependencies == "YES":
        download('4️⃣  IMPORTING DEFAULT DEPENDENCIES',175)
        print("%s📚 \033[1mDEPENDENCIES IMPORTED\033[0m %s\n" % (fg(3), attr(0)))

    print("%s✅ \033[1m{} PROJECT CREATED SUCCESSFULLY\033[0m %s".format(name.upper()) % (fg(2), attr(0))
)

def download(message, value):
    bar = ChargingBar(message, max=value)
    for i in range(value):
        time.sleep(0.010)
        bar.next()
    bar.finish()
