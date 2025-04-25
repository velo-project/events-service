using Velo.EventsService.Core.Contracts;
using Velo.EventsService.Core.Entities;
using Velo.EventsService.Persistence.Context;

namespace Velo.EventsService.Persistence.Repositories.Transactional;

public class TransactionalEventsRepository(DatabaseContext context) : ITransactionalEventsRepository
{
    public async Task<EventEntity> PersistEventAsync(EventEntity eventEntity, CancellationToken cancellationToken)
    {
        var result = await context.Events.AddAsync(eventEntity, cancellationToken);
        
        await context.SaveChangesAsync(cancellationToken);

        return result.Entity;
    }
}