---
marp: true
author: Rasmus Lystrøm
theme: default
class: invert
---

# C♯ 5 - Dependency Injection and Testing Entity Framework Core

![bg right:60%](https://samueleresca.net/content/images/wordpress/2017/02/How-To-Use-Dependency-Injection-in-ASP.NET-MVC-6.jpg)

Rasmus Lystrøm
Associate Professor
ITU

---

![bg](https://miro.medium.com/max/1400/1*wgpae8kWxzQPG0XBAzUWYw.jpeg)

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
<br>
<br>

# Repository pattern

---

# Repository Pattern

- Enable CRUD on domain objects (entities)

- Usually: one repository per entity

- Debatable: has a `Save()` method

---

# Generic repository

```csharp
public interface Repository<T, K>
{
    T Create(T entity);
    IReadOnlyCollection<T> Read();
    T Read(K id);
    void Update(T entity);
    void Delete(K id);
}
```

---

# Repository Pattern

> ... but wait ...Entity Framework already does that for me!?

---

# Recommended Repository: Per entity e.g., `Character`

```csharp
public interface ICharacterRepository
{
    int Create(CharacterCreateDTO character);
    CharacterDetailsDTO Read(int characterId);
    IReadOnlyCollection<CharacterDTO> Read();
    void Update(CharacterUpdateDTO character);
    void Delete(int characterId);
}
```

...or something similar...

---

# Testing...

- Testing live databases is hard
- Testing live full systems is hard
- By transitivity: Testing ... is hard...

---

<!-- _class: default -->

![bg contain](https://samueleresca.net/content/images/wordpress/2017/02/How-To-Use-Dependency-Injection-in-ASP.NET-MVC-6.jpg)

---

# Dependency Injection

Software design pattern which implements Inversion of Control (IoC)

- Constructor Injection
- Property (Setter) Injection
- Interface Injection

---

# Dependency Injection

- Structured readable code
- Testable code
- Dependency Inversion Principle
- Separation of Concerns
- Rock SOLI***D***!!! <-- Pun intented
- AWESOME!!

---

# Programming to interface, not implementation…

```csharp
public interface IFooService
{
    bool Bar(Foo foo);
}

public class FooService : IFooService
{
    bool Bar(Foo foo)
    {
         // Implementation
    }
}

public interface IFooMapper
{
    Foo Map(Qux qux);
}
```

---

![bg right](https://cdn.shopify.com/s/files/1/1155/3572/articles/Untitled_design_71_1800x.png?v=1547597527)

# Using IFooService?

```csharp
public class Baz
{
    public bool Grault(Qux qux)
    {
        IFooMapper mapper = new FooMapper();
        var foo = mapper.Map(qux);
        IFooService service = new FooService();
        return service.Bar(foo);
    }
}
```

---

# Constructor Injection (preferred)

```csharp
public class Baz
{
    private readonly IFooMapper _mapper;
    private readonly IFooService _service;

    public Baz(IFooMapper mapper,  IFooService service)
    {
        _mapper = mapper;
        _service = service;
    }

    public bool Grault(Qux qux)
    {
        var foo = _mapper.Map(qux);
        return _service.Bar(foo);
    }
}
```

---

# <strike>Property Injection</strike>

```csharp
public class Baz
{
    public IFooService Service { private get; set; }

    public bool Grault(Qux qux)
    {
        ...

        return Service?.Update(foo);
    }
}
```

---

# <strike>Interface Injection</strike>

```csharp
public interface IServiceSetter<T>
{
    T Service { set; }
}

public class Baz : IServiceSetter<IFooService>
{
    public IFooService Service { private get; set; }

    public bool Grault(Qux qux)
    {
        ...

        return Service?.Update(foo);
    }
}
```

---

# Best practices

- Use Adapter to enable interface if needed
- Use constructor injection
- Program to interface
- Use an IoC container (more on that later)

---

# Testing Entity Framework

---

# SQLite in-memory database

```bash
dotnet add package Microsoft.EntityFrameworkCore.Sqlite
```

```csharp
using var connection = new SqliteConnection("Filename=:memory:");
connection.Open();
var builder = new DbContextOptionsBuilder<MyContext>().UseSqlite(connection);
using var context = new MyContext(builder.Options);
```

---

# Demo

## Testing Entity Framework with *SQLite in-memory*

---

# Best practices

- Wrap in logical units/service classes/repositories
- Don’t test built-in code…
- Program to interface
- Repositories should not depend on other repositories
- Use integration testing to ensure it all works

---

# Integration testing

- Most test against a real database eventually
- More on this later in the course

---

# Thank you
