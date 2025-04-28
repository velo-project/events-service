using Velo.EventsService.Dependencies.Mediator.Contracts;

namespace Velo.EventsService.Queries.FetchEventImage;

public class FetchEventImageQuery : IQuery
{
    public int EventId { get; set; }
}