using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace Velo.EventsService.Core.Entities;

public class EventEntity
{
    [Key]
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    [Column("id_event")]
    public int Id { get; set; }

    [Column("name_event")]
    public string Name { get; set; } = string.Empty;
    
    [Column("description_event")]
    public string Description { get; set; } = string.Empty;

    [Column("photo_event")]
    public string? PhotoPath { get; set; }
    
    [Column("active_event")]
    public bool IsActive { get; set; }
    
    [Column("canceled_event")]
    public bool IsCanceled { get; set; }
    
    [Column("when_will_happen_event")]
    public DateTime WhenWillHappen { get; set; }
    
    [Column("created_at")]
    public DateTime CreatedAt { get; set; }
}