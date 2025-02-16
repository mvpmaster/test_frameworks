A simple REST API to add and list products for the company using a lightweight scala framework Finatra.
[Finatra](https://twitter.github.io/finatra/) is used for building applications in Scala on top of [TwitterServer](https://github.com/twitter/twitter-server) and [Finagle](https://github.com/twitter/finagle).


1. I created a server OrgsServer which extends com.twitter.finatra.http.HttpServer.
2. object OrgsApp is an object which will be used to launch the server.
3. We override configureHttp and define router controllers. Endpoints are defined in these controllers.
4. Right now I am only using mutable Map, it will act as a database. TODO: Connect to the database using Slick.
5. This mutable Map will hold Company Names as a key and their list of products for values.

Follow these steps to get started:

1. Git-clone this repository.

        $ git clone https://github.com/mmchougule/finatra-example.git finatra-example

2. Open this project in IntelliJ IDEA -> Go to OrgsApp.scala and run OR Ctrl+Shift+F10

3. Browse to [http://localhost:8888/]
	Example calls

        > http://localhost:8888/hello

        > http://localhost:8888/products/ (GET) -> To show the list of products by all companies.

        > (POST) -> POST data to the application.
         curl -H "Content-Type: application/json" -X POST -d '{"company":"Actio", "no_of_users":"100","productname": "Vault", "revenue": "2000000"}' http://localhost:8888/products
         curl -H "Content-Type: application/json" -X POST -d '{"company":"Test Company", "no_of_users":"1000","productname": "Ticketing", "revenue": "1000000"}' http://localhost:8888/products

