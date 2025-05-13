using Velo.EventsService.Persistence.Context;
using Velo.EventsService.Persistence.Contracts;
using Velo.EventsService.Persistence.Entities;
using Velo.EventsService.Persistence.Repositories.Exceptions;

namespace Velo.EventsService.Persistence.Repositories;

public class EventsRepository(DatabaseContext context) : IEventsRepository
{
    public async Task DeleteEventAsync(int id, CancellationToken cancellationToken)
    {
        cancellationToken.ThrowIfCancellationRequested();

        var result = await context.Events.FindAsync([id], cancellationToken: cancellationToken);

        if (result == null)
            throw new EntityNotFoundException();

        context.Events.Remove(result);
        await context.SaveChangesAsync();
    }

    public async Task<EventEntity> PersistEventAsync(EventEntity eventEntity, CancellationToken cancellationToken)
    {
        cancellationToken.ThrowIfCancellationRequested();

        var result = await context.Events.AddAsync(eventEntity, cancellationToken);

        await context.SaveChangesAsync(cancellationToken);

        return result.Entity;
    }

    public async Task<EventEntity> UpdateEventAsync(int eventId, EventEntity eventEntity, CancellationToken cancellationToken)
    {
        cancellationToken.ThrowIfCancellationRequested();

        var entity = await context.Events.FindAsync([eventId, cancellationToken], cancellationToken);

        if (entity == null)
        {
            throw new EntityNotFoundException();
        }

        entity.Name = eventEntity.Name;
        entity.IsCanceled = eventEntity.IsCanceled;
        entity.WhenWillHappen = eventEntity.WhenWillHappen;
        entity.Description = eventEntity.Description;
        entity.Embeddings = eventEntity.Embeddings;

        await context.SaveChangesAsync(cancellationToken);

        return eventEntity;
    }

    public async Task UpdateEventImageAsync(int eventId, string path, CancellationToken cancellationToken)
    {
        var eventEntity = await context.Events.FindAsync([eventId, cancellationToken], cancellationToken);
        if (eventEntity == null)
        {
            throw new EntityNotFoundException();
        }
        eventEntity.PhotoPath = path;
        await context.SaveChangesAsync(cancellationToken);
    }
}