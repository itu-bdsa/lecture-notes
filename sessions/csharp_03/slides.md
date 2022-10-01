---
marp: true
author: Rasmus Lystrøm
theme: default
class: invert
---

# Survey during the *academic quarter*

## <https://www.menti.com/> code: `5942 1139`

---

# Questions

- How to setup tests?
- Testing private methods?
- How to use fluent assertions?

---

# C♯ 03: Lambdas and LINQ

![bg right:50%](images/lambda.png)

Rasmus Lystrøm
Associate Professor
ITU

---

# Topics

![bg right:60%](images/core.jpg)

Delegates
Anonymous methods
Lambda expressions
Local functions
Anonymous types
Tuples
Records
Extension methods
LINQ

---

![bg](images/delegate.jpg)

# Delegates

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

---

<!-- _class: default -->

# Delegates – Building block for Higher-order functions

```csharp
public delegate int BinaryOperation(int x, int y);

var add = new BinaryOperation(
    delegate (int x, int y)
    {
        return x + y;
    }
);
```

---

# Delegates

## Demo

---

# Lambda Expressions

![bg right:50%](images/lambda.png)

---

<!-- _class: default -->

# Lambda Expressions

```csharp
Action<string> write = s => Console.WriteLine(s);

Func<int, int> square = a => a * a;
```

```csharp
Predicate<City> b = c => c.Name.StartsWith("B");

Converter<double, double> ftoC = c => c * 9.0 / 5.0 + 32.0;

...
```

---

<!-- _class: default -->

# Local Functions

```csharp
static void Main(string[] args)
{
    int square(int a) { return a * a; };

    Console.WriteLine(square(16));
}
```

---

# Local Functions

## Demo

---

![bg](images/belt.png)

# Prerequisites

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

---

<!-- _class: default -->

# Anonymous types

```csharp
var question = new
{
    Title = "The answer...?",
    Answer = 42
};
```

---

<!-- _class: default -->

# (Tuples)

```csharp
var s = Tuple.Create("Clark Kent", "Superman");

var b = ("Bruce Wayne", "Batman");

var f = (name: "Barry Allen", alterEgo: "The Flash");

var random = new Random();

IEnumerable<(float x, float y)> GenerateCoordinates()
{
    yield return (random.NextSingle() * 100, random.NextSingle() * 100);
}
```

---

<!-- _class: default -->

# Records

```csharp
public record Superhero(string Name, string AlterEgo, DateTime FirstAppearance);
```

---

<!-- _class: default -->

# Data: Collection Initializer

```csharp
IEnumerable<City> cities = new[]
{
    new City(1, "Berlin"),
    new City(2, "Hamburg"),
    new City(3, "Frankfurt")
};
```

---

<!-- _class: default -->

# Data: Collection + Object Initializer

```csharp
IEnumerable<City> cities = new[]
{
    new City { Id = 1, Name = "Berlin" },
    new City { Id = 2, Name = "Hamburg" },
    new City { Id = 3, Name = "Frankfurt" }
};
```

---

<!-- _class: default -->

![bg](images/extensions.jpg)

# Extension Methods

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

---

<!-- _class: default -->

# Built-in Extension methods

```csharp
var count = cities.Count();

var sort = cities.OrderBy(c => c.Name);

var filter = cities.Where(c => c.Name.Contains("i"));

var pick = cities.FirstOrDefault(c => c.Id == 2);

var all = cities.All(c => c.Name.Length < 10);

var any = cities.Any(c => c.Name.StartsWith("B"));

var project = cities.Select(c => c.Name);
```

---

<!-- _class: default -->

# Create your own extension method

```csharp
public static class Extensions
{
    public static int WordCount(this string str) =>
        str.Split(new[] { ' ', '.', '?' },
                  StringSplitOptions.RemoveEmptyEntries)
           .Length;
}
```

---

# Extension methods

## Demo

---

![bg](images/questions.jpg)

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

# LINQ - Language INtegrated Query

---

<!-- _class: default -->

# LINQ

<!-- _class: default -->

```csharp
var sorted = from c in cities
             where c.Name.Contains("i")
             orderby c.Name descending
             select new { Name = c.Name };
```

## Extension Methods Version

```csharp
var sorted = cities.Where(c => c.Name.Contains("i"))
                   .OrderByDescending(c => c.Name)
                   .Select(c => new { Name = c.Name });
```

---

# LINQ

## Demo

---

![bg](images/thank-you.jpg)

# Thank you
