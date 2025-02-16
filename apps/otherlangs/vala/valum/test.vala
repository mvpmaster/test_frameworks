using Valum;
using VSGI;

var app = new Router ();

app.use (basic ()); /* handle stuff like 404 errors and more */

app.get ("/", (req, res) => {
    res.headers.set_content_type ("text/plain", null);
    return res.expand_utf8 ("Hello world!");
});

Server.@new ("http", handler: app).run ({"app", "--address=0.0.0.0:3003", "--forks=4"});