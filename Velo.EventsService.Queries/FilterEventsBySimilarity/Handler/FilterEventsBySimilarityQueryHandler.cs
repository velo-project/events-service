using System.Globalization;
using Microsoft.EntityFrameworkCore;
using Pgvector;
using Velo.EventsService.Dependencies.Gemini;
using Velo.EventsService.Dependencies.Mediator.Handlers;
using Velo.EventsService.Persistence.Context;
using Velo.EventsService.Persistence.Entities;

namespace Velo.EventsService.Queries.FilterEventsBySimilarity.Handler;

public class FilterEventsBySimilarityQueryHandler : IQueryHandler<FilterEventsBySimilarityQuery, FilterEventsBySimilarityQueryResult>
{
    private readonly DbSet<EventEntity> _events;
    private readonly IGeminiService _geminiService;
    
    public FilterEventsBySimilarityQueryHandler(DatabaseContext context, IGeminiService geminiService)
    {
        _events = context.Set<EventEntity>();
        _geminiService = geminiService;
    }
    public async Task<FilterEventsBySimilarityQueryResult> Handle(FilterEventsBySimilarityQuery query, CancellationToken cancellationToken)
    {
        var embeddings = await _geminiService.GenerateEmbeddingsAsync(query.Text);

        var queryEmbedding = new Vector(embeddings.Embedding.Values.ToArray());
        
        var events = await _events.FromSqlInterpolated($@"
            SELECT *
            FROM ""tb_events""
            ORDER BY ""embeddings_event"" <=> {queryEmbedding}
            LIMIT 10
        ")
            .ToListAsync(cancellationToken);

        return new FilterEventsBySimilarityQueryResult()
        {
            Events = events
        };
    }
}