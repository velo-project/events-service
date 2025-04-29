using Velo.EventsService.Dependencies.Gemini.Models;

namespace Velo.EventsService.Dependencies.Gemini;

public interface IGeminiService
{
    Task<EmbeddingsResponseModel> GenerateEmbeddingsAsync(string text);
}