
require "toro"

puts "start app"


class App < Toro::Router
  @@test="1"
  # @@test = "glogal var"

  a = 1
  while a < 5000
    @@test += "test"
    a +=1
  end
  puts @@test.size
  # def initialize()
  #   @test= ""
  #   a = 1
  #   while a < 5000
  #     @test += "test"
  #     a +=1
  #   end
  #   puts @test.size
  # end

  def routes
    get do
      text "hello world"
    end
    on "test" do
      get do
        text @@test
      end
    end
  end

end


 App.run(8080) # do |server|
#   server.test="test"
# # ssl = OpenSSL::SSL::Context::Server.new
# # ssl.private_key = "path/to/private_key"
# # ssl.certificate_chain = "path/to/certificate_chain"
# # server.tls = ssl
# end