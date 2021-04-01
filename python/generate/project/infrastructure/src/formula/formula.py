from colored import fg, attr
import time
from progress.bar import ChargingBar

def run(name, pipeline_tool, commit_tool, build_tool, test_tool, deploy_tool, operation_tool):

    download('1️⃣  CREATING CI PIPELINE',100)
    print("🔄%s \033[1mCI PIPELINE CREATED USING {}\033[0m %s\n".format(pipeline_tool) % (fg(3), attr(0)))

    download('2️⃣  CONFIGURING WEBHOOK',125)
    print("🔗%s \033[1mWEBHOOK CONFIGURED WITH {}\033[0m %s\n".format(commit_tool) % (fg(3), attr(0)))

    download('3️⃣  CONFIGURING BUILD JOB',150)
    print("🔨%s \033[1mBUILD JOB CONFIGURED THROUGH {}\033[0m %s\n".format(build_tool) % (fg(3), attr(0)))

    download('4️⃣  CONFIGURING TESTS JOB',175)
    print("💻%s \033[1mTESTS JOB CONFIGURED USING {}\033[0m %s\n".format(test_tool) % (fg(3), attr(0)))

    download('5️⃣  CONFIGURING DEPLOY JOB',200)
    print("🆙%s \033[1mDEPLOY JOB CONFIGURED THROUGH {}\033[0m %s\n".format(deploy_tool) % (fg(3), attr(0)))

    download('6️⃣  CONFIGURING OPERATING JOB',225)
    print("🔎%s \033[1mOPERATING JOB CONFIGURED WITH {}\033[0m %s\n".format(operation_tool) % (fg(3), attr(0)))

    print("%s✅ \033[1m{} PROJECT INFRASTRUCTURE CONFIGURED SUCCESSFULLY\033[0m %s".format(name.upper()) % (fg(2), attr(0)))

def download(message, value):
    bar = ChargingBar(message, max=value)
    for i in range(value):
        time.sleep(0.010)
        bar.next()
    bar.finish()
