using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Velo.EventsService.Dependencies.Mediator.Handlers;
using Velo.EventsService.Persistence.Contracts;

namespace Velo.EventsService.Commands.DeleteEvent.Handler
{
    public class DeleteEventCommandHandler(IEventsRepository repository) : ICommandHandler<DeleteEventCommand, DeleteEventCommandResult>
    {
        public async Task<DeleteEventCommandResult> Handle(DeleteEventCommand command, CancellationToken cancellationToken)
        {
            await repository.DeleteEventAsync(command.Id, cancellationToken);

            return new DeleteEventCommandResult();
        }
    }
}
