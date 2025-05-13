using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Velo.EventsService.Dependencies.Mediator.Contracts;

namespace Velo.EventsService.Commands.DeleteEvent
{
    public class DeleteEventCommand : ICommand
    {
        public int Id { get; set; }
    }
}
