using Microsoft.AspNetCore.Mvc;
using Velo.EventsService.Commands.DeleteEvent;
using Velo.EventsService.Dependencies.Mediator.Dispatchers;

namespace Velo.EventsService.Api.Commands
{
    [ApiController]
    [Route("/api")]
    [Tags("Delete Event", "Event Management")]
    public class DeleteEventController : ControllerBase
    {
        [HttpDelete("v1/event/{id:int}")]
        public async Task<IActionResult> PerformV1(
            int id,
            ICommandDispatcher mediator,
            CancellationToken cancellationToken)
        {
            var command = new DeleteEventCommand
            {
                Id = id,
            };
            await mediator.Dispatch<DeleteEventCommand, DeleteEventCommandResult>(command, cancellationToken);

            return NoContent();
        }
    }
}
