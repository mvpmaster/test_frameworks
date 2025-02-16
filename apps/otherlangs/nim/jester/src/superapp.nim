# example.nim
import htmlgen
import jester
import std/strutils

let test = "test".indent(0).repeat(5000).strip

routes:
  get "/":
    resp test
    #resp h1(test)
