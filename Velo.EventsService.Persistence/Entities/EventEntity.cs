using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text.Json.Serialization;
using Pgvector;

namespace Velo.EventsService.Persistence.Entities;

[Table("tb_events")]
public class EventEntity
{
    [Key]
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    [Column("id_event")]
    public int Id { get; set; }

    [Column("name_event")]
    public string Name { get; set; } = string.Empty;
    
    [Column("description_event")]
    public string? Description { get; set; }

    [Column("photo_event")]
    [JsonIgnore]
    public string? PhotoPath { get; set; }
    
    [Column("active_event")]
    public bool IsActive { get; set; }
    
    [Column("canceled_event")]
    public bool IsCanceled { get; set; }

    [JsonIgnore]
    [Column("embeddings_event", TypeName = "vector(3072)")]
    public Vector? Embeddings { get; set; }
    
    [Column("when_will_happen_event")]
    public DateTime WhenWillHappen { get; set; }
    
    [Column("created_at")]
    public DateTime CreatedAt { get; set; }
}