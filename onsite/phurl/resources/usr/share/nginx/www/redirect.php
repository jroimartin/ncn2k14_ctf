<?php
require_once("config.php");
require_once("functions.php");

$last_request = @$_SESSION['last_request'];
$current_time = time();
@$_SESSION['last_request'] = $current_time;
if (!is_null($last_request) && ($current_time - $last_request) < 1) {
  header("HTTP/1.1 429 Too Many Requests");
  exit();
}

db_connect();

$alias = trim(mysql_real_escape_string($_GET['alias']));

if (!preg_match("/^[a-zA-Z0-9_-]+$/", $alias)) {
  header("Location: ".SITE_URL, true, 301);
  exit();
}

if (($url = get_url($alias))) {
    header("Location: $url", true, 301);
    exit();
}

header("Location: ".SITE_URL, true, 301);
