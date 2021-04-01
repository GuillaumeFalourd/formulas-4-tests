#!/usr/bin/python3
import requests

def run():
    url = "https://api.github.com/users/andrew/repos"
    result = get_repositories(url)
    print(result)
    print(len(result))

def get_repositories(url):
    result = []
    r = requests.get(url=url)
    if 'next' in r.links :
        result += get_repositories(r.links['next']['url'])

    for repository in r.json():
        result.append(repository.get('name'))

    return result
