<?php

use Workerman\Worker;

require_once __DIR__ . '/vendor/autoload.php';


$GLOBALS["TEST"]= str_repeat("test" , 5000);


$myjson = '{"content":[
	{
		"id": 272732,
		"tag": "",
		"label": [
			"meta_aaa"
		],
		"short_text": "smart title",
		"text": "это текст-\"рыба\", часто используемый в печати и вэб-дизайне. Lorem Ipsum является стандартной \"рыбой\" для текстов на латинице с начала XVI века. В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, используя Lorem Ipsum для распечатки образцов. Lorem Ipsum не только успешно пережил без заметных изменений пять веков, но и перешагнул в электронный дизайн. Его популяризации в новое время послужили публикация листов Letraset с образцами Lorem Ipsum в 60-х годах и, в более недавнее время, программы электронной вёрстки типа Aldus PageMaker, в шаблонах которых используется Lorem Ipsum.",
		"id_page": 1312312,
		"date_created": "2022-05-05",
		"img": null,
		"id_page_group": [1,2,3,4444]
	},
	{
		"id": 272732,
		"tag": "",
		"label": [
			"meta_aaa"
		],
		"short_text": "smart title",
		"text": "это текст-\"рыба\", часто используемый в печати и вэб-дизайне. Lorem Ipsum является стандартной \"рыбой\" для текстов на латинице с начала XVI века. В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, используя Lorem Ipsum для распечатки образцов. Lorem Ipsum не только успешно пережил без заметных изменений пять веков, но и перешагнул в электронный дизайн. Его популяризации в новое время послужили публикация листов Letraset с образцами Lorem Ipsum в 60-х годах и, в более недавнее время, программы электронной вёрстки типа Aldus PageMaker, в шаблонах которых используется Lorem Ipsum.",
		"id_page": 1312312,
		"date_created": "2022-05-05",
		"img": null,
		"id_page_group": [1,2,3,4444]
	},
	{
		"id": 272732,
		"tag": "",
		"label": [
			"meta_aaa"
		],
		"short_text": "smart title",
		"text": "это текст-\"рыба\", часто используемый в печати и вэб-дизайне. Lorem Ipsum является стандартной \"рыбой\" для текстов на латинице с начала XVI века. В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, используя Lorem Ipsum для распечатки образцов. Lorem Ipsum не только успешно пережил без заметных изменений пять веков, но и перешагнул в электронный дизайн. Его популяризации в новое время послужили публикация листов Letraset с образцами Lorem Ipsum в 60-х годах и, в более недавнее время, программы электронной вёрстки типа Aldus PageMaker, в шаблонах которых используется Lorem Ipsum.",
		"id_page": 1312312,
		"date_created": "2022-05-05",
		"img": null,
		"id_page_group": [1,2,3,4444]
	},
	{
		"id": 272732,
		"tag": "",
		"label": [
			"meta_aaa"
		],
		"short_text": "smart title",
		"text": "это текст-\"рыба\", часто используемый в печати и вэб-дизайне. Lorem Ipsum является стандартной \"рыбой\" для текстов на латинице с начала XVI века. В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, используя Lorem Ipsum для распечатки образцов. Lorem Ipsum не только успешно пережил без заметных изменений пять веков, но и перешагнул в электронный дизайн. Его популяризации в новое время послужили публикация листов Letraset с образцами Lorem Ipsum в 60-х годах и, в более недавнее время, программы электронной вёрстки типа Aldus PageMaker, в шаблонах которых используется Lorem Ipsum.",
		"id_page": 1312312,
		"date_created": "2022-05-05",
		"img": null,
		"id_page_group": [1,2,3,4444]
	},
	{
		"id": 272732,
		"tag": "",
		"label": [
			"meta_aaa"
		],
		"short_text": "smart title",
		"text": "это текст-\"рыба\", часто используемый в печати и вэб-дизайне. Lorem Ipsum является стандартной \"рыбой\" для текстов на латинице с начала XVI века. В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, используя Lorem Ipsum для распечатки образцов. Lorem Ipsum не только успешно пережил без заметных изменений пять веков, но и перешагнул в электронный дизайн. Его популяризации в новое время послужили публикация листов Letraset с образцами Lorem Ipsum в 60-х годах и, в более недавнее время, программы электронной вёрстки типа Aldus PageMaker, в шаблонах которых используется Lorem Ipsum.",
		"id_page": 1312312,
		"date_created": "2022-05-01",
		"img": null,
		"id_page_group": [1,2,3,4444]
	}
	]}';

$GLOBALS["TEST_OBJ"] =   json_decode($myjson, true);
var_dump($GLOBALS["TEST_OBJ"]['content'][0]['text']);



// #### http worker ####
$http_worker = new Worker('http://0.0.0.0:2345');

// 4 processes
$http_worker->count = 4;

// Emitted when data received
$http_worker->onMessage = function ($connection, $request) {
    $_GET = $request->get();
    $path = $request->path();
    //var_dump($request);
    //$request->get();
    //$request->post();
    //$request->header();
    //$request->cookie();
    //$request->session();
    //$request->uri();
    //$request->path();
    //$request->method();

    // 404 - 130k

    // Send data to client
    // 60 - 63k
    if(substr($path,0,5) == '/test'){
        $connection->send($GLOBALS["TEST"]);
    // 59-70k
    }else if(substr($path,0,10) == '/json_text'){
        $j=json_encode($GLOBALS["TEST"],JSON_UNESCAPED_UNICODE);
        $connection->send($j);
    // 52k - 65k
    }else if(substr($path,0,5) == '/json'){
        $j=json_encode($GLOBALS["TEST_OBJ"],JSON_UNESCAPED_UNICODE);
        $connection->send($j);
    }else{ //if ($path=="/"){
        $connection->send("Hello World");
    }
};

// Run all workers
Worker::runAll();