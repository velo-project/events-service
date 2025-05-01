using Velo.EventsService.Dependencies.Mediator.Contracts;

namespace Velo.EventsService.Queries.FilterEventsBySimilarity;

public class FilterEventsBySimilarityQuery : IQuery
{
    public required string Text { get; set; }
}