using System.ComponentModel.DataAnnotations;
using Velo.EventsService.Dependencies.Mediator.Contracts;

namespace Velo.EventsService.Commands.CreateEvent;

public class CreateEventCommand : ICommand
{
    public string Name { get; set; } = string.Empty;
    public string? Description { get; set; }
    public DateTime WhenWillHappen { get; set; }
    public byte[] PhotoBytes { get; set; } = [];
}