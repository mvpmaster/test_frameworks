name := "TicketsApp"

version := "1.0"

scalaVersion := "2.11.8"

lazy val versions = new {
  val finatra = "2.7.0"
  val logback = "1.1.3"
}

resolvers ++= Seq(
  Resolver.sonatypeRepo("releases"),
  "Twitter Maven" at "https://maven.twttr.com"
)

libraryDependencies += "com.twitter" %% "finatra-http" % versions.finatra