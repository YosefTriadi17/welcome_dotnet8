using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using System;
using System.Diagnostics;
using welcome.Models;
using StackExchange.Redis;

namespace welcome.Controllers
{
    public class HomeController : Controller
    {
        private readonly ILogger<HomeController> _logger;

        public HomeController(ILogger<HomeController> logger)
        {
            _logger = logger;
        }

        public IActionResult Index()
        {
            //var muxer = ConnectionMultiplexer.Connect(
            //    new ConfigurationOptions
            //    {
            //        EndPoints = { { "redis-14506.c228.us-central1-1.gce.redns.redis-cloud.com", 14506 } },
            //        User = "default",
            //        Password = "AQp8KNpFXBcvCy9YgVqUWQFwTd7sIFeJ"
            //    }
            //);
            //var db = muxer.GetDatabase();

            //db.StringSet("foo", "barikade sudosu");
            //db.StringSet("fee", "bar");
            //RedisValue result = db.StringGet("foo");
            //Console.WriteLine(result); // >>> bar
            //return Ok(result);
            return View();
        }

        public IActionResult Privacy()
        {
            return View();
        }

        [ResponseCache(Duration = 0, Location = ResponseCacheLocation.None, NoStore = true)]
        public IActionResult Error()
        {
            return View(new ErrorViewModel { RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier });
        }

        public IActionResult LoadData()
        {
            return View();
        }
    }
}
