using Microsoft.AspNetCore.Mvc;
using Velo.EventsService.Commands.CreateEvent;
using Velo.EventsService.Core.Dependencies.Dispatchers;

namespace Velo.EventsService.Api.Commands;

[ApiController]
[Route("/api")]
public class CreateEventController : ControllerBase
{
    [HttpPost("v1/event")]
    public async Task<IActionResult> PerformV1(
        [FromBody] CreateEventCommand command, 
        [FromServices] ICommandDispatcher mediator, 
        CancellationToken cancellationToken)
    {
        var result = await mediator.Dispatch<CreateEventCommand, CreateEventCommandResult>(command, cancellationToken);
        return StatusCode(result.StatusCode, result);
    }
}