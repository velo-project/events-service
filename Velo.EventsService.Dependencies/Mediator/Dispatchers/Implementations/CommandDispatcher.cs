using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using Velo.EventsService.Dependencies.Mediator.Contracts;
using Velo.EventsService.Dependencies.Mediator.Exceptions;
using Velo.EventsService.Dependencies.Mediator.Handlers;

namespace Velo.EventsService.Dependencies.Mediator.Dispatchers.Implementations;

public class CommandDispatcher(IServiceProvider serviceProvider) : ICommandDispatcher
{
    public async Task<TCommandResult> Dispatch<TCommand, TCommandResult>(TCommand command, CancellationToken cancellationToken) where TCommand : ICommand
    {
        var logger = serviceProvider.GetRequiredService<ILogger<TCommand>>();
        var handler = serviceProvider.GetRequiredService<ICommandHandler<TCommand, TCommandResult>>();

        if (handler == null)
            throw new NoCommandHandlerFoundException<TCommand, TCommandResult>();
        try
        {
            logger.LogInformation("Dispatching command: {CommandName}", typeof(TCommand).Name);
            var result = await handler.Handle(command, cancellationToken);
            logger.LogInformation("Dispatch success: {CommandName}", typeof(TCommand).Name);

            return result;
        }
        catch (Exception ex)
        {
            logger.LogError("Dispatch failed for {CommandName}", typeof(TCommand).Name);
            logger.LogError("Exception message: {Message}", ex.Message);
            logger.LogError("Stack trace: {StackTrace}", ex.StackTrace);

            if (ex.InnerException != null)
            {
                logger.LogError("Inner exception message: {InnerException}", ex.InnerException.Message);
            }
            
            throw;
        }
    }
}