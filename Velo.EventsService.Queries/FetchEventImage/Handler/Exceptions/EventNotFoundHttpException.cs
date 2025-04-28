using Velo.EventsService.Dependencies.Exceptions;

namespace Velo.EventsService.Queries.FetchEventImage.Handler.Exceptions;

public class EventNotFoundHttpException(int id) : HttpException($"Event {id} not found", 404);