---
marp: true
author: Rasmus Lystrøm
theme: default
class: invert
---

# C♯ 10: .NET Mobile and Desktop Applications

![bg right](https://upload.wikimedia.org/wikipedia/commons/f/f4/Ordi-portable-milouf-img_0999.jpg)

Rasmus Lystrøm
Associate Professor
ITU

---

<!-- _class: default -->

![height:300px](images/ali-farahnak.jpg)

# Guest Lecture

## Mobile Application Development with C♯

Ali Reza Farahnak
Senior Cloud Solution Architect - Engineering

---

![bg](https://enginuitypowersystems.com/wp-content/uploads/2021/04/Time-to-Upgrade-700x467.jpg)

---

# Upgrading .NET Applications

Use Dependabot: <https://docs.github.com/en/code-security/dependabot>.

Get the latest .NET SDK: <https://dot.net/>.

```bash
dotnet tool install --global upgrade-assistant

dotnet tool install --global try-convert

upgrade-assistant analyze <csproj/sln>

upgrade-assistant upgrade <csproj/sln> --skip-backup

dotnet tool list --global

dotnet tool update --global dotnet-ef

dotnet list package --outdated
```

---

<!-- _class: default -->

# Update your .csproj files

```xml
<Project Sdk="Microsoft.NET.Sdk">
  <PropertyGroup>
    <TargetFramework>net7.0</TargetFramework>
    <Nullable>enable</Nullable>
    <ImplicitUsings>enable</ImplicitUsings>
  </PropertyGroup>
</Project>
```

---

# Upgrading .NET Applications

## Demo

---

# C♯ 11

## New features

- [UTF-8 string literals](https://learn.microsoft.com/en-us/dotnet/csharp/whats-new/csharp-11#utf-8-string-literals)
- [Required members](https://learn.microsoft.com/en-us/dotnet/csharp/whats-new/csharp-11#required-members)
- Other stuff: [What's new in C♯ 11](https://learn.microsoft.com/en-us/dotnet/csharp/whats-new/csharp-11)

Check out what's new in C♯ 8.0, 9.0, and 10.0 also.

---

# .NET 7

## New features

- [What's new in .NET 7](https://learn.microsoft.com/en-us/dotnet/core/whats-new/dotnet-7)
- [What's new in ASP.NET Core 7.0](https://learn.microsoft.com/en-us/ef/core/what-is-new/ef-core-7.0/whatsnew)
- [What's New in EF Core 7.0](https://learn.microsoft.com/en-us/ef/core/what-is-new/ef-core-7.0/whatsnew)

Check out what's new in C♯ 5.0 and 6.0.

---

# Mobile and Desktop Applications in .NET

Windows Forms (WinForms) (2002)

Windows Presentation Foundation (WPF) (2006)

Silverlight (2007-2019)

Xamarin.Forms (2014) (acquired by Microsoft in 2016)

Universal Windows Platform (UWP) (2015-2021?)

Windows UI Library (WinUI) (2021)

.NET Multi-platform App UI (MAUI) (2021)

---

# XAML: eXtensible Application Markup Language

Windows Desktop (WPF)
.NET MAUI
WinUI

![](images/rip.gif)

Xamarin.Forms (iOS, Android, Windows)
Silverlight (web)
Universal Windows Platform (anything windowsy)

---

# XAML

Markup language for declaratively designing and creating application UIs

XAML maps XML markup to objects in the .NET Framework

Every tag maps to a class and every attribute to a property

Markup and procedural code are peers in functionality and performance

Code and markup are both first class citizens

Consistent model between UI, documents, and media

Compiled to code

---

<!-- _class: default -->

# XAML Markup vs. Code

```xml
<Button Width="100">OK
    <Button.Background>
        Purple
    </Button.Background>
</Button>
```

![height:100px](images/button.png)

```csharp
var button = new Button();
button.Content = "OK";
button.Background = new SolidColorBrush(Colors.Purple);
button.Width = 100;
```

---

<!-- _class: default -->

# MainPage.xaml

```xml
<Page>
    <Grid>
        <StackPanel>
            <Ellipse Name="Light" Fill="Red"
                     Height="200" Width="200" Margin="50" />
            <Button Width="150"
                    Content="Change Lights"
                    HorizontalAlignment="Center"
                    Click="Button_Click" />
        </StackPanel>
    </Grid>
</Page>
```

---

<!-- _class: default -->

# MainPage.xaml.cs

```csharp
namespace App;

public sealed partial class MainPage : Page
{
    public MainPage()
    {
        this.InitializeComponent();
    }

    private void Button_Click(object sender, RoutedEventArgs e)
    {
        var current = Light.Fill as SolidColorBrush;

        if (current.Color == Colors.Red)
            Light.Fill = new SolidColorBrush(Colors.Green);
        else
            Light.Fill = new SolidColorBrush(Colors.Red);
    }
}
```

---

# Desktop Applications

## Demo

---

<!-- _class: default -->

# XAML + Code Behind

```xml
<StackPanel>
    <Ellipse Name="Light" Fill="Red" Height="200" Width="200"
             Margin="50" />
    <Button Width="150" Content="Change Lights"
            HorizontalAlignment="Center" Click="Button_Click" />
</StackPanel>
```

```csharp
private void Button_Click(object sender, RoutedEventArgs e)
{
    var current = Light.Fill as SolidColorBrush;

    if (current.Color == Colors.Red)
    {
        Light.Fill = new SolidColorBrush(Colors.Green);
    }
    else
    {
        Light.Fill = new SolidColorBrush(Colors.Red);
    }
}
```

---

![bg](https://www.reactiongifs.com/r/tumblr_mazjy9e9Dg1rvbizho2_250.gif)

---

# MVVM

---

# The Model-View-ViewModel (MVVM) Pattern

Separation of logic and presentation

Having event handlers in the code-behind is bad for testing, since you cannot mock away the view

Changing the design of the view often also requires changes in the code, since every element has its different event handlers

The logic is tightly bound to the view. It's not possible to reuse the logic in an other view

---

<!-- _class: default -->

# MVVM

<br>

![height:300px](https://learn.microsoft.com/en-us/dotnet/architecture/maui/media/mvvm-pattern.png)

---

# MVVM

## Demo

---

# MVVM Concepts

There is conceptually only ever one MODEL

Code in code-behind should be ABSOLUTELY MINIMAL

A ViewModel should ALWAYS implement INotifyPropertyChanged

A ViewModel may be used for more than one view

---

# MVVM Design Patterns

Observer Pattern:

- `INotifyPropertyChanged`
- `ObservableCollection<T>`

Command Pattern:

- `ICommand`

---

# Thank You
