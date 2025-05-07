using Microsoft.AspNetCore.Mvc;
using Velo.EventsService.Commands.UpdateEvent;
using Velo.EventsService.Dependencies.Mediator.Dispatchers;

namespace Velo.EventsService.Api.Commands
{
    [ApiController]
    [Route("/api")]
    [Tags("Update Event", "Event Management")]
    public class UpdateEventController : ControllerBase
    {
        [HttpPut("v1/event/{id:int}")]
        public async Task<IActionResult> PerformV1(
            int id, 
            UpdateEventCommand command, 
            ICommandDispatcher mediator, 
            CancellationToken cancellationToken)
        {
            await mediator.Dispatch<UpdateEventCommand, UpdateEventCommandResult>(command, cancellationToken);

            return NoContent();
        }
    }
}
