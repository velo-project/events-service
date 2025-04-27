using Velo.EventsService.Dependencies.Mediator.Contracts;

namespace Velo.EventsService.Dependencies.Mediator.Exceptions;

public class NoCommandHandlerFoundException<TCommand, TCommandResult>()
    : MediatorException($"No CommandHandler Found for <{typeof(TCommand).Namespace}, {typeof(TCommandResult).Namespace}>")
    where TCommand : ICommand;
    
public class NoQueryHandlerFoundException<TQuery, TQueryResult>()
    : MediatorException($"No QueryHandler Found for <{typeof(TQuery).Namespace}, {typeof(TQueryResult).Namespace}>")
    where TQuery : IQuery;