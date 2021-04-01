#!/usr/bin/python3
import requests
import json
from datetime import date

def run(api_token, user_id, token_name):

    url = f'https://gitlab.com/api/v4/users/{user_id}/personal_access_tokens'
    
    today = date.today()
    expires_at = str(today.replace(year = today.year + 1))

    data = {}
    data['name'] = token_name
    data['expires_at'] = expires_at
    data['scopes'] = ['api', 'read_user', 'read_api', 'read_repository', 'write_repository']

    json_data = json.dumps(data)
    
    headers = {"PRIVATE-TOKEN":"{}".format(api_token)}
    
    r = requests.post(
        url=url,
        data=json_data,
        headers=headers
        )
    
    print(r)

    if r.status_code == 201:
        print("Personal token successfully created!")

    else:
        print("Couldn't create new personal token")
        print (r.status_code, r.reason)