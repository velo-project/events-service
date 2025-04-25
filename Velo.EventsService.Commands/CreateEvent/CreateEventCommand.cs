using System.ComponentModel.DataAnnotations;
using Velo.EventsService.Persistence.Dependencies.Contracts;

namespace Velo.EventsService.Commands.CreateEvent;

public class CreateEventCommand : ICommand
{
    [Required]
    public string Name { get; set; } = string.Empty;
    public string? Description { get; set; }
    
    [Required]
    public DateTime WhenWillHappen { get; set; }
}