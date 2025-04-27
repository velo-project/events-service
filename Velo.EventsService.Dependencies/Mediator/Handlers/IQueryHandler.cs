using Velo.EventsService.Dependencies.Mediator.Contracts;

namespace Velo.EventsService.Dependencies.Mediator.Handlers;

public interface IQueryHandler<in TQuery, TQueryResult> where TQuery : IQuery
{
    Task<TQueryResult> Handle(TQuery query, CancellationToken cancellationToken);
}