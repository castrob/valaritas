import requests
import json

BASE_URL = "http://localhost:8080/api/"
JWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkZlbGlwZSBNZWdhbGUiLCJpYXQiOjE1MTYyMzkwMjJ9.vfQESuSJEOB9HJNtXGCYG-kvV9oTn7y5Oxooh4c9ZUI"
HEADERS = {'Authorization': 'Bearer {0}'.format(JWT)}
command = ""

while command != "exit":
    command = input("~ ")
    cmd_arr = command.split(" ")
    operation = cmd_arr[0]

    if operation == "create":
        coll_name = cmd_arr[1]
        coll_vals = cmd_arr[2]
        url_path = "{0}/_create".format(coll_name)
        r = requests.post(BASE_URL+url_path, data=coll_vals, headers=HEADERS)
        print(r.text)
    elif operation == "search":
        coll_name = cmd_arr[1]
        # coll_vals = cmd_arr[2]
        url_path = "{0}/_search".format(coll_name)
        r = requests.get(BASE_URL+url_path, headers=HEADERS)
        print(r.text)
    elif operation == "update":
        print(cmd_arr)
    elif operation == "delete":
        print(cmd_arr)
    elif operation == "help":
        print(
            '''create <collection name> <collection values> (e.g. create user {"name":"john doe"})''')
        print('''search <collection name>''')
        print(
            '''update <collection name> <collection values> (e.g. update user {"name":"johnny cash"})''')
        print('''delete <collection name>''')
    elif operation == "exit":
        break
    else:
        print("unknown op")
