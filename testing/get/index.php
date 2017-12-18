<?php
if(!isset($_GET["id"])){
    echo("need ?id=something");
    return;
}
$id = $_GET["id"];
echo("id is \"".$id."\"");
?>