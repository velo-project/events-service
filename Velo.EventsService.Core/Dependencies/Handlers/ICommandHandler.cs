using Velo.EventsService.Core.Dependencies.Contracts;

namespace Velo.EventsService.Core.Dependencies.Handlers;

public interface ICommandHandler<in TCommand, TCommandResult> where TCommand : ICommand
{
    Task<TCommandResult> Handle(TCommand command, CancellationToken cancellationToken);
}