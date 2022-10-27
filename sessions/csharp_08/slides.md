---
marp: true
author: Rasmus Lystrøm
theme: default
class: invert
---

# C♯ 08: JSON and the REST part deux

![bg right](https://i0.wp.com/farm1.staticflickr.com/207/476245464_2b59cc08f2_z.jpg)

Rasmus Lystrøm
Associate Professor
ITU

---

# Documenting ASP&period;NET Core

## Demo

---

<!-- _class: default -->

# Support the OpenAPI Specification (Swagger)

```bash
dotnet add package Swashbuckle.AspNetCore
```

```csharp
builder.Services.AddSwaggerGen();

var app = builder.Build();

if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}
```

---

<!-- _class: default -->

# Support HTTPS only

Cf. [Enforce HTTPS in ASP.NET Core](https://learn.microsoft.com/en-us/aspnet/core/security/enforcing-ssl)

```bash
dotnet dev-certs https --trust
```

```csharp
app.UseHttpsRedirection();
```

---

<!-- _class: default -->

# Standardize on lowercase urls

```csharp
builder.Services.AddRouting(options => options.LowercaseUrls = true);
```

---

# Secure your Web API

Azure AD
Azure AD B2C
3rd party

**Do not write your own security layer!!!**

---

![bg](https://upload.wikimedia.org/wikipedia/commons/2/2a/CSIRO_ScienceImage_2819_Preparing_samples_of_soy_sauce_for_testing.jpg)

<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>

# Testing ASP&period;NET Core

---

# Stub --> Mock --> Fake

---

# Stubs

```csharp
public class FooServiceFalseStub : IFooService
{
    public bool Bar(Foo foo)
    {
        return false;
    }
}
```

---

# Test using stub

```csharp

[Fact]
public void Exec_when_IFooService_Update_returns_false_returns_false()
{
    IFooService service = new FooServiceTrueStub();
    var baz = new Baz(service);
    var result = baz.Exec(new Foo());
    result.Should().Be(false);
}
```

---

# Faking it using Moq

```bash
dotnet add package Moq
```

```csharp
[Fact]
public void Exec_when_IFooService_Update_returns_false_returns_false()
{
    var mock = new Mock<IFooService>();
    IFooService service = mock.Object;
    var baz = new Baz(service);
    
    var result = baz.Exec(new Foo());
    result.Should().Be(false);
}
```

---

# Faking it using Moq (advanced)

```csharp
[Fact]
public void Exec_when_IFooService_Update_true_returns_true()
{
    var mock = new Mock<IFooService>();
    mock.Setup(m => m.Update(It.IsAny<Foo>())).Returns(true);
    var baz = new Baz(mock.Object);
    
    var result = baz.Exec(new Foo());
    result.Should().Be(true);
}
```

---

# Faking it using NSubstitute

```bash
dotnet add package NSubstitute
```

```csharp
[Fact]
public void Exec_when_IFooService_Update_returns_false_returns_false()
{
    var service = Substitute.For<IFooService>();
    var baz = new Baz(service);
    
    var result = baz.Exec(new Foo());
    result.Should().Be(false);
}
```

---

# Faking it using NSubstitute (advanced)

```csharp
[Fact]
public void Exec_when_IFooService_Update_true_returns_true()
{
    var service = Substitute.For<IFooService>();
    service.Update(Arg.Any<Foo>()).Returns(true);
    var baz = new Baz(service);
    
    var result = baz.Exec(new Foo());
    result.Should().Be(true);
}
```

---

# Testing ASP&period;NET Core with fakes

## Demo

---

# [ASP&period;NET Minimal APIs](https://learn.microsoft.com/en-us/aspnet/core/fundamentals/minimal-apis)

`dotnet new webapi -minimal`

*Controllers are so yesterday...*

## Demo

---

# [Integration tests in ASP&period;NET Core](https://learn.microsoft.com/en-us/aspnet/core/test/integration-tests?view=aspnetcore-6.0)

## Demo

---

# gRPC

> gRPC is a modern, open source, high-performance remote procedure call (RPC) framework that can run anywhere. gRPC enables client and server applications to communicate transparently, and simplifies the building of connected systems.

Source: <https://github.com/grpc/grpc>

```bash
dotnet new grpc
```

---

# GraphQL

> GraphQL is a query language for APIs and a runtime for fulfilling those queries with your existing data. GraphQL provides a complete and understandable description of the data in your API, gives clients the power to ask for exactly what they need and nothing more, makes it easier to evolve APIs over time, and enables powerful developer tools.

Source: <https://graphql.org/>

---

# Thank You
