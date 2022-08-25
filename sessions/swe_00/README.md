The following describes what we would like you to do before the first session of the course _Analysis, Design and Software Architecture_ on Tuesday, Aug. 30th (10:00).

We would like that you prepare the following four things:

  * [A) Reading](#A-reading)
  * [B) Register accounts](#B-register-accounts)
  * [C) Setup workstation](#C-setup-workstation)
  * [D) Get to know Docker](#D-get-to-know-docker)

In case you have problems with installation of any of the tools, do not despair.
You will have to ask the TAs for help during the first exercise session on Tuesday, Aug. 30th (12:00 - 14:00).

--------------------------------------------------------------------------------

# A) Reading

Read chapter 15 from *Objects First with Java: A Practical Introduction Using BlueJ* (Sixth Edition). It is the book that you used in the first semester in *"Grundlæggende Programmering"*.

--------------------------------------------------------------------------------

# B) Register accounts

## Create an account on [GitHub.com](https://github.com)

If you do not have an account on [GitHub.com](https://github.com) yet, sign-up via https://github.com/signup

Note, in this course we are using GitHub.com **not*- https://github.itu.dk/.

## Discord

In case you are not registered for Discord, navigate to https://discord.com/ and create a user there.

## Overleaf

You might want to author LaTeX documents with the browser-based tool [Overleaf](https://www.overleaf.com/login)

To use it choose:

1. `Login with ORCID`
1. `Access through your institution`
1. Enter `IT University of Copenhagen`
1. Login with your credentials

## *Optional*: Activate Azure credits

Navigate to <https://my.visualstudio.com/> and log in with your `@itu.dk` credentials.

--------------------------------------------------------------------------------

# C) Setup workstation

## Install a package manager

- Windows
  - Install [winget](https://docs.microsoft.com/en-us/windows/package-manager/winget/)
- MacOS
  - Install [homebrew](https://brew.sh/) by running the following in your terminal:

    ```bash
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    ```

- Linux
  - Depending on your distribution, you have a package manager like `apt` (Debian-based distributions) or `rpm` (RedHat-based distributions), etc. already installed.

## Install software (latest versions)

1. Windows only:
    1. Windows Terminal: <https://docs.microsoft.com/en-us/windows/terminal/>
    1. Custom prompt: <https://docs.microsoft.com/en-us/windows/terminal/tutorials/custom-prompt-setup>
    1. Windows Subsystem for Linux: <https://docs.microsoft.com/en-us/windows/wsl/> with Ubuntu
    1. Ubuntu
1. Visual Studio Code: <https://code.visualstudio.com/> with extensions:
    - C#
    - .NET Core Test Explorer
    - markdownlint
    - PowerShell
    - Remote Development
    - REST Client
    - Visual Studio IntelliCode
    - vscode-icons
    - Docker
1. .NET <https://dotnet.microsoft.com/> (for Windows also in WSL)
1. PowerShell <https://docs.microsoft.com/en-us/powershell/scripting/overview> (for Windows also in WSL)
1. Docker: <https://www.docker.com/products/docker-desktop/>
1. LaTeX: <https://www.latex-project.org/get/>
1. Git: <https://git-scm.com/downloads>
1. *Optional*: Windows and Mac only:
    - Visual Studio Community: <https://visualstudio.microsoft.com/downloads/>
1. *Optional*: JetBrains Rider: <https://www.jetbrains.com/rider/>
    - Note that you can get a free student license for all JetBrais IDEs that you use during your education, see https://www.jetbrains.com/community/education/#students

## Choose and install a diagramming tool

Identify a tool that you can use in the course to draw diagrams in general and UML diagrams in particular.
Make sure that the tool can export the diagrams to image files that can include them in a LaTeX document.

In case you do not have a favorite diagramming tool yet, you might consider the following:

- Diagram markup languages:

  - [MermaidJS](https://mermaid-js.github.io/)
  - [PlantUML](https://plantuml.com/)
  - [GraphViz](https://graphviz.org/)

- WYSIWYG diagram editors:

  - [Umbrello](https://umbrello.kde.org/)
  - [Dia](http://dia-installer.de/index.html.en)
  - [ArgoUML](http://argouml.tigris.org/)
  - [StarUML](http://staruml.io/)
  - [Papyrus](https://www.eclipse.org/papyrus/)
  - [OmniGraffle](https://www.omnigroup.com/omnigraffle) (commercial product)

--------------------------------------------------------------------------------

# D) Get to know Docker

1. Follow the tutorial presented on Docker Desktop: <https://docs.docker.com/desktop/get-started/>
1. Follow this guide: [Quickstart: Run SQL Server container images with Docker](https://docs.microsoft.com/en-us/sql/linux/quickstart-install-connect-docker).

By now you should be able to run SQL Server in a container on your local machine, connect to it and run T-SQL queries against it.
