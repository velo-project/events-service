using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Pgvector;
using Velo.EventsService.Dependencies.Gemini;
using Velo.EventsService.Dependencies.Mediator.Handlers;
using Velo.EventsService.Persistence.Contracts;

namespace Velo.EventsService.Commands.UpdateEvent.Handler
{
    public class UpdateEventCommandHandler(IEventsRepository repository, IGeminiService geminiService) : ICommandHandler<UpdateEventCommand, UpdateEventCommandResult>
    {
        public async Task<UpdateEventCommandResult> Handle(UpdateEventCommand command, CancellationToken cancellationToken)
        {
            var embeddingsToGenerate = $"{command.Event.Name} {command.Event.Description}";
            var embeddings = await geminiService.GenerateEmbeddingsAsync(embeddingsToGenerate);

            command.Event.Embeddings = new Vector(embeddings.Embedding.Values.ToArray());

            await repository.UpdateEventAsync(command.EventId, command.Event, cancellationToken);

            return new UpdateEventCommandResult();
        }
    }
}
