using Microsoft.Extensions.DependencyInjection;
using Velo.EventsService.Dependencies.Mediator.Contracts;
using Velo.EventsService.Dependencies.Mediator.Handlers;

namespace Velo.EventsService.Dependencies.Mediator.Dispatchers.Implementations;

public class QueryDispatcher(IServiceProvider serviceProvider) : IQueryDispatcher
{
    public Task<TQueryResult> Dispatch<TQuery, TQueryResult>(TQuery query, CancellationToken cancellationToken) where TQuery : IQuery
    {
        var handler = serviceProvider.GetRequiredService<IQueryHandler<TQuery, TQueryResult>>();
        return handler.Handle(query, cancellationToken);
    }
}