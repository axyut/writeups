<?php
function xor_encrypt($in)
{
    $urlDecodedCookie = "MGw7JCQ5OC04PT8jOSpqdmkgJ25nbCorKCEkIzlscm5oKC4qLSgubjY=";
    $key = base64_decode($urlDecodedCookie);
    //echo "\n\nKEY\n" . $key;
    $text = $in;
    $outText = '';

    // Iterate through each character
    for ($i = 0; $i < strlen($text); $i++) {
        $outText .= $text[$i] ^ $key[$i % strlen($key)];
    }

    return $outText;
}
$payload = array("showpassword" => "no", "bgcolor" => "#ffffff");
$key = xor_encrypt((json_encode($payload)));
echo "\nXOR KEY\n" . $key;
?>