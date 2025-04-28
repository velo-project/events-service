using Velo.EventsService.Dependencies.Exceptions;

namespace Velo.EventsService.Queries.FetchEventImage.Handler.Exceptions;

public class EventWithNoImageHttpException() : HttpException("This event dont have an image", 400);