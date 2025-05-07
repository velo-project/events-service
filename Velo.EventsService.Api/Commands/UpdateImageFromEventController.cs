using Microsoft.AspNetCore.Mvc;
using Velo.EventsService.Api.Commands.DTOs;
using Velo.EventsService.Commands.UpdateImageFromEvent;
using Velo.EventsService.Dependencies.Mediator.Dispatchers;

namespace Velo.EventsService.Api.Commands
{
    [ApiController]
    [Route("/api")]
    [Tags("Update Event", "Event Management")]
    public class UpdateImageFromEventController : ControllerBase
    {
        [HttpPut("v1/event/image/{id:int}")]
        public async Task<IActionResult> PerformV1(
            int id, 
            IFormFile image,
            ICommandDispatcher mediator, 
            CancellationToken cancellation)
        {
            var bytes = await GeneratePhotoBytesFrom(image);
            var command = new UpdateImageFromEventCommand
            {
                EventId = id,
                ImageBytes = bytes
            };
            await mediator.Dispatch<UpdateImageFromEventCommand, UpdateImageFromEventCommandResult>(command, cancellation);
            return NoContent();
        }

        private static async Task<byte[]> GeneratePhotoBytesFrom(IFormFile file)
        {
            if (file == null)
                return [];

            using var stream = new MemoryStream();
            await file.CopyToAsync(stream);
            return stream.ToArray();
        }
    }
}
