//use serde::{Serialize, Deserialize};
use warp::{self, http::Uri, path::FullPath, Filter,  Reply , http::Response};

use lazy_static::lazy_static;
use std::fs;
use serde::{Deserialize, Serialize};
use serde_json::Result;

// extern crate json;
// use json::JsonValue;


// use rand::Rng;
//use std::collections::HashMap;
//use core::hash::Hash;
// use std::str::FromStr;

// #[derive(Deserialize)]
// struct InputVideoUrl {
//     video: String,
// }

// mod redirect_url;

// struct Err404 {
//     text: String
// }

// fn resp(video: &String) -> impl Reply {//&HashMap<String, String>) -> impl Reply {
//     //if query.contains_key("video") {
//         //let video = query.get("1" )//&"video".to_string() ); // println!("{}",value);
//     //     println!("{}", q.video.as_str());
//     //format!("path: {}\nquery: {:?}", path, query)
//     //     //"test"
//         //let res = redirect_url::ab_redirect_by_rand(video, cdn_host, &rng.gen::<f64>());
//         //println!("{}",&video);
//         //format!("{}", &video)
//         // match query.get(&"vidoe".to_string()) {
//         //     Some(&url) => warp::redirect(Uri::from_static("url")) 
//         //     _ => warp::redirect(Uri::from_static("/error"))
//         // }
//         warp::redirect(Uri::from_static("url"))
//         // match video{
//         //     Some(url) =>{
//         //         warp::redirect(Uri::from_static(url))
//         //     }
//         //     None => warp::redirect(Uri::from_static("/error"))
//         // }
        
//     //     // res;
//     // }else{
//     // }else{
//     //     warp::redirect(Uri::from_static("/error"))
//     //     // warp::reply::json(Err404{
//     //     //     text: "Not Found"
//     //     // })
//     //     //"error 404".to_string()
//     // }
// }


#[derive(Serialize, Deserialize, Debug, Clone)] //, RustcDecodable, RustcEncodable)]
struct MyItem {
    id: i32,
    tag: String,
    label: Vec<String>,
    short_text: String,
    text: String,
    id_page: i32,
    date_created: String,
    img: Option<String>,
    id_page_group: Vec<i32>,
}

#[derive(Serialize, Deserialize, Debug, Clone)] //, RustcDecodable, RustcEncodable)]
struct Content {
    content: Vec<MyItem>,
}


lazy_static! {
    #[derive(Debug)]
    static ref MyJson:Result<Content> = {
        let contents = fs::read_to_string("test.json")
        .expect("Should have been able to read the file");         //unsafe {
         let contents_json_obj: Content = serde_json::from_str(&contents)?;
         Ok(contents_json_obj)          // }
    }; }

lazy_static! { static ref MyJsonText:String = {
        serde_json::to_string(MyJson.as_ref().unwrap()).unwrap()     }; }

//lazy_static! {  static ref TestString:String = { "test".repeat(5000)     };}


#[tokio::main]
async fn main() {
    // GET /hello/warp => 200 OK with body "Hello, warp!"

    //http://balancer-domain/?video=http://s1.origin-cluster/video/1488/xcg2djHckad.m3u8
    //let mut rng = rand::thread_rng();
    //println!("{}", rng.gen::<f64>());

    //let v: Vec<&str> = "abc1,def,2ghi".split(",").collect();
    //println!("{:?} days", v);

    // let path = "http://s1.origin-cluster/video/1488/xcg2djHckad.m3u8";
    // let cdn_host = "http://localhost:3200";

    // let res = redirect_url::format_path(&path);
    // println!("{:?}", &res);
    // println!("{} {}", res.server, res.url);
    // let new_url = redirect_url::format_video_redirect_url(&cdn_host,res);
    // println!("{}", new_url);

    // let res = redirect_url::ab_redirect_by_rand(path, &cdn_host); //, &rng.gen::<f64>());
    // println!("{}",res);

    

    let error404 = warp::path("error").map(|| "error 404");
    // let head_only = warp::head() //.map(warp::reply)
    // //.and(warp::query::<HashMap<String, String>>())
    // //.and(warp::query::<ListOptions>())
    // .and(warp::query::<InputVideoUrl>())
    // .and(warp::path::full())
    // .map(|query: InputVideoUrl, p: FullPath| {
    //     resp(&query.video)
    // });
    // let filter = warp::any()
    //     //.and(warp::query::<HashMap<String, String>>())
    //     .and(warp::query::<InputVideoUrl>())
    //     //.and(warp::query()
    //     .map(|i: InputVideoUrl| {
    //         // let mut rng = rand::thread_rng();
    //         // if rng.gen::<f64>()<0.3 {
    //             let cdn_host = "http://localhost:3200";
    //             //let res = redirect_url::format_path(&i.video);
    //             //let new_url = redirect_url::format_video_redirect_url(&cdn_host, res);
    //             //let test = format!("{}", &new_url)
    //             let url = redirect_url::ab_redirect_by_rand(&i.video, &cdn_host); //, rng.gen::<f64>()); //redirect_url::format_video_redirect_url(&cdn_host, &*i.video);
    //             warp::redirect(Uri::from_str(&url).unwrap()) 
    //             //&i.video).unwrap()) //Uri::from_static(&url))
    //         // }else{
    //         //     warp::redirect(Uri::from_str(&i.video).unwrap())
    //         // }
    //         } );//);
        //.and(warp::query::<InputVideoUrl>())
        // .and(warp::path::full())
        // .map(|query: InputVideoUrl, p: FullPath| {
        //     resp(&query.video)
        // });
        // let hello = warp::path!("hello")
        // .map(|param: String| {
        //     format!("Hello {}", param)
        // });
    
    println!("{:?}", MyJson.as_ref().unwrap().content[0]); //MyJson.content[0]);


    let json_warp = warp::path!("json_warp").map(|| {
            let mystruct = MyJson.as_ref().unwrap();
            //.header("my-custom-header", "some-value")
            warp::reply::json(&mystruct)
    });
    let json_serde = warp::path!("json").map(|| {
        let mystruct = MyJson.as_ref().unwrap();
        Response::builder()
        .header("Content-Type", "application/json")
        .body( serde_json::to_string(mystruct).unwrap() ) //"and a custom body")
});

let json_text = warp::path!("json_text").map(|| {
    Response::builder()
    //.header("my-custom-header", "some-value")
    .body( MyJsonText.clone() ) //"and a custom body")
});
        let smalljson_warp = warp::path!("smalljson_warp").map(|| {
            
                //.header("my-custom-header", "some-value")
                let our_ids = vec![1, 3, 7, 13];
                warp::reply::json(&our_ids)
        });
        // index
        let test = warp::path!("test").map(|| {
            Response::builder()
                //.header("my-custom-header", "some-value")
                .body("test".repeat(5000) ) //"and a custom body")
        });
        // index
        let any = warp::path::end().map(|| {
            Response::builder()
                // .header("my-custom-header", "some-value")
                .body("hello")
        });
        let routes = warp::get().and(
            error404
            .or(smalljson_warp)
            .or(json_warp)
            .or(json_serde)
            .or(json_text)
            .or(test)
            .or(any)
               // .or(hello),
            //    .or(filter),
        ); //.or(head_only);
        //.map(|name| format!("video, {}!", name));

    warp::serve(routes)
        .run(([0, 0, 0, 0], 3200))
        .await;
}
 
