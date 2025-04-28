using Microsoft.AspNetCore.Mvc;
using Velo.EventsService.Dependencies.Mediator.Dispatchers;
using Velo.EventsService.Queries.FetchEventImage;

namespace Velo.EventsService.Api.Queries;

[ApiController]
[Route("/api")]
public class FetchEventImageController : Controller
{
    [HttpGet("v1/event/image/{id:int}")]
    public async Task<IActionResult> PerformV1(
        int id,
        IQueryDispatcher mediator,
        CancellationToken cancellationToken)
    {
        var command = new FetchEventImageQuery() { EventId = id };
        var result = await mediator.Dispatch<FetchEventImageQuery, FetchEventimageQueryResult>(command, cancellationToken);
        return File(result.Image, "image/jpeg");
    }
}