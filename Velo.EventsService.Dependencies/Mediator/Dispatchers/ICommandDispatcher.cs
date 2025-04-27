using Velo.EventsService.Dependencies.Mediator.Contracts;

namespace Velo.EventsService.Dependencies.Mediator.Dispatchers;

public interface ICommandDispatcher
{
    Task<TCommandResult> Dispatch<TCommand, TCommandResult>(TCommand command, CancellationToken cancellationToken)
        where TCommand : ICommand;
}