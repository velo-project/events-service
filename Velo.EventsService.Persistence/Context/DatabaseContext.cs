using Microsoft.EntityFrameworkCore;
using Velo.EventsService.Core.Entities;

namespace Velo.EventsService.Persistence.Context;

public class DatabaseContext(DbContextOptions options) : DbContext(options)
{
    public DbSet<EventEntity> Events { get; protected set; }
}