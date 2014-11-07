<?php
define('PHURL', true);
ini_set('display_errors', 0);
$prefix[0] = '';
$filename = 'install';
if (is_dir($filename)) {
    die ("To get Phurl up and running, you first need to go through the <a href=\"install\">installation wizard</a> which will help you set up your new URL shortener in a matter of moments.<br/><br/>If you've already installed Phurl, then you MUST delete the install directory before it will function.");
}
?>
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

if (count($_GET) > 0) {
    $url   = mysql_real_escape_string(trim($_GET['url']));
    $alias = mysql_real_escape_string(trim($_GET['alias']));

    if (!preg_match("/^(".URL_PROTOCOLS.")\:\/\//i", $url)) {
 	$prefix = explode(":", $url);
 	if ($prefix[0] == 'mailto') {
 		$url = $url;
 	} else {
        $url = "http://".$url;
 	}
    }

    $last = $url[strlen($url) - 1];

    if ($last == "/") {
        $url = substr($url, 0, -1);
    }

    $data = @parse_url($url);
		if ($prefix[0] == 'mailto') {
			$data['scheme'] = 'mailto';
			$data['host'] = 'none';
		}
    if (strlen($url) == 0) {
        $_ERROR[] = "Please enter a URL to shorten.";
    }
    else if (empty($data['scheme']) || empty($data['host'])) {
        $_ERROR[] = "Please enter a valid URL to shorten.";
    }
    else {
        $hostname = get_hostname();
        $domain   = get_domain();

        if (preg_match("/($hostname)/i", $data['host'])) {
            $_ERROR[] = "The URL you have entered is not allowed.";
        }
    }

    if (strlen($alias) > 0) {
        if (!preg_match("/^[a-zA-Z0-9_-]+$/", $alias)) {
            $_ERROR[] = "Custom aliases may only contain letters, numbers, underscores and dashes.";
        }
        else if (code_exists($alias) || alias_exists($alias)) {
            $_ERROR[] = "The custom alias you entered already exists.";
        }
        else if (preg_match("/NCN/", $alias)) { //(preg_match("/[Nn][Cc][Nn][_-]?[a-zA-Z0-9]+/", $alias)) {
            $_ERROR[] = "There can be only one flag in the database, stop trying to troll the other teams ;)";
        }
    }

    if (count($_ERROR) == 0) {
        $create = true;

        if (($url_data = url_exists($url))) {
            $create    = false;
            $id        = $url_data[0];
            $code      = $url_data[1];
            $old_alias = $url_data[2];

            if (strlen($alias) > 0) {
                if ($old_alias != $alias) {
                    $create = true;
                }
            }
        }

        if ($create) {
            do {
                $code = generate_code(get_last_number());

                if (!increase_last_number()) {
                    die("System error!");
                }

                if (code_exists($code) || alias_exists($code)) {
                    continue;
                }

                break;
            } while (1);

            $id = insert_url($url, $code, $alias);
        }

        if (strlen($alias) > 0) {
            $code = $alias;
        }

        $short_url = SITE_URL."/".$code;

        $_GET['url']   = "";
        $_GET['alias'] = "";

        require_once("html/header.php");
        require_once("html/index_done.php");
        require_once("html/index_form.php");
        require_once("html/footer.php");
        exit();
    }
}

require_once("html/header.php");
require_once("html/index_form.php");
require_once("html/footer.php");
