namespace Velo.EventsService.Dependencies.FileManager;

public interface IFileService
{
    /// <summary>
    /// Save photo (on .png format) on ${ImageFolderPath}/${Year}/${Month}/${Day}
    /// </summary>
    /// <param name="file">The bytes from the file</param>
    /// <returns>The full path of new archive</returns>
    Task<string?> SavePhotoAsync(byte[] file);
}