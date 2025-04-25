using Velo.EventsService.Persistence.Entities;

namespace Velo.EventsService.Commands.CreateEvent;

public record CreateEventCommandResult(
    string Message,
    EventEntity SavedEvent,
    int StatusCode);