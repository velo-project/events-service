using Microsoft.EntityFrameworkCore;
using Velo.EventsService.Commands.CreateEvent;
using Velo.EventsService.Commands.CreateEvent.Handler;
using Velo.EventsService.Core.Contracts;
using Velo.EventsService.Core.Dependencies.Dispatchers;
using Velo.EventsService.Core.Dependencies.Dispatchers.Implementations;
using Velo.EventsService.Core.Dependencies.Handlers;
using Velo.EventsService.Persistence.Context;
using Velo.EventsService.Persistence.Repositories.Transactional;

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

    public static void AddPersistence(this IServiceCollection services, IConfiguration configuration)
    {
        var connectionString = configuration["PostgreSQL:ConnectionString"] ?? throw new Exception("PostgreSQL:ConnectionString Is Null");

        services.AddDbContext<DatabaseContext>(options =>
        {
            options.UseNpgsql(connectionString, m =>
            {
                m.MigrationsAssembly(typeof(DatabaseContext).Assembly.FullName);
            });
        });
        services.AddScoped<ITransactionalEventsRepository, TransactionalEventsRepository>();
    }
}