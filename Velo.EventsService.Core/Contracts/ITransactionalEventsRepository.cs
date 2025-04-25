using Velo.EventsService.Core.Entities;

namespace Velo.EventsService.Core.Contracts;

public interface ITransactionalEventsRepository
{
    Task<EventEntity> PersistEventAsync(EventEntity eventEntity, CancellationToken cancellationToken);
}