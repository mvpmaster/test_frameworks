package handlers

import (
	"github.com/gofiber/fiber/v2"
	//"superapp/internal/service/redirect_url"
	"strings"
	"encoding/json"
	//"github.com/mailru/easyjson"
	j "github.com/goccy/go-json"
	jay "github.com/francoispqt/gojay"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	segmentio "github.com/segmentio/encoding/json"
)

var test=strings.Repeat("test", 5000)
var test2 []byte //[]byte
var test3 string

//easyjson:json
type Json_struct struct {
	Content []struct {
		ID          int         `json:"id"`
		Tag         string      `json:"tag"`
		Label       []string    `json:"label"`
		ShortText   string      `json:"short_text"`
		Text        string      `json:"text"`
		IDPage      int         `json:"id_page"`
		DateCreated string      `json:"date_created"`
		Img         interface{} `json:"img"`
		IDPageGroup []int       `json:"id_page_group"`
	} `json:"content"`
}

var myjson = `{"content":[
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
	]}`

var json_test Json_struct

func Init_Test(app *fiber.App) { //go service.SomeWorker(); 
	json.Unmarshal([]byte(myjson), &json_test)
	test2,_=json.Marshal(map[string]string{"test":test})
	test3=string(test2)
	//go redirect_url.WorkerBlanaceForgotten() // start worker thread

	// 404 = 80k

	// 29-35 K
    app.Get("/test", Go_Test) 
	// 38-39 k
	app.Get("/fiber_json", FiberJson) 
	// 30k - 33k
	app.Get("/fiber_json_text", FiberJsonText) 
	// 43k - 47k (52-53k)
	app.Get("/json", SimpleJson) // best
	// 29k - 34k
	app.Get("/json_text", SimpleJson_text) 

	app.Get("/json_gojson", SimpleJson_MailRu) // bad

	app.Get("/json_gojay", SimpleJson_Gojay) // not valid
	
	// 51-55k
	app.Get("/json_iter", SimpleJson_Jsoniter) // +/- bad

	app.Get("/json_segmentio", SimpleJson_Segmentio) // ok

	

	
	} // some events

func Go_Test(c *fiber.Ctx) error {

	//url := c.Query("video") ; // log.Println(url)
	//url, err := redirect_url.RootRedirect(url) ; if err != nil { return c.Status(404).SendString(err.Error()) } // log.Println(url)
	//c.Status(201).JSON( map[string]string{"test": test, }) // 40000
	// медленнее
	// test2,_=json.Marshal(map[string]string{"test":test})
	// c.SendString(string(test2))
	c.SendString(test) // https://github.com/gofiber/fiber/issues/809 // 51000
	
	//c.SendString(test3, 200) 
	; return nil 	
	} // перенаправление 


// https://github.com/gofiber/fiber/issues/809
// c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

func FiberJsonText(c *fiber.Ctx) error {

		//a, _:= json.Marshal(json_test)
		//url := c.Query("video") ; // log.Println(url)
		//url, err := redirect_url.RootRedirect(url) ; if err != nil { return c.Status(404).SendString(err.Error()) } // log.Println(url)
		c.Status(201).JSON( test3) // 40000
		// медленнее
		// test2,_=json.Marshal(map[string]string{"test":test})
		// c.SendString(string(test2))
		//c.SendString(test) // https://github.com/gofiber/fiber/issues/809 // 51000
		
		//c.SendString(test3, 200) 
		; return nil 	
				} // перенаправление 

func FiberJson(c *fiber.Ctx) error {

//a, _:= json.Marshal(json_test)
//url := c.Query("video") ; // log.Println(url)
//url, err := redirect_url.RootRedirect(url) ; if err != nil { return c.Status(404).SendString(err.Error()) } // log.Println(url)
c.Status(201).JSON( &json_test) // 40000
// медленнее
// test2,_=json.Marshal(map[string]string{"test":test})
// c.SendString(string(test2))
//c.SendString(test) // https://github.com/gofiber/fiber/issues/809 // 51000

//c.SendString(test3, 200) 
; return nil 	
		} // перенаправление 


func SimpleJson(c *fiber.Ctx) error {

	a, _:= json.Marshal(&json_test)
	//url := c.Query("video") ; // log.Println(url)
	//url, err := redirect_url.RootRedirect(url) ; if err != nil { return c.Status(404).SendString(err.Error()) } // log.Println(url)
	//c.Status(201).JSON( json_test) // 40000
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	c.Status(201).SendString(string(a))
	// медленнее
	// test2,_=json.Marshal(map[string]string{"test":test})
	// c.SendString(string(test2))
	//c.SendString(test) // https://github.com/gofiber/fiber/issues/809 // 51000
	
	//c.SendString(test3, 200) 
	; return nil 	
			} // перенаправление 

// benchmarks
//	https://github.com/goccy/go-json
func SimpleJson_MailRu(c *fiber.Ctx) error {

	// https://github.com/mailru/easyjson
	a, _:= j.Marshal(json_test)

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	c.Status(201).SendString(string(a))

	return nil

}

func SimpleJson_Gojay(c *fiber.Ctx) error {

	// https://github.com/francoispqt/gojay
	a, err:= jay.Marshal(json_test)

	if err != nil { fmt.Println(err) }

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	c.Status(201).SendString(string(a))

	return nil

}

func SimpleJson_Jsoniter(c *fiber.Ctx) error {

var json = jsoniter.ConfigCompatibleWithStandardLibrary
a, err:=json.Marshal(&json_test)

if err != nil { fmt.Println(err) }

c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
c.Status(201).SendString(string(a))

return nil

}

func SimpleJson_Segmentio(c *fiber.Ctx) error {

	
	a, err:=segmentio.Marshal(&json_test)
	
	if err != nil { fmt.Println(err) }
	
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	c.Status(201).SendString(string(a))
	
	return nil
	
	}

func SimpleJson_text(c *fiber.Ctx) error {

	a, _:= json.Marshal(test3)
	//url := c.Query("video") ; // log.Println(url)
	//url, err := redirect_url.RootRedirect(url) ; if err != nil { return c.Status(404).SendString(err.Error()) } // log.Println(url)
	//c.Status(201).JSON( json_test) // 40000
	c.SendString(string(a))
	// медленнее
	// test2,_=json.Marshal(map[string]string{"test":test})
	// c.SendString(string(test2))
	//c.SendString(test) // https://github.com/gofiber/fiber/issues/809 // 51000
	
	//c.SendString(test3, 200) 
	; return nil 	
			} 

	// return c.Status(fiber.StatusOK).JSON( &url )
// interface
// func SetSettings(cdn_host string, percent float64, history_balance float64){
// 	redirect_url.SetSettings(cdn_host, percent, history_balance) }