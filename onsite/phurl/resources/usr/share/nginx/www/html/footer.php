<?php
if( !defined('PHURL' ) ) {
    header('HTTP/1.0 404 Not Found');
    exit();
}
ini_set('display_errors', 0);
?>
<h2>Recent shortened URLs</h2>
<div id="recent"></div>
<h2>Browser Bookmarklets</h2>
<div id="bookmarklets">
<p>Drag these links to your browser toolbar.</p>
<p><a href="javascript:(function(){var%20ali=prompt('Enter%20a%20custom%20alias:');if(ali){location.href='<?php echo SITE_URL ?>/index.php?url='%20+%20escape(location.href)%20+%20'&alias='+ali;}})();" title="Shorten with a custom alias">Shorten with a custom alias</a><br />
<a href="javascript:void(location.href='<?php echo SITE_URL ?>/index.php?alias=&url='+escape(location.href))">Shorten without a custom alias</a></p>
</div>
</div>
<div id="footer">
<p>&copy; <?php echo date("Y"); ?> <?php echo SITE_TITLE ?> - Powered by <a href="http://code.google.com/p/phurl">Phurl <?php echo PHURL_VERSION ?></a></p>
</div>
</div>
<script>
$.ajax({
    type: "GET",
    url: "html/recent.php",
    data: 'n=' + (Math.round(new Date().getTime() / 1000) - 7200),
    success: function(msg){
        $('#recent').html(msg);
    }
});
</script>
</body>
</html>
