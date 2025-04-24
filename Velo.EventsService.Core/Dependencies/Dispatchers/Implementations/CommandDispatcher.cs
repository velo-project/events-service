using Microsoft.Extensions.DependencyInjection;
using Velo.EventsService.Core.Dependencies.Contracts;
using Velo.EventsService.Core.Dependencies.Handlers;

namespace Velo.EventsService.Core.Dependencies.Dispatchers.Implementations;

public class CommandDispatcher(IServiceProvider serviceProvider) : ICommandDispatcher
{
    public Task<TCommandResult> Dispatch<TCommand, TCommandResult>(TCommand command, CancellationToken cancellationToken) where TCommand : ICommand
    {
        var handler = serviceProvider.GetRequiredService<ICommandHandler<TCommand, TCommandResult>>();
        return handler.Handle(command, cancellationToken);
    }
}