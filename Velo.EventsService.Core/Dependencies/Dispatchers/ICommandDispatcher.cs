using Velo.EventsService.Core.Dependencies.Contracts;

namespace Velo.EventsService.Core.Dependencies.Dispatchers;

public interface ICommandDispatcher
{
    Task<TCommandResult> Dispatch<TCommand, TCommandResult>(TCommand command, CancellationToken cancellationToken)
        where TCommand : ICommand;
}