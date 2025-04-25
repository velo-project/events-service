using Velo.EventsService.Persistence.Dependencies.Contracts;

namespace Velo.EventsService.Persistence.Dependencies.Dispatchers;

public interface ICommandDispatcher
{
    Task<TCommandResult> Dispatch<TCommand, TCommandResult>(TCommand command, CancellationToken cancellationToken)
        where TCommand : ICommand;
}