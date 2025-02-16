import ../mike/src/mike
import strutils
import strformat
import tables

#
# Simple routing
#

get "/":
    ctx.send "Mike is running!"


run()
