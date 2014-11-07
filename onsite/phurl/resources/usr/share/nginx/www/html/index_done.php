<?php
if( !defined('PHURL' ) ) {
    header('HTTP/1.0 404 Not Found');
    exit();
}
ini_set('display_errors', 0);
?>
<h2>Short URL has been created!</h2>
<p>Long URL: <strong><?php echo htmlentities($url) ?></strong> (<?php echo strlen($url) ?> characters)</p>
<p>Short URL: <strong><a href="<?php echo htmlentities($short_url) ?>" target="_blank"><?php echo htmlentities($short_url) ?></a></strong> (<?php echo strlen($short_url) ?> characters)<br />

