//using static Globals;

public class MyEndpoint : Endpoint<MyRequest>
{
    //public static test String; //new string('test', 10);

    //public string _test = "test";
    
    public override void Configure()
    {
        //test = string.Concat(Enumerable.Repeat("test", 10));//new string('test', 2);
        //Config.Application["test"] = "noodle";
        Get("/"); //api/user/create");
        AllowAnonymous();
    }

    public override async Task HandleAsync(MyRequest req, CancellationToken ct)
    {
        if (Globals.test.Length <10){ //_test.Length < 50) { 
            Console.WriteLine("test {0}", Globals.test);
            //Globals.test+="1";
            Console.WriteLine("Init Test String {0}",Globals.test); //_test.Length);
            //_test = 
            Globals.test=string.Concat(Enumerable.Repeat("test", 5000)); 
            }
        var response = new MyResponse()
        {
            FullName = Globals.test, //_test, //Config["test"],  //"Test", //req.FirstName + " " + req.LastName,
            IsOver18 = true //req.Age > 18
        };

        await SendAsync(response);
    }
}
