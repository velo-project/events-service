using System.Text.Json;
using Microsoft.Extensions.Configuration;
using Velo.EventsService.Dependencies.Gemini.Models;

namespace Velo.EventsService.Dependencies.Gemini.Services;

public class GeminiService : IGeminiService
{
    private readonly string _geminiApiKey;
    private readonly string _geminiModel;
    private readonly Uri _geminiApiCurl;
    private readonly HttpClient _httpClient;

    public GeminiService(IConfiguration configuration, HttpClient httpClient)
    {
        var apiKey = configuration["Gemini:ApiKey"] ?? throw new Exception("Gemini:ApiKey Is Null");
        var model = configuration["Gemini:EmbeddingsModel"] ?? throw new Exception("EmbeddingsModel Is Null");

        _geminiApiKey = apiKey;
        _geminiModel = $"models/{model}";
        _httpClient = httpClient;
        _geminiApiCurl = new Uri($"https://generativelanguage.googleapis.com/v1beta/models/gemini-embedding-exp-03-07:embedContent?key={_geminiApiKey}");
    }
    
    public async Task<EmbeddingsResponseModel> GenerateEmbeddingsAsync(string text)
    {
        var body = new
        {
            Model = _geminiModel,
            Parts = new[]
            {
                new { Text = text }
            },
            TaskType = "SEMANTIC_SIMILARITY"
        };
        var jsonBody = JsonSerializer.Serialize(body);
        var content = new StringContent(jsonBody);
        var response = await _httpClient.PostAsync(_geminiApiCurl, content);

        response.EnsureSuccessStatusCode();

        var responseContent = await response.Content.ReadAsStringAsync();
        return JsonSerializer.Deserialize<EmbeddingsResponseModel>(responseContent) ?? new EmbeddingsResponseModel();
    }
}