using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using Velo.EventsService.Dependencies.Mediator.Handlers;
using Velo.EventsService.Persistence.Context;
using Velo.EventsService.Persistence.Entities;
using Velo.EventsService.Queries.FetchEventImage.Handler.Exceptions;

namespace Velo.EventsService.Queries.FetchEventImage.Handler;

public class FetchEventImageQueryHandler : IQueryHandler<FetchEventImageQuery, FetchEventimageQueryResult>
{
    private readonly DbSet<EventEntity> _events;
    private readonly string _imagePath;

    public FetchEventImageQueryHandler(DatabaseContext context, IConfiguration configuration)
    {
        _events = context.Set<EventEntity>();
    }
    
    public async Task<FetchEventimageQueryResult> Handle(FetchEventImageQuery query, CancellationToken cancellationToken)
    {
        cancellationToken.ThrowIfCancellationRequested();
        
        var eventEntity = await _events.Where(_ => _.Id == query.EventId)
            .SingleOrDefaultAsync(cancellationToken);

        // TODO: Refactor for NotFoundExceptions
        if (eventEntity == null)
            throw new EventNotFoundHttpException(query.EventId);

        if (eventEntity.PhotoPath == null)
            throw new EventWithNoImageHttpException();
        
        var image = await FetchImageBytesFrom(eventEntity.PhotoPath);
        return new FetchEventimageQueryResult()
        {
            Image = image
        };
    }

    private static async Task<byte[]> FetchImageBytesFrom(string path)
    {
        return await File.ReadAllBytesAsync(path);
    }
}