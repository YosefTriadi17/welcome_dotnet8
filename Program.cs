var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddControllersWithViews();

var app = builder.Build();

// === Path base untuk /welcome ===
//app.UsePathBase("/welcome");
app.UseRouting();

app.Use(async (context, next) =>
{
    // Pastikan URL tetap benar saat diakses tanpa /welcome/
    if (!context.Request.Path.StartsWithSegments("/welcome"))
    {
        var path = "/welcome" + context.Request.Path + context.Request.QueryString;
        context.Response.Redirect(path, permanent: false);
        return;
    }
    await next();
});

// === Penting: file statis harus aware path base ===
app.UseStaticFiles(new StaticFileOptions
{
    OnPrepareResponse = ctx =>
    {
        ctx.Context.Response.Headers["Cache-Control"] = "public,max-age=600";
    }
});

// Middleware pipeline standar
if (!app.Environment.IsDevelopment())
{
    app.UseExceptionHandler("/Home/Error");
    app.UseHsts();
}

app.UseHttpsRedirection();
app.UseAuthorization();

app.MapControllerRoute(
    name: "default",
    pattern: "{controller=Home}/{action=Index}/{id?}");

app.Run();
