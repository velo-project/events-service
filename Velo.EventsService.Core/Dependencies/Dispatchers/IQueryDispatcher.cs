using Velo.EventsService.Core.Dependencies.Contracts;

namespace Velo.EventsService.Core.Dependencies.Dispatchers;

public interface IQueryDispatcher
{
    Task<TQueryResult> Dispatch<TQuery, TQueryResult>(TQuery query, CancellationToken cancellationToken)
        where TQuery : IQuery;
}