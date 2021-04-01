#!/usr/bin/python3
import json

class Github:
	username = 'username'
	token = 'token'
	email = 'email'

def Run(username, token, email):
    #create object
    github = Github()
    github.username = username
    github.token = token
    github.email = email

    #convert to JSON string
    github_json = json.dumps(github.__dict__)

    #print json string
    print(github_json)
