
using Microsoft.Extensions.Configuration;
using Pgvector;
using Velo.EventsService.Dependencies.FileManager;
using Velo.EventsService.Dependencies.Gemini;
using Velo.EventsService.Dependencies.Mediator.Handlers;
using Velo.EventsService.Persistence.Contracts;
using Velo.EventsService.Persistence.Entities;

namespace Velo.EventsService.Commands.CreateEvent.Handler;

public class CreateEventCommandHandler(IEventsRepository eventsRepository, IGeminiService geminiService, IFileService fileService)
    : ICommandHandler<CreateEventCommand, CreateEventCommandResult>
{
    
    public async Task<CreateEventCommandResult> Handle(CreateEventCommand command, CancellationToken cancellationToken)
    {
        cancellationToken.ThrowIfCancellationRequested();

        var imagePath = await ProcessAndSaveImage(command);
        var eventToSave = ExtractAndPrepareEventFrom(command, imagePath);
        
        eventToSave = await PrepareAndGenerateEmbeddings(eventToSave);
        eventToSave = await eventsRepository.PersistEventAsync(eventToSave, cancellationToken);

        return PrepareSuccessResult(eventToSave);
    }

    private async Task<EventEntity> PrepareAndGenerateEmbeddings(EventEntity eventEntity)
    {
        var embeddingsToGenerate = $"{eventEntity.Name} {eventEntity.Description}";
        var response = await geminiService.GenerateEmbeddingsAsync(embeddingsToGenerate);
        eventEntity.Embeddings = new Vector(response.Embedding.Values.ToArray());
        return eventEntity;
    }

    private Task<string?> ProcessAndSaveImage(CreateEventCommand command)
    {
        return fileService.SavePhotoAsync(command.PhotoBytes);
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