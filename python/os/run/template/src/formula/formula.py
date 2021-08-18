#!/usr/bin/python3
#!/usr/bin/python3
import requests
import inquirer
import git
import os

def run():
    url = 'https://raw.githubusercontent.com/fabianofernandeszup/backend-formulas-template-engine/main/templates.txt'
    r = requests.get(url, allow_redirects=True)

    open('templates.txt', 'wb').write(r.content)
    file1 = open('templates.txt', 'r')
    Lines = file1.readlines()
    StripLines = []

    for line in Lines:
        stripline = line.strip()
        StripLines.append(stripline)

    file1.close()
    os.remove("templates.txt")

    question = [
            inquirer.Checkbox("templates",
                    message = f"\033[1mSelect templates:\033[0m ",
                    choices = StripLines,
                ),
        ]
    answer = inquirer.prompt(question)
    templates = answer["templates"]

    for t in templates:
        print(f"🛠  \033[36mCloning {t}\033[0m repository...\033[0m")
        gitRepoUrl = t + ".git"
        print(f"URL: {gitRepoUrl}")
        git.Git("").clone(gitRepoUrl)
        print(f"Repository cloned 🚀\033[0m\n")

    print(f"Selected templates cloned successfully ✅")
