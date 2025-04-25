using Microsoft.Extensions.DependencyInjection;
using Velo.EventsService.Persistence.Dependencies.Contracts;
using Velo.EventsService.Persistence.Dependencies.Dispatchers;
using Velo.EventsService.Persistence.Dependencies.Handlers;

namespace Velo.EventsService.Core.Dependencies.Dispatchers.Implementations;

public class CommandDispatcher(IServiceProvider serviceProvider) : ICommandDispatcher
{
    public Task<TCommandResult> Dispatch<TCommand, TCommandResult>(TCommand command, CancellationToken cancellationToken) where TCommand : ICommand
    {
        var handler = serviceProvider.GetRequiredService<ICommandHandler<TCommand, TCommandResult>>();
        return handler.Handle(command, cancellationToken);
    }
}