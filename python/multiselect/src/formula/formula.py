#!/usr/bin/python3

def run(multi):
    print("Hello World!")
    items = multi.split("|")
    print("items:", items)
    
    loop = 0
    for item in items:
        loop = loop +1
        print("Item {}:".format(loop), item)
