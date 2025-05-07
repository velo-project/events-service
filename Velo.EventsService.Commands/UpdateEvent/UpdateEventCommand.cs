using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Velo.EventsService.Dependencies.Mediator.Contracts;
using Velo.EventsService.Persistence.Entities;

namespace Velo.EventsService.Commands.UpdateEvent
{
    public class UpdateEventCommand : ICommand
    {
        public int EventId { get; set; }
        public required EventEntity Event { get; set; }
    }
}
