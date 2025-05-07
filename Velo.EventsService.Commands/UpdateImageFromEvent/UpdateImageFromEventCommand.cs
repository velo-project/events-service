using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Velo.EventsService.Dependencies.Mediator.Contracts;

namespace Velo.EventsService.Commands.UpdateImageFromEvent
{
    public class UpdateImageFromEventCommand : ICommand
    {
        public byte[] ImageBytes { get; set; } = [];
        public int EventId { get; set; }
    }
}
