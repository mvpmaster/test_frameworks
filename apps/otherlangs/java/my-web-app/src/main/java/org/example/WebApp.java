package org.example;


import io.activej.http.AsyncServlet;
import io.activej.http.HttpResponse;
import io.activej.inject.annotation.Provides;
import io.activej.launcher.Launcher;
import io.activej.launchers.http.HttpServerLauncher;
import io.activej.promise.Promise;

// http://www.mastertheboss.com/java-ee/json/how-to-parse-json-in-java/
// jeson model
// https://go.libhunt.com/jsonparser-alternatives
// https://www.javatpoint.com/how-to-convert-string-to-json-object-in-java
// class example
// https://mkyong.com/java/how-do-convert-java-object-to-from-json-format-gson-api/
// https://www.google.com/search?q=gson+java+example
// code examle
// https://www.java67.com/2016/10/3-ways-to-convert-string-to-json-object-in-java.html
// https://javarevisited.blogspot.com/2013/02/how-to-convert-json-string-to-java-object-jackson-example-tutorial.html#axzz7lszgvq62
// https://stackoverflow.com/questions/10308452/how-to-convert-the-following-json-string-to-java-object
import org.json.JSONException;
import org.json.JSONObject;    
import com.google.gson.Gson;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonToken;
// https://www.geeksforgeeks.org/how-to-setup-jackson-in-java-application/
import com.fasterxml.jackson.core.JsonParseException;
import com.fasterxml.jackson.databind.JsonMappingException;
// https://www.tabnine.com/blog/how-to-convert-a-java-object-into-a-json-string/
// https://www.google.com/search?q=java+class+to+string+json+alternatives
import com.fasterxml.jackson.databind.ObjectMapper;
//import org.json.simple.parser.JSONParser;
//import java.io.*;

// https://github.com/activej/activej/blob/examples-5.4.3/examples/core/http/src/main/java/RoutingServletExample.java
import io.activej.http.RoutingServlet;
import static io.activej.http.HttpMethod.GET;
// https://github.com/activej/activej/blob/examples-5.4.3/examples/core/http/src/main/java/BlockingServletExample.java
import java.util.concurrent.Executor;



//import java.io.FileReader;


// https://stackoverflow.com/questions/326390/how-do-i-create-a-java-string-from-the-contents-of-a-file
// https://www.geeksforgeeks.org/files-class-readstring-method-in-java-with-examples/
// Implementation fo readString() method in Java
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

import java.nio.charset.StandardCharsets;

//import  InputJson_TypeStructClass;




public class WebApp extends HttpServerLauncher {

    public static String test;
    public static JSONObject myJsonObject_Simple;
    public static ContentList_InputJson_TypeStructClass myJsonObject_Gson;
    public static ContentList_InputJson_TypeStructClass myJsonObject_Jackson;
    public static Gson gson;
	public static ObjectMapper jacksonMapper;



@Provides
 AsyncServlet servlet(){ //Executor executor) { 
     return RoutingServlet.create() // Pool
   // https://www.javacodegeeks.com/2020/10/creating-high-performance-web-applications-with-activej.html
   //https://activej.io/http/examples
          // 35000 rps - 64000 rps (70000) / 170000
          .map(GET,
          "/", request ->
                 HttpResponse.ok200()
                        .withPlainText(WebApp.test))
                    //    .withHtml("<h1>Go to some pages</h1>" +
                    //          "<a href=\"/path1\"> Path 1 </a><br>" +
                    //          "<a href=\"/path2\"> Path 2 </a><br>"))
           // 25700 rps (47000~)
           .map(GET, 
           "/jackson", request ->
                 HttpResponse.ok200()
                       .withPlainText(WebApp.jacksonMapper.writeValueAsString(WebApp.myJsonObject_Jackson)))
           // 11000- 19000 rps (20000+)
           .map(GET, 
           "/gson", request ->
                 HttpResponse.ok200()
                       .withPlainText(WebApp.gson.toJson(WebApp.myJsonObject_Gson)))
           // 13000 - 19000 rps (26000)
           .map(GET, 
           "/json_jackson_save", request ->
                 HttpResponse.ok200()
                       .withPlainText(WebApp.gson.toJson(WebApp.myJsonObject_Simple)))
           // 9000 - 16000 rps (17-19 max)
           .map(GET, 
           "/json", request ->
                 HttpResponse.ok200()
                       .withPlainText(WebApp.myJsonObject_Simple.toString()))
           .map("/*", request ->
                 HttpResponse.ofCode(404)
                       .withHtml("<h1>404</h1><p>Path '" + request.getRelativePath() + "' not found</p>" +
                         "<a href=\"/\">Go home</a>"));
    }


    // @Provides
    // AsyncServlet servlet() {
    //     return request -> Promise.of(
    //             HttpResponse.ok200()
    //                 .withPlainText(WebApp.test)); //"Hello, World!"));
    // }

    public static void main(String[] args) throws Exception {
        //JSONParser parser = new JSONParser();
        //Gson gson = new Gson();
        WebApp.gson = new Gson();
        WebApp.jacksonMapper = new ObjectMapper();
        System.out.println("Hello, World!");
        try {
            Path filePath = Path.of("./test.json");
            String jsonString = Files.readString(filePath,  StandardCharsets.UTF_8); //encoding);
            // System.out.println(content);
            //https://stackoverflow.com/questions/10926353/how-to-read-json-file-into-java-with-simple-json-library
            //System.out.println(new FileReader("./test.json"));
            //JSONObject jsonObject = 
            
            //InputJson_TypeStructClass myJsonObject = 
            //System.out.println(jsonString);

            // 16500 - 19000 rps
            WebApp.myJsonObject_Gson =  WebApp.gson.fromJson(jsonString, ContentList_InputJson_TypeStructClass.class);
            // 25700 rps
            WebApp.myJsonObject_Jackson = new ObjectMapper().readValue(jsonString, ContentList_InputJson_TypeStructClass.class);

            //Read more: https://www.java67.com/2016/10/3-ways-to-convert-string-to-json-object-in-java.html#ixzz7lt2Kjhg8

            // 11000 - 16500 rps
            WebApp.myJsonObject_Simple = new JSONObject(jsonString); //(JSONObject) parser.parse(new FileReader("./test.json"));
            //JSONObject jsonObject = new JSONObject(//"{\"phonetype\":\"N95\",\"cat\":\"WP\"}");
            System.out.println(WebApp.gson.toJson(WebApp.myJsonObject_Gson)); //WebApp.myJsonObject); //WebApp.gson.toJson(WebApp.myJsonObject)); //WebApp.myJsonObject); //jsonObject.toString());
        }catch (JSONException err){
           System.out.println(err.getMessage());

            //Log.d("Error", err.toString());
        }

        WebApp.test = new String(new char[5000]).replace("\0", "test");
        Launcher launcher = new WebApp();
        launcher.launch(args);
    }
}
