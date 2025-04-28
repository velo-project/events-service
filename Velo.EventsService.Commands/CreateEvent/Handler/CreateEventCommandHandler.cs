
using Microsoft.Extensions.Configuration;
using Velo.EventsService.Dependencies.Mediator.Handlers;
using Velo.EventsService.Persistence.Contracts;
using Velo.EventsService.Persistence.Entities;

namespace Velo.EventsService.Commands.CreateEvent.Handler;

public class CreateEventCommandHandler(IEventsRepository eventsRepository, IConfiguration configuration)
    : ICommandHandler<CreateEventCommand, CreateEventCommandResult>
{
    private readonly string _imageFolderPath = configuration["ImageFolderPath"] ?? throw new InvalidOperationException();
    
    public async Task<CreateEventCommandResult> Handle(CreateEventCommand command, CancellationToken cancellationToken)
    {
        cancellationToken.ThrowIfCancellationRequested();

        var imagePath = await ProcessAndSaveImage(command);
        var eventToSave = ExtractAndPrepareEventFrom(command, imagePath);
        
        var savedEvent = await eventsRepository.PersistEventAsync(eventToSave, cancellationToken);

        if (savedEvent.PhotoPath != null)
            savedEvent.PhotoPath = null;

        return PrepareSuccessResult(savedEvent);
    }

    private async Task<string?> ProcessAndSaveImage(CreateEventCommand command)
    {
        if (command.PhotoBytes.Length == 0)
            return null;

        var folderCombination = $"{DateTime.UtcNow.Year}/{DateTime.UtcNow.Month}/{DateTime.UtcNow.Day}";
        var folderPath = Path.Combine(_imageFolderPath, folderCombination);
        if (!Directory.Exists(folderPath))
            Directory.CreateDirectory(folderPath);

        var imageName = Guid.NewGuid().ToString().Replace("-", "") + ".png";
        var filePath = Path.Combine(folderPath, imageName);

        await File.WriteAllBytesAsync(filePath, command.PhotoBytes);

        return filePath;
    }

    private static EventEntity ExtractAndPrepareEventFrom(CreateEventCommand command, string? imagePath)
    {
        return new EventEntity()
        {
            Name = command.Name,
            Description = command.Description,
            IsActive = true,
            IsCanceled = false,
            WhenWillHappen = command.WhenWillHappen,
            PhotoPath = imagePath,
            CreatedAt = DateTime.UtcNow
        };
    }

    private static CreateEventCommandResult PrepareSuccessResult(EventEntity eventEntity) => new(
            "Created.",
            eventEntity,
            201);
}