import sbt._

val AkkaVersion = "2.7.0"
val AkkaHttpVersion = "10.4.0"


object Dependencies {
  lazy val scalaTest = "org.scalatest" %% "scalatest" % "3.2.11"

}

libraryDependencies ++= Seq(
  "com.typesafe.akka" %% "akka-actor-typed" % AkkaVersion,
  "com.typesafe.akka" %% "akka-stream" % AkkaVersion,
  "com.typesafe.akka" %% "akka-http" % AkkaHttpVersion
)