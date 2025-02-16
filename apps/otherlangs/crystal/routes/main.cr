require "router"
require "json"

puts "start app"

# https://crystal-lang.org/api/1.6.2/JSON.html

class WebServer
  include Router

  @@test = "glogal var"
  #@@test_json = "a"
  # https://crystal-lang.org/api/1.6.2/JSON/Any.html
  @test_json : JSON::Any

  def initialize()

    json= File.open("test.json") do |file|
      JSON.parse(file)
    end
    #puts json
    @test_json = json


    @test= ""
    a = 1
    while a < 5000
      @test += "test"
      a +=1
    end
    puts @test.size
  end

  def draw_routes
    get "/" do |context, params|
      context.response.print "Hello router.cr!"
      context
    end
    get "/test" do |context, params|
        context.response.print @test
        context
    end
    get "/json" do |context, params|
        context.response.print @test_json.to_json
        context
    end
  end
  
  def run
    server = HTTP::Server.new(route_handler)
    server.bind_tcp("0.0.0.0", 8080)
    #server.bind_tcp 8080
    server.listen
  end
end


web_server = WebServer.new
web_server.draw_routes
web_server.run
