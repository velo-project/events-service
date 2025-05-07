using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Velo.EventsService.Dependencies.Exceptions;

namespace Velo.EventsService.Persistence.Repositories.Exceptions
{
    public class EntityNotFoundException() : HttpException("No entity found", 404)
    {
    }
}
