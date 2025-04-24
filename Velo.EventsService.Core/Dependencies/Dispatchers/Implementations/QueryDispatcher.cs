using Microsoft.Extensions.DependencyInjection;
using Velo.EventsService.Core.Dependencies.Contracts;
using Velo.EventsService.Core.Dependencies.Handlers;

namespace Velo.EventsService.Core.Dependencies.Dispatchers.Implementations;

public class QueryDispatcher(IServiceProvider serviceProvider) : IQueryDispatcher
{
    public Task<TQueryResult> Dispatch<TQuery, TQueryResult>(TQuery query, CancellationToken cancellationToken) where TQuery : IQuery
    {
        var handler = serviceProvider.GetRequiredService<IQueryHandler<TQuery, TQueryResult>>();
        return handler.Handle(query, cancellationToken);
    }
}