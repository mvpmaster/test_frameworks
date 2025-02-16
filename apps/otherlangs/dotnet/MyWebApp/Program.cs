global using FastEndpoints;

var builder = WebApplication.CreateBuilder();

 Console.WriteLine("CPU count: {0}", "hello");

// https://learn.microsoft.com/en-us/aspnet/core/fundamentals/logging/?view=aspnetcore-7.0
// builder.Logging.ClearProviders();
// builder.Logging.AddConsole();

builder.Services.AddFastEndpoints();

var app = builder.Build();
app.UseAuthorization();
app.UseFastEndpoints();
app.Run();
