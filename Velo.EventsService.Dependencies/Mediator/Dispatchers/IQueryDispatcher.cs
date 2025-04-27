using Velo.EventsService.Dependencies.Mediator.Contracts;

namespace Velo.EventsService.Dependencies.Mediator.Dispatchers;

public interface IQueryDispatcher
{
    Task<TQueryResult> Dispatch<TQuery, TQueryResult>(TQuery query, CancellationToken cancellationToken)
        where TQuery : IQuery;
}