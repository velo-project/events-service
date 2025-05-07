using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Velo.EventsService.Commands.UpdateImageFromEvent.Handler.Exceptions;
using Velo.EventsService.Dependencies.FileManager;
using Velo.EventsService.Dependencies.Mediator.Handlers;
using Velo.EventsService.Persistence.Contracts;

namespace Velo.EventsService.Commands.UpdateImageFromEvent.Handler
{
    public class UpdateImageFromEventCommandHandler(IFileService fileService, IEventsRepository repository) : ICommandHandler<UpdateImageFromEventCommand, UpdateImageFromEventCommandResult>
    {
        public async Task<UpdateImageFromEventCommandResult> Handle(UpdateImageFromEventCommand command, CancellationToken cancellationToken)
        {
            var path = await fileService.SavePhotoAsync(command.ImageBytes);

            if (path == null)
            {
                throw new FailedToUpdateImageException();
            }

            await repository.UpdateEventImageAsync(command.EventId, path, cancellationToken);

            return new UpdateImageFromEventCommandResult();
        }
    }
}
