Mix.install([:plug, :plug_cowboy])



# defmodule MyRouter do
#   use Plug.Router

#   plug :match
#   plug :dispatch

#   get "/hello" do
#     send_resp(conn, 200, "world")
#   end

#   forward "/users", to: UsersRouter

#   match _ do
#     send_resp(conn, 404, "oops")
#   end
# end

defmodule MyPlug do
  import Plug.Conn
  @test String.duplicate("test", 5000 )

  def init(options) do
    # initialize options

    options
  end

  def call(conn, _opts) do
    conn
    |> put_resp_content_type("text/plain")
    |> send_resp(200, "Hello world")
  end
end

require Logger
{:ok, _} = Plug.Cowboy.http(MyPlug, [])
Logger.info("Plug now running on localhost:4000")
