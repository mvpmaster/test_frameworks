var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();


string test = string.Concat(Enumerable.Repeat("test", 5000)); 

app.MapGet("/", () => "Hello World!");
app.MapGet("/test", () => test);

app.Run();
