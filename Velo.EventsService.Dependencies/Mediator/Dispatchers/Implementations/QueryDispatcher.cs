using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using Velo.EventsService.Dependencies.Mediator.Contracts;
using Velo.EventsService.Dependencies.Mediator.Handlers;

namespace Velo.EventsService.Dependencies.Mediator.Dispatchers.Implementations;

public class QueryDispatcher(IServiceProvider serviceProvider) : IQueryDispatcher
{
    public async Task<TQueryResult> Dispatch<TQuery, TQueryResult>(TQuery query, CancellationToken cancellationToken) where TQuery : IQuery
    {
        var logger = serviceProvider.GetRequiredService<ILogger<TQuery>>();
        var handler = serviceProvider.GetRequiredService<IQueryHandler<TQuery, TQueryResult>>();

        if (handler == null)
        {
            var message = $"No Handler Found for <{typeof(TQuery).Namespace}, {typeof(TQueryResult).Namespace}>";
            logger.LogError(message);
            throw new Exception(message);
        }

        try
        {
            logger.LogInformation("Dispatching query: {QueryName}", typeof(TQuery).Name);
            var result = await handler.Handle(query, cancellationToken);
            logger.LogInformation("Dispatch success: {QueryName}", typeof(TQuery).Name);

            return result;
        }
        catch (Exception ex)
        {
            logger.LogError("Dispatch failed for {QueryName}", typeof(TQuery).Name);
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