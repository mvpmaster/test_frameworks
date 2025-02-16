from sanic.response import text, json, redirect, raw

big_text='test'*5000

import json as def_json
import ujson
import orjson, nujson


ensure=False

# ensure=False
# orjson     = 28800-38000
# ujson      = 27700-35800
# nujson     = 27500-35400
# sanic_json = 29000-35000
# json       = 25000-30300

# dumps orjson > json > nujson > ujson  > sanic_json # но в целом почти разницы нет
# 33519 - 31298

# changed cpu settings
# orjson text   31100-38400 (видимо текст эффективнее, при этом json короче из-за utf-8, а byte не сильно влияет)+ стабильнее из-за аллокации переменных ? и выделения памяти?
# cached j raw  34000-36000
# cached json   30000-34500
# cached j raw-r29000-34600
# sanic_json    32000-33400
# ujson         29800-33822
# nujson        30000-33500
# json          27500-29000

# py 3.11
# obj dumps
# orjson raw    28800-29500
# orjson text   26600-35303 (ответ меньше весит в раза), уже utf-8, и выдается в bytes (а не string)
# cached j raw-r27000-31000
# cached j raw  25000-31000
# json   		26000-28140
# nujson 		26300-27300
# ujson  		25400-27800
# cached resp   24000-31000
# sanic_json 	24361-26703
# cached json   23600
# sanic j resp 	7500

# text 5kb dumps

# nujson        26000-32200
# orjson		25500-31600 (ответ меньше весит в раза), уже utf-8, и выдается в bytes (а не string)
# sanic_json    26500-31500
# json          25500-26600
# ujson         23600-26800

#https://github.com/ultrajson/ultrajson


obj={"content":[
			{
				"id": 272732,
				"tag": "",
				"label": [
					"meta_aaa"
				],
				"short_text": "smart title",
				"text": "это текст-\"рыба\", часто используемый в печати и вэб-дизайне. Lorem Ipsum является стандартной \"рыбой\" для текстов на латинице с начала XVI века. В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, используя Lorem Ipsum для распечатки образцов. Lorem Ipsum не только успешно пережил без заметных изменений пять веков, но и перешагнул в электронный дизайн. Его популяризации в новое время послужили публикация листов Letraset с образцами Lorem Ipsum в 60-х годах и, в более недавнее время, программы электронной вёрстки типа Aldus PageMaker, в шаблонах которых используется Lorem Ipsum.",
				"id_page": 1312312,
				"date_created": "2022-05-05T19:49:55.084894Z",
				"img": None,
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
				"date_created": "2022-05-05T19:49:55.084894Z",
				"img": None,
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
				"date_created": "2022-05-05T19:49:55.084894Z",
				"img": None,
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
				"date_created": "2022-05-05T19:49:55.084894Z",
				"img": None,
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
				"date_created": "2022-05-05T19:49:55.084894Z",
				"img": None,
				"id_page_group": [1,2,3,4444]
			},
			]}

cached_json=orjson.dumps(obj)
cached_response=text(cached_json.decode("utf-8"))
cached_response_byte=raw(cached_json)
cached_json=cached_json.decode("utf-8")

#orjson 4048
#ujson 14923
#nujson 14923
#json 15028


print('orjson',len(cached_json))
print('ujson',len(ujson.dumps(obj)))
print('nujson',len(nujson.dumps(obj)))
print('json',len(def_json.dumps(obj)))

print('orjson',len(orjson.dumps(obj).decode("utf-8")))
print('ujson -',len(ujson.dumps(obj, ensure_ascii=False)))
print('nujson -',len(nujson.dumps(obj, ensure_ascii=False)))
print('json -',len(def_json.dumps(obj, ensure_ascii=False)))



async def root_redirect(request): 
    try: test = request.args['test'][0] # валидация
    except: return text('not found',status=404)

    return text(big_text, status=200)
    #return redirect(url, status=301) # редирект на основной сервер


class InitRedirect:
    def __init__(self, app=None, interfaces=None): 
        # global AnyModule; AnyModule = interfaces.AnyModule

        # 404 = 24000
        # 404 = 11000-12000

		# swarm 3 = 28000 - 31000
        # = 14000-18100
        app.add_route( root_redirect, f"/test" , 
                       methods=["HEAD", "GET"] )
        
        # swarm 3 = 21395 - 25000
        # 10710
        @app.get("/api/sanic_json_text")
        async def app_text_json(request): 
            return json(big_text, ensure_ascii=ensure) 

        @app.get("/api/sanic_json_text_response")
        async def app_text_json(request): 
            return big_text
        
        # swarm 3 = 21395 - 28000
        # 10700
        @app.get("/api/sanic_json")
        async def app_test_json(request): 
            return json(obj, ensure_ascii=ensure)     
        
        @app.get("/api/sanic_json_response")
        async def app_test_json(request): 
            return obj      

		# swarm 3 = 21197 - 27000
        # 11300
        @app.get("/api/json")
        async def app_nativ_json(request): 
            return text(def_json.dumps(obj, ensure_ascii=ensure))   
        
        @app.get("/api/json_text")
        async def app_nativ_json(request): 
            return text(def_json.dumps(big_text, ensure_ascii=ensure))   
        
        
        # swarm 3 = 23315-30000 (28142)
        # 10800
        @app.get("/api/ujson")
        async def app_nativ_ujson(request): 
            return text(ujson.dumps(obj, ensure_ascii=ensure))   

        @app.get("/api/ujson_text")
        async def app_nativ_ujson(request): 
            return text(ujson.dumps(big_text, ensure_ascii=ensure))   

        @app.get("/api/orjson")
        async def app_nativ_orjson(request): 
            return text(orjson.dumps(obj).decode("utf-8") )

        @app.get("/api/orjson_raw")
        async def app_nativ_orjson(request): 
            return raw(orjson.dumps(obj) )  
        
        
        # 33k - 34k vs 29kk (json), но почти нет разницы
        @app.get("/api/orjson_text")
        async def app_nativ_orjson(request): 
            return text(orjson.dumps(big_text, ensure_ascii=ensure).decode("utf-8") )   
        
        @app.get("/api/nujson")
        async def app_nativ_orjson(request): 
            return text(nujson.dumps(obj, ensure_ascii=ensure))  

        @app.get("/api/nujson_text")
        async def app_nativ_orjson(request): 
            return text(nujson.dumps(big_text, ensure_ascii=ensure))  

        @app.get("/api/cached_json")
        async def app_nativ_orjson(request): 
            return text(cached_json)  
        
        @app.get("/api/cached_json_raw")
        async def app_nativ_orjson(request): 
            return raw(cached_json)  
        @app.get("/api/cached_json_raw_response")
        async def app_nativ_orjson(request): 
            return cached_response_byte 
        
        
        @app.get("/api/cached_json_response")
        async def app_nativ_orjson(request): 
            return cached_response 
        
        
        
        


        # Для сверки RPS
        # @app.get("/api/rps_test_json")
        # async def app_test_json(request): 
        #     return json("test"*5000) 

        # @app.get("/api/rps_test_redirect")
        # async def app_redirect(request):
        #     return redirect("http://", status=301)

        #app.add_route(root, f"/video/<video_id>/<m3u8_file>", methods=["HEAD", "GET"])
