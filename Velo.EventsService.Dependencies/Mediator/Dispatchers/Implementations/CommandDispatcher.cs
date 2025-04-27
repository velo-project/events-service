using Microsoft.Extensions.DependencyInjection;
using Velo.EventsService.Dependencies.Mediator.Contracts;
using Velo.EventsService.Dependencies.Mediator.Handlers;

namespace Velo.EventsService.Dependencies.Mediator.Dispatchers.Implementations;

public class CommandDispatcher(IServiceProvider serviceProvider) : ICommandDispatcher
{
    public Task<TCommandResult> Dispatch<TCommand, TCommandResult>(TCommand command, CancellationToken cancellationToken) where TCommand : ICommand
    {
        var handler = serviceProvider.GetRequiredService<ICommandHandler<TCommand, TCommandResult>>();
        return handler.Handle(command, cancellationToken);
    }
}