using Velo.EventsService.Persistence.Entities;

namespace Velo.EventsService.Queries.FilterEventsBySimilarity;

public class FilterEventsBySimilarityQueryResult
{
    public List<EventEntity> Events { get; set; } = [];
}