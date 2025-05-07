using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Velo.EventsService.Dependencies.Exceptions;

namespace Velo.EventsService.Commands.UpdateImageFromEvent.Handler.Exceptions
{
    public class FailedToUpdateImageException() : HttpException("Failed to update image", 500)
    {
    }
}
