using Microsoft.EntityFrameworkCore;
using Velo.EventsService.Commands.CreateEvent;
using Velo.EventsService.Commands.CreateEvent.Handler;
using Velo.EventsService.Commands.DeleteEvent;
using Velo.EventsService.Commands.DeleteEvent.Handler;
using Velo.EventsService.Commands.UpdateEvent;
using Velo.EventsService.Commands.UpdateEvent.Handler;
using Velo.EventsService.Commands.UpdateImageFromEvent;
using Velo.EventsService.Commands.UpdateImageFromEvent.Handler;
using Velo.EventsService.Dependencies.FileManager;
using Velo.EventsService.Dependencies.FileManager.Services;
using Velo.EventsService.Dependencies.Gemini;
using Velo.EventsService.Dependencies.Gemini.Services;
using Velo.EventsService.Dependencies.Mediator.Dispatchers;
using Velo.EventsService.Dependencies.Mediator.Dispatchers.Implementations;
using Velo.EventsService.Dependencies.Mediator.Handlers;
using Velo.EventsService.Persistence.Context;
using Velo.EventsService.Persistence.Contracts;
using Velo.EventsService.Persistence.Repositories;
using Velo.EventsService.Queries.FetchEventImage;
using Velo.EventsService.Queries.FetchEventImage.Handler;
using Velo.EventsService.Queries.FilterEventsBySimilarity;
using Velo.EventsService.Queries.FilterEventsBySimilarity.Handler;

namespace Velo.EventsService.Api;

public static class Extensions
{
    public static void AddMediator(this IServiceCollection services)
    {
        services.AddScoped<IQueryDispatcher, QueryDispatcher>();
        services.AddScoped<ICommandDispatcher, CommandDispatcher>();
    }

    public static void AddGemini(this IServiceCollection services)
    {
        services.AddHttpClient<IGeminiService, GeminiService>();
    }

    public static void AddFileManager(this IServiceCollection services)
    {
        services.AddScoped<IFileService, FileService>();
    }

    public static void AddCommands(this IServiceCollection services)
    {
        services.AddScoped<ICommandHandler<CreateEventCommand, CreateEventCommandResult>, CreateEventCommandHandler>();
        services.AddScoped<ICommandHandler<UpdateEventCommand, UpdateEventCommandResult>, UpdateEventCommandHandler>();
        services.AddScoped<ICommandHandler<UpdateImageFromEventCommand, UpdateImageFromEventCommandResult>, UpdateImageFromEventCommandHandler>();
        services.AddScoped<ICommandHandler<DeleteEventCommand, DeleteEventCommandResult>, DeleteEventCommandHandler>();
    }

    public static void AddQueries(this IServiceCollection services)
    {
        services.AddScoped<IQueryHandler<FetchEventImageQuery, FetchEventimageQueryResult>, FetchEventImageQueryHandler>();
        services
            .AddScoped<IQueryHandler<FilterEventsBySimilarityQuery, FilterEventsBySimilarityQueryResult>,
                FilterEventsBySimilarityQueryHandler>();
    }

    public static void AddPersistence(this IServiceCollection services, IConfiguration configuration)
    {
        var connectionString = configuration["PostgreSQL:ConnectionString"] ?? throw new Exception("PostgreSQL:ConnectionString Is Null");

        services.AddDbContext<DatabaseContext>(options =>
        {
            options.UseNpgsql(connectionString, npgssqlOptions =>
            {
                npgssqlOptions.UseVector();
                npgssqlOptions.MigrationsAssembly(typeof(DatabaseContext).Assembly.FullName);
            });
        });
        services.AddScoped<IEventsRepository, EventsRepository>();
    }
}