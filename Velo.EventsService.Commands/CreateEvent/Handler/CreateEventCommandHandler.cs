using Velo.EventsService.Core.Contracts;
using Velo.EventsService.Core.Dependencies.Handlers;
using Velo.EventsService.Core.Entities;

namespace Velo.EventsService.Commands.CreateEvent.Handler;

public class CreateEventCommandHandler(ITransactionalEventsRepository eventsRepository)
    : ICommandHandler<CreateEventCommand, CreateEventCommandResult>
{
    public async Task<CreateEventCommandResult> Handle(CreateEventCommand command, CancellationToken cancellationToken)
    {
        cancellationToken.ThrowIfCancellationRequested();

        var eventToSave = ExtractAndPrepareEventFrom(command);
        var savedEvent = await eventsRepository.PersistEventAsync(eventToSave, cancellationToken);

        return PrepareSuccessResult(savedEvent);
    }

    private static EventEntity ExtractAndPrepareEventFrom(CreateEventCommand command)
    {
        return new EventEntity()
        {
            Name = command.Name,
            Description = command.Description,
            IsActive = true,
            IsCanceled = false,
            WhenWillHappen = command.WhenWillHappen,
            CreatedAt = DateTime.UtcNow
        };
    }

    private static CreateEventCommandResult PrepareSuccessResult(EventEntity eventEntity) => new(
            "Created.",
            eventEntity,
            201);
}