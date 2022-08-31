using System.IO;
using System.Reflection;

using FluentAssertions;
using Xunit;

namespace HelloWorld.Tests;

public class ProgramTests
{
    [Fact]
    public void Main_given_no_args_prints_Hello_World()
    {
        // Arrange
        using var writer = new StringWriter();
        Console.SetOut(writer);

        // Act
        var program = Assembly.Load(nameof(HelloWorld));
        program.EntryPoint?.Invoke(null, new[] { Array.Empty<string>() });

        // Assert
        var output = writer.GetStringBuilder().ToString().TrimEnd();
        output.Should().Be("Hello, World!");
    }
}