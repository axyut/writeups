import requests
import binascii
import urllib
import base64
import string

charset = string.ascii_lowercase

url = "http://natas28.natas.labs.overthewire.org/index.php"
s = requests.Session()
s.auth = ('natas28', 'skrwxciAe6Dnb0VfFDzDEHcCzQmv3Gd4')

sample = "aaaaaaaaa"

for x in charset:
    data = {'query':sample+x}
    r = s.post(url, data=data)
    cipher = r.url.split('=')[1]
    cipher = urllib.parse.unquote(cipher)
    print("[*] last char. = %s | %s" % (x, cipher))
