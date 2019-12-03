import requests
import json

BASE_URL = "http://localhost:8080/api/"
JWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkZlbGlwZSBNZWdhbGUiLCJpYXQiOjE1MTYyMzkwMjJ9.vfQESuSJEOB9HJNtXGCYG-kvV9oTn7y5Oxooh4c9ZUI"
HEADERS = {'Authorization': 'Bearer {0}'.format(JWT), 'Content-Type': 'application/json'}
command = ""

while command != "exit":
    command = input("~ ")
    collection = ""
    operation = ""
    data = {}
    tempString = ""
    if command == "help" or command == "\\?":
        print('''Insert/Search/Delete''')
        print( '''<collection name>.<command>(<data>)''')
        print( '''(e.g. user.insert({"name":"john doe"}) )''')
        print( '''(e.g. user.Search({"name":"john doe"}) )''')
        print( '''(e.g. user.Delete({"name":"john doe"}) )''')
        print('''update''')
        print( '''<collection name>.<command>(<searchData>,<newData>)''')
        print( '''(e.g. user.update({"search": {"name": "john doe"}, "data": {"name": "Joao Castro"} }))''')
    elif command != "exit": 
        for char in command:
            if collection == "":
                if char != '.':
                    tempString += char
                else:
                    collection = tempString
                    tempString = ""
            elif operation == "":
                if char != '(':
                    tempString += char
                else:
                    operation = tempString
                    tempString = ""
            else:
                if char != ")":
                    tempString += char
                else:
                    data = tempString
                    tempString = ""
        try:
            requestBody = json.loads(data)
            if operation == "insert":
                url_path = "{0}/_create".format(collection)
            elif operation == "update":
                url_path = "{0}/_update".format(collection)
            elif operation == "search":
                url_path = "{0}/_search".format(collection)
            elif operation == "remove":
                url_path = "{0}/_delete".format(collection)
            else:
                raise ValueError("")
            r = requests.post(BASE_URL+url_path, data=data, headers=HEADERS)
            print(r.text)
        except ValueError:
            print("unknown op")
            continue