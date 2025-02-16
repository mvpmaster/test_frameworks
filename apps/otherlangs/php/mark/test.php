<?php
use Mark\App;

require 'vendor/autoload.php';

$api = new App('http://0.0.0.0:3000');

$api->count = 4; // process count

$GLOBALS["TEST"]= str_repeat("test" , 5000);

$api->any('/', function ($requst) {
    return 'Hello world';
});

$api->any('/test', function ($requst) {
    return $GLOBALS["TEST"];
});

$api->get('/hello/{name}', function ($requst, $name) {
    return $GLOBALS["TEST"];
});

$api->post('/user/create', function ($requst) {
    return json_encode(['code'=>0 ,'message' => 'ok']);
});

$api->start();