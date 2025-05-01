using Microsoft.Extensions.Configuration;
using Velo.EventsService.Dependencies.FileManager.Exceptions;

namespace Velo.EventsService.Dependencies.FileManager.Services;

public class FileService : IFileService
{
    private readonly string _imageFolderPath;

    public FileService(IConfiguration configuration)
    {
        _imageFolderPath = configuration["ImageFolderPath"] ?? throw new NoImageFolderPathFoundException();
    }
    
    public Task<string?> SavePhotoAsync(byte[] file)
    {
        return SaveFileInternalLogicAsync(file, ".png");
    }

    private async Task<string?> SaveFileInternalLogicAsync(byte[] file, string format)
    {
        if (file.Length == 0)
            return null;

        var folderCombination = $"{DateTime.UtcNow.Year}/{DateTime.UtcNow.Month}/{DateTime.UtcNow.Day}";
        var folderPath = Path.Combine(_imageFolderPath, folderCombination);
        if (!Directory.Exists(folderPath))
            Directory.CreateDirectory(folderPath);

        var fileName = Guid.NewGuid().ToString().Replace("-", "") + format;
        var filePath = Path.Combine(folderPath, fileName);

        await File.WriteAllBytesAsync(filePath, file);

        return filePath;
    }
}