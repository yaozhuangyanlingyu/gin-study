<?php
function p($arr){
	echo '<pre>';
	print_r($arr);
	die;
}

function http_request($url, $data = null)
{
    $curl = curl_init();
    curl_setopt($curl, CURLOPT_URL, $url);
    curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, FALSE);
    curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, FALSE);
    if (!empty($data)){
        curl_setopt($curl, CURLOPT_POST, 1);
        curl_setopt($curl, CURLOPT_POSTFIELDS, $data);
    }
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, TRUE);
    $output = curl_exec($curl);
    curl_close($curl);
    return $output;
}

$tmp = file_get_contents("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=wwde96ee9bc90ebb43&corpsecret=h3MleGzWRq7j6SUkHWMm1WJrE7cR5yscGxLQunSLKfk");
$tmp = json_decode($tmp, true);
$accessToken = '';
if( $tmp['errcode'] == 0 ){
    $accessToken = $tmp['access_token'];
}

$url = 'https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=';
$res = http_request(sprintf("%s%s", $url, $accessToken), json_encode([
    'touser'    => '@all',
    'msgtype'   => 'text',
    'agentid'   => '1000002',
    'text'      => [
        'content'   => '下辈子娶高圆圆当老婆',
    ],
    'safe'      => 0,
]));
var_dump( $res );





