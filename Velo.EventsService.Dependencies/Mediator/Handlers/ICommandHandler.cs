using Velo.EventsService.Dependencies.Mediator.Contracts;

namespace Velo.EventsService.Dependencies.Mediator.Handlers;

public interface ICommandHandler<in TCommand, TCommandResult> where TCommand : ICommand
{
    Task<TCommandResult> Handle(TCommand command, CancellationToken cancellationToken);
}