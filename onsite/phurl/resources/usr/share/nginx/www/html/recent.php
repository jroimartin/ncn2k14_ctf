<?php
require_once("../config.php");
require_once("../functions.php");

db_connect();

$num = $_GET['n'];
if (is_null($num) || $num === "") {
    header('HTTP/1.0 404 Not Found');
    exit();
}
$query = "SELECT code, alias, url FROM ".DB_PREFIX."urls WHERE date_added > FROM_UNIXTIME($num) ORDER BY id DESC LIMIT 5";
//echo $query;
$db_result = mysql_query($query) or db_die(__FILE__, __LINE__, mysql_error());
$row_count = mysql_num_rows($db_result);
if ($row_count > 1000) {
    db_die(__FILE__, __LINE__, "Dude, seriously? Are you trying to dump the whole database? C'mon, you can make a smarter query!");
}
echo "<ul>\n";
while ($row_count > 0) {
    $row_count = $row_count - 1;
    $db_row = mysql_fetch_row($db_result);
    $url = SITE_URL."/".$db_row[0];
    if (!is_null($db_row[1]) && $db_row[1] != "") {
        $url = SITE_URL."/".$db_row[1];
    }
    $url = htmlentities($url);
    $target = htmlentities($db_row[2]);
    echo "<li><a href=$url>$target</a></li>\n";
}
echo "</ul>\n";
?>
