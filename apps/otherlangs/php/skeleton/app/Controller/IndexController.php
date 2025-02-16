<?php

declare(strict_types=1);
/**
 * This file is part of Simps.
 *
 * @link     https://simps.io
 * @document https://doc.simps.io
 * @license  https://github.com/simple-swoole/simps/blob/master/LICENSE
 */
namespace App\Controller;

$GLOBALS["TEST"]= str_repeat("test" , 5000);

class IndexController
{
    public function index($request, $response)
    {
        $response->end(
            $GLOBALS["TEST"]
            //json_encode(
            //    [
            //        'method' => $request->server['request_method'],
            //        'message' => $GLOBAL["TEST"], //'Hello Simps.',
            //    ]
            //)
        );
    }

    public function hello($request, $response, $data)
    {
        $name = $data['name'] ?? 'Simps';

        $response->end(
            json_encode(
                [
                    'method' => $request->server['request_method'],
                    'message' => "Hello {$name}.",
                ]
            )
        );
    }
}
