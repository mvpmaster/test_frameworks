using HTTP

# start a blocking server
HTTP.listen() do http::HTTP.Stream
    @show http.message
    @show HTTP.header(http, "Content-Type")
    while !eof(http)
        println("body data: ", String(readavailable(http)))
    end
    HTTP.setstatus(http, 200)
    HTTP.setheader(http, "Content-Type" => "application/json")
    HTTP.startwrite(http)
    write(http, "response body")
    write(http, "more response body")
end