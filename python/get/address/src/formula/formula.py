#!/usr/bin/python3
from urllib.request import urlopen
import json

def Run(cep):
    cep = cep.replace("-", "")
    url = 'http://viacep.com.br/ws/%s/json/' % cep
    response = urlopen(url)
    address_json = json.loads(response.read().decode('utf-8'))
    if "erro" in address_json:
        print("Invalid CEP")
    else:
        print(address_json)
