package com.mchougule.productsapp

import com.twitter.finagle.http.Request
import com.twitter.finatra.http.Controller
import com.twitter.finatra.validation.Size

import scala.collection.mutable

/**
  * Created by mchougule on 1/22/2017.
  */
class ProductResource extends Controller {
  val db = mutable.Map[String, List[Product]]()

  get("/products") { request: Request =>
    db
  }

  get("/products/:company") {request: Request =>
    db.getOrElse(request.params("company"), List())
  }

  post("/products") { company: Company =>
    val r = time(s"Total time for posting this product is %d ms") {
      val productsForCompany = db.get(company.company) match {
          // Find if this name already exists. If so then append to the list.
        case Some(products) => products :+ company
          // Create a new Map of company name and products.
        case None => List(company)
      }
      db.put(company.company, productsForCompany)
      response.created.location(s"/products/${company.productname}")
    }
    r
  }
}

case class Company(
                  @Size(min = 1, max = 25) productname: String,
                  company: String,
                  revenue: Int,
                  no_of_users: Int
                  )
