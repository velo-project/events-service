using Velo.EventsService.Persistence.Entities;

namespace Velo.EventsService.Persistence.Contracts;

public interface IEventsRepository
{
    Task<EventEntity> PersistEventAsync(EventEntity eventEntity, CancellationToken cancellationToken);
    Task<EventEntity> UpdateEventAsync(int eventId, EventEntity eventEntity, CancellationToken cancellationToken);
}