<?php
function encrypt($in)
{
  $key = "KNHL";
  $text = $in;
  $outText = '';

  // Iterate through each character
  for ($i = 0; $i < strlen($text); $i++) {
    $outText .= $text[$i] ^ $key[$i % strlen($key)];
  }

  return $outText;
}

$payload = array("showpassword" => "yes", "bgcolor" => "#ffffff");
echo "\nPassword\n" . base64_encode(encrypt((json_encode($payload))));

?>