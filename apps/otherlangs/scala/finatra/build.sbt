import Dependencies._

ThisBuild / scalaVersion     := "2.13.8"
ThisBuild / version          := "0.1.0-SNAPSHOT"
ThisBuild / organization     := "com.example"
ThisBuild / organizationName := "example"

lazy val root = (project in file("."))
  .settings(
    name := "finatra",
    //libraryDependencies += scalaTest % Test
    libraryDependencies ++= Seq(
      "com.twitter" %% "finatra-http-server" % "22.4.0",
    "com.twitter" %% "finatra-jackson" % "22.4.0" ,
     "com.twitter" %% "inject-server" % "22.4.0" ,
     "com.twitter" %% "inject-app" % "22.4.0" ,
      "com.twitter" %% "inject-core" % "22.4.0" ,
      "com.twitter" %% "inject-modules" % "22.4.0" ,
      // "com.typesafe.akka" %% "akka-http"                % akkaHttpVersion,
      // "com.typesafe.akka" %% "akka-http-spray-json"     % akkaHttpVersion,
      // "com.typesafe.akka" %% "akka-actor-typed"         % akkaVersion,
      // "com.typesafe.akka" %% "akka-stream"              % akkaVersion,
      "ch.qos.logback"    % "logback-classic"           % "1.2.11",

      // "com.typesafe.akka" %% "akka-http-testkit"        % akkaHttpVersion % Test,
      // "com.typesafe.akka" %% "akka-actor-testkit-typed" % akkaVersion     % Test,
      "org.scalatest"     %% "scalatest"                % "3.2.9"         % Test
    )
  )

// See https://www.scala-sbt.org/1.x/docs/Using-Sonatype.html for instructions on how to publish to Sonatype.

