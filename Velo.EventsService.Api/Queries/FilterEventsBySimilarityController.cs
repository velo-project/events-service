using Microsoft.AspNetCore.Mvc;
using Velo.EventsService.Dependencies.Mediator.Dispatchers;
using Velo.EventsService.Queries.FilterEventsBySimilarity;

namespace Velo.EventsService.Api.Queries;

[ApiController]
[Route("/api")]
[Tags("Filter Events", "Search By Similarity", "Get Event", "Event Management")]
public class FilterEventsBySimilarityController(IQueryDispatcher mediator) : ControllerBase
{
    [HttpGet("v1/event")]
    public async Task<IActionResult> PerformV1([FromQuery] string queryValue, CancellationToken cancellationToken)
    {
        var query = new FilterEventsBySimilarityQuery()
        {
            Text = queryValue
        };
        
        var result =
            await mediator.Dispatch<FilterEventsBySimilarityQuery, FilterEventsBySimilarityQueryResult>(query, cancellationToken);
        
        return Ok(new
        {
            Message = "Ok",
            Events = result.Events
        });
    }
}