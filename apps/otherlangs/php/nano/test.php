<?php

require_once 'vendor/autoload.php';

use laylatichy\Nano;

//default settings for nano app - http://0.0.0.0:3000
$nano = new Nano();

// GET route
$nano->get(
    '/',
    function ($request) use ($nano) {
        $nano->response->code(200);
        $nano->response->plain('hello plain');
    }
);

// POST route
$nano->post(
    '/hello/{name}',
    function ($request, $name) use ($nano) {
        $nano->response->code(200);
        $nano->response->json(
            [
                'hello' => $name,
            ]
        );
    }
);

// Start nano
$nano->start();
