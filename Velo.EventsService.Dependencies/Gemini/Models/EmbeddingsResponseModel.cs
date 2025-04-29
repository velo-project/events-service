namespace Velo.EventsService.Dependencies.Gemini.Models;

public class EmbeddingsResponseModel
{
    public EmbeddingsResponse Embedding { get; set; } 
}

public class EmbeddingsResponse
{
    public List<float> Values { get; set; } = [];
}