using Microsoft.AspNetCore.Mvc;
using Velo.EventsService.Api.Commands.DTOs;
using Velo.EventsService.Commands.CreateEvent;
using Velo.EventsService.Dependencies.Mediator.Dispatchers;

namespace Velo.EventsService.Api.Commands;

[ApiController]
[Route("/api")]
[Tags("Create Event", "Event Management")]
public class CreateEventController : ControllerBase
{
    [HttpPost("v1/event")]
    public async Task<IActionResult> PerformV1(
        [FromForm] CreateEventDTO dto,
        [FromServices] ICommandDispatcher mediator, 
        CancellationToken cancellationToken)
    {
        var photoBytes = await GeneratePhotoBytesFrom(dto);
        var command = MapDtoForCommand(dto, photoBytes);
        
        var result = await mediator.Dispatch<CreateEventCommand, CreateEventCommandResult>(command, cancellationToken);
        return StatusCode(result.StatusCode, result);
    }

    private static async Task<byte[]> GeneratePhotoBytesFrom(CreateEventDTO dto)
    {
        if (dto.File == null)
            return [];
        
        using var stream = new MemoryStream();
        await dto.File.CopyToAsync(stream);
        return stream.ToArray();
    } 

    private static CreateEventCommand MapDtoForCommand(CreateEventDTO dto, byte[] photoByes) => new()
    {
        Name = dto.Name,
        Description = dto.Description,
        PhotoBytes = photoByes,
        WhenWillHappen = dto.WhenWillHappen
    };
}