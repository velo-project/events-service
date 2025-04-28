using System.Text.Json;
using Velo.EventsService.Dependencies.Exceptions;

namespace Velo.EventsService.Api.Middlewares;

public class HttpExceptionMiddleware
{
    private readonly RequestDelegate _next;

    public HttpExceptionMiddleware(RequestDelegate next)
    {
        _next = next;
    }

    public async Task Invoke(HttpContext context)
    {
        try
        {
            await _next(context);
        }
        catch (HttpException ex)
        {
            context.Response.StatusCode = ex.StatusCode;
            context.Response.ContentType = "application/json";

            var response = new
            {
                Message = ex.Message,
                StatusCode = ex.StatusCode
            };

            var jsonResponse = JsonSerializer.Serialize(response);
            await context.Response.WriteAsync(jsonResponse);
        }
    }
}