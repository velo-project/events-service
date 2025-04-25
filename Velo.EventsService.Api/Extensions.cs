using Velo.EventsService.Core.Dependencies.Dispatchers;
using Velo.EventsService.Core.Dependencies.Dispatchers.Implementations;

namespace Velo.EventsService.Api;

public static class Extensions
{
    public static IServiceCollection AddMediator(this IServiceCollection services)
    {
        services.AddScoped<IQueryDispatcher, QueryDispatcher>();
        services.AddScoped<ICommandDispatcher, CommandDispatcher>();
        
        return services;
    }
}