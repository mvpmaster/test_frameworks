/*
 * Copyright (C) 2020-2022 Lightbend Inc. <https://www.lightbend.com>
 */

package docs.http.scaladsl

import akka.actor.typed.ActorSystem
import akka.actor.typed.scaladsl.Behaviors
import akka.http.scaladsl.Http
import akka.http.scaladsl.model._
import akka.http.scaladsl.server.Directives._

import scala.io.StdIn

import spray.json._

// https://blog.rockthejvm.com/akka-http-json/
// https://www.baeldung.com/scala/circe-json





case class ContentList_InputJson_TypeStructClass (
    content: List[InputJson_TypeStructClass],
)


case class InputJson_TypeStructClass(
  id: Int,
  tag: String,
  label:  List[String],
  short_text: String,
  text: String,
  id_page: Int,
  date_created: String,
  img: String,
  id_page_group: List[Int]
)


object QuickstartApp {

  def main(args: Array[String]): Unit = {

    implicit val system = ActorSystem(Behaviors.empty, "my-system")
    // needed for the future flatMap/onComplete in the end
    implicit val executionContext = system.executionContext

    //implicit val testJson = jsonFormat2(ContentList_InputJson_TypeStructClass)

    var test = "test"*5000

    val route = concat(
      path("hello") {
        get {
          complete(HttpEntity(ContentTypes.`text/html(UTF-8)`, "<h1>Say hello to akka-http</h1>"))
        }
      },
      path("test") {
        get {
          complete(HttpEntity(ContentTypes.`text/html(UTF-8)`, test))
        }
      }
    )

    val bindingFuture = Http().newServerAt("0.0.0.0", 8080).bind(route)

    println(s"Server now online. Please navigate to http://localhost:8080/hello\nPress RETURN to stop...")
    StdIn.readLine() // let it run until user presses return
    // bindingFuture
    //   .flatMap(_.unbind()) // trigger unbinding from the port
    //   .onComplete(_ => system.terminate()) // and shutdown when done
  }
}