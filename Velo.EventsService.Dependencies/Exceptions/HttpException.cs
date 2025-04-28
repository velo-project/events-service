namespace Velo.EventsService.Dependencies.Exceptions;

public class HttpException(string message, int statusCode) : Exception(message)
{
    public int StatusCode { get; set; } = statusCode;
}