using Microsoft.EntityFrameworkCore;
using Velo.EventsService.Commands.CreateEvent;
using Velo.EventsService.Commands.CreateEvent.Handler;
using Velo.EventsService.Dependencies.Mediator.Dispatchers;
using Velo.EventsService.Dependencies.Mediator.Dispatchers.Implementations;
using Velo.EventsService.Dependencies.Mediator.Handlers;
using Velo.EventsService.Persistence.Context;
using Velo.EventsService.Persistence.Contracts;
using Velo.EventsService.Persistence.Repositories;
using Velo.EventsService.Queries.FetchEventImage;
using Velo.EventsService.Queries.FetchEventImage.Handler;

namespace Velo.EventsService.Api;

public static class Extensions
{
    public static void AddMediator(this IServiceCollection services)
    {
        services.AddScoped<IQueryDispatcher, QueryDispatcher>();
        services.AddScoped<ICommandDispatcher, CommandDispatcher>();
    }

    public static void AddCommands(this IServiceCollection services)
    {
        services.AddScoped<ICommandHandler<CreateEventCommand, CreateEventCommandResult>, CreateEventCommandHandler>();
    }

    public static void AddQueries(this IServiceCollection services)
    {
        services.AddScoped<IQueryHandler<FetchEventImageQuery, FetchEventimageQueryResult>, FetchEventImageQueryHandler>();
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