using Velo.EventsService.Persistence.Dependencies.Contracts;

namespace Velo.EventsService.Persistence.Dependencies.Handlers;

public interface IQueryHandler<in TQuery, TQueryResult> where TQuery : IQuery
{
    Task<TQueryResult> Handle(TQuery query, CancellationToken cancellationToken);
}