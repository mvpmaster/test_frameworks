defmodule MyApp do
  @moduledoc """
  Documentation for MyApp.
  """

  @doc """
  Hello world.

  ## Examples

      iex> MyApp.hello()
      :world

  """
  def hello do
    :world
  end
end



# defmodule MyApp.Application do
#   # See https://hexdocs.pm/elixir/Application.html
#   # for more information on OTP Applications
#   @moduledoc false

#   use Application

#   def start(_type, _args) do
#     # List all child processes to be supervised
#     children = [
#       {Plug.Cowboy, scheme: :http, plug: MyPlug, options: [port: 4001]}
#     ]

#     # See https://hexdocs.pm/elixir/Supervisor.html
#     # for other strategies and supported options
#     opts = [strategy: :one_for_one, name: MyApp.Supervisor]
#     Supervisor.start_link(children, opts)
#   end
# end
