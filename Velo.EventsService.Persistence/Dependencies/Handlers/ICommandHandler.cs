using Velo.EventsService.Persistence.Dependencies.Contracts;

namespace Velo.EventsService.Persistence.Dependencies.Handlers;

public interface ICommandHandler<in TCommand, TCommandResult> where TCommand : ICommand
{
    Task<TCommandResult> Handle(TCommand command, CancellationToken cancellationToken);
}