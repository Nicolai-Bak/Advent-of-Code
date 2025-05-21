namespace Helpers;

public static class FileOperations
{
    public static IEnumerable<string> ReadLines(string path = "data.txt") => File.ReadLines(path);
}
