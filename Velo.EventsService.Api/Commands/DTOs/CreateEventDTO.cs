using System.ComponentModel.DataAnnotations;
using Velo.EventsService.Commands.CreateEvent;

namespace Velo.EventsService.Api.Commands.DTOs;

public class CreateEventDTO
{
    [Required]
    public string Name { get; set; } = string.Empty;
    public string? Description { get; set; }
    
    [Required]
    public DateTime WhenWillHappen { get; set; }
    
    [Required]
    public IFormFile File { get; set; }
}