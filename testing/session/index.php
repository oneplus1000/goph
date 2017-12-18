<?php
session_start();
$_SESSION["name"] = "oneplus";
echo("set session done<br />");
echo("<a href=\"./session.php\">View session value</a>")
?>