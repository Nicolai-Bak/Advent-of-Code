using Helpers;

Console.WriteLine("Hello, World!");


var firstLine = FileOperations.ReadLines()
.FirstOrDefault();

Console.WriteLine(firstLine);


