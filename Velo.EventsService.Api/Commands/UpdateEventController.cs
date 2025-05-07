using Microsoft.AspNetCore.Mvc;
using Velo.EventsService.Api.Commands.DTOs;
using Velo.EventsService.Commands.UpdateEvent;
using Velo.EventsService.Dependencies.Mediator.Dispatchers;
using Velo.EventsService.Persistence.Entities;

namespace Velo.EventsService.Api.Commands
{
    [ApiController]
    [Route("/api")]
    [Tags("Update Event", "Event Management")]
    public class UpdateEventController : ControllerBase
    {
        [HttpPatch("v1/event/{id:int}")]
        public async Task<IActionResult> PerformV1(
            int id, 
            UpdateEventDTO dto, 
            ICommandDispatcher mediator, 
            CancellationToken cancellationToken)
        {
            var command = new UpdateEventCommand
            {
                EventId = id,
                Event = new EventEntity
                {
                    Name = dto.Name,
                    Description = dto.Description,
                    WhenWillHappen = dto.WhenWillHappen,
                    IsCanceled = dto.IsCanceled,
                }
            };

            await mediator.Dispatch<UpdateEventCommand, UpdateEventCommandResult>(command, cancellationToken);

            return NoContent();
        }
    }
}
