using Velo.EventsService.Api.Middlewares;

namespace Velo.EventsService.Api;

public class Program
{
    public static void Main(string[] args)
    {
        var builder = WebApplication.CreateBuilder(args);

        builder.Configuration.AddJsonFile(
            $"appsettings.{builder.Environment.EnvironmentName}.json", optional: false);
        
        // Add services to the container.
        builder.Services.AddMediator();
        builder.Services.AddCommands();
        builder.Services.AddQueries();
        builder.Services.AddPersistence(
            builder.Configuration);
        
        builder.Services.AddControllers();
        // Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
        builder.Services.AddEndpointsApiExplorer();
        builder.Services.AddSwaggerGen();

        var app = builder.Build();

        // Configure the HTTP request pipeline.
        if (app.Environment.IsDevelopment())
        {
            app.UseSwagger();
            app.UseSwaggerUI();
        }

        app.UseHttpsRedirection();

        app.UseAuthorization();
        app.UseMiddleware<HttpExceptionMiddleware>();

        app.MapControllers();

        app.Run();
    }
}