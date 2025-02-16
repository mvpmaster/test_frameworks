package com.mchougule.orgsapp

import com.twitter.finagle.http.Request
import com.twitter.finatra.http.{Controller, HttpServer}
import com.twitter.finatra.http.routing.HttpRouter
import com.mchougule.productsapp.ProductResource
import com.twitter.finatra.http.filters.CommonFilters
/**
  * Created by mchougule on 1/22/2017.
  */
object OrgsApp extends OrgsServer

class OrgsServer extends HttpServer {
  override protected def configureHttp(router: HttpRouter): Unit = {
    router
        .filter[CommonFilters]
      .add[HiController]
      .add[ProductResource]
  }
}
class HiController extends Controller {
  get("/hello") { request: Request =>
    "Hello there you!"
  }
}
