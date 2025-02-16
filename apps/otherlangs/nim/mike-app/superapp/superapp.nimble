# Package

version       = "0.1.0"
author        = "North it"
description   = "A new awesome nimble package"
license       = "MIT"
srcDir        = "src"
bin           = @["superapp"]


# Dependencies

requires "nim >= 1.6.8"
requires "httpx >= 0.3.2, zippy >= 0.10.4, websocketx >= 0.1.2"
