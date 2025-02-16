


1)

https://learn.microsoft.com/ru-ru/dotnet/core/install/linux-ubuntu#2004
https://dotnet.microsoft.com/en-us/download
https://dotnet.microsoft.com/en-us/download/dotnet/3.1

dotnet --version

or 

sudo snap install dotnet-sdk --classic



2)

works
https://fast-endpoints.com/docs/get-started#add-an-endpoint-class


  dotnet new web -n MyWebApp
  cd MyWebApp
  dotnet add package FastEndpoints



dotnet run


(old not work)

https://github.com/CarterCommunity/Carter

otnet new install CarterTemplate
dotnet new carter -n MyCarterApp -o MyCarterApp

or

otnet new web -n MyCarterApp
dotnet add package carter

dotnet run