<?php
if( !defined('PHURL' ) ) {
    header('HTTP/1.0 404 Not Found');
    exit();
}
ini_set('display_errors', 0);
?>
<html>
<head>
<title><?php echo SITE_TITLE ?></title>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<link rel="stylesheet" type="text/css" href="assets/default.css" />
<script src="assets/jquery-1.11.1.min.js"></script>
</head>
<body>
<div id="container">
<div id="header">
<a href="<?php echo SITE_URL ?>" ><h1 style="color:black;"><?php echo SITE_TITLE ?></h1></a>
</div>
<div id="content">
