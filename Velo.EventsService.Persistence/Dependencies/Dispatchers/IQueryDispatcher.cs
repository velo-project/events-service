using Velo.EventsService.Persistence.Dependencies.Contracts;

namespace Velo.EventsService.Persistence.Dependencies.Dispatchers;

public interface IQueryDispatcher
{
    Task<TQueryResult> Dispatch<TQuery, TQueryResult>(TQuery query, CancellationToken cancellationToken)
        where TQuery : IQuery;
}