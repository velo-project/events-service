using Microsoft.EntityFrameworkCore;
using Velo.EventsService.Persistence.Entities;

namespace Velo.EventsService.Persistence.Context;

public class DatabaseContext(DbContextOptions options) : DbContext(options)
{
    public DbSet<EventEntity> Events { get; protected set; }

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.HasPostgresExtension("vector");
        modelBuilder.Entity<EventEntity>(builder =>
        {
            builder.Property(_ => _.Embeddings).IsRequired(false);
            builder.HasIndex(_ => _.Embeddings)
                .HasMethod("hnsw")
                .HasOperators("vector_cosine_ops");
        });
        
        base.OnModelCreating(modelBuilder);
    }
}