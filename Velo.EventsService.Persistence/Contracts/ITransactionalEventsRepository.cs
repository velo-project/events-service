using Velo.EventsService.Persistence.Entities;

namespace Velo.EventsService.Persistence.Contracts;

public interface ITransactionalEventsRepository
{
    Task<EventEntity> PersistEventAsync(EventEntity eventEntity, CancellationToken cancellationToken);
}