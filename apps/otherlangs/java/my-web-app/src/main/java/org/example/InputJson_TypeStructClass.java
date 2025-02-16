
package org.example;



public class InputJson_TypeStructClass {

    public int id;
    public String tag;
    public String[] label;
    public String short_text;
    public String text;  
    public int id_page;
    public String date_created;
    public String img;
    public int[] id_page_group;                

    //getters and setters
}

// type Json_struct struct {
// 	Content []struct {
// 		ID          int         `json:"id"`
// 		Tag         string      `json:"tag"`
// 		Label       []string    `json:"label"`
// 		ShortText   string      `json:"short_text"`
// 		Text        string      `json:"text"`
// 		IDPage      int         `json:"id_page"`
// 		DateCreated string      `json:"date_created"`
// 		Img         interface{} `json:"img"`
// 		IDPageGroup []int       `json:"id_page_group"`
// 	} `json:"content"`
// }

// https://mkyong.com/java/how-do-convert-java-object-to-from-json-format-gson-api/

// public class Staff {

//     private String name;
//     private int age;
//     private String[] position;              // array
//     private List<String> skills;            // list
//     private Map<String, BigDecimal> salary; // map

//     //getters and setters
// }
