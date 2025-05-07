using System.ComponentModel.DataAnnotations;

namespace Velo.EventsService.Api.Commands.DTOs
{
    public class UpdateEventDTO
    {
        [Required]
        public string Name { get; set; } = string.Empty;

        [Required]
        public string? Description { get; set; }

        [Required]
        public bool IsCanceled { get; set; }

        [Required]
        public DateTime WhenWillHappen { get; set; }
    }

}
