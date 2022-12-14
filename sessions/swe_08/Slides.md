class: center, middle

<img src="https://www.saa-authors.eu/picture/739/ftw_768/saa-mtcwmza4nzq5mq.jpg" width="50%">

# Analysis, Design and Software Architecture

## Software Engineering Session 8

Helge Pfeiffer, Assistant Professor,<br>
[Research Center for Government IT](https://www.itu.dk/forskning/institutter/institut-for-datalogi/forskningscenter-for-offentlig-it),<br> 
[IT University of Copenhagen, Denmark](https://www.itu.dk)<br>
`me@itu.dk`

---

class: center, middle

# Info and Feedback

---

### Important: Late submissions of assignments

In case you did not yet submit your solution to assignment 05 due to some unfortunate reason, you can do so before this Friday (Oct. 28th) 23:59.

You hand-in your late assignment via LearnIT (<https://learnit.itu.dk/mod/assign/view.php?id=166606>) by providing a link to your repository and by uploading a PDF file as specified in the original task description (<https://github.com/itu-bdsa/assignment-05/blob/main/README.md>)

Important, this is the final deadline for late hand-ins of assignment 05. If you do not submit your solution before the deadline, it cannot be approved. Remember approval of mandatory assignments is necessary to be eligible for the exam.

---

### Project

Assignments are over, the project starts now.

<https://github.com/itu-bdsa/project-description>

---

class: center, middle

# Building Software

---

### A super simple program

The program (binary) `hello` is built from the following [source code](./bdsa_greeter/csharp/Program.cs):

```
string msg = @"
 █     █░▓█████  ██▓     ▄████▄   ▒█████   ███▄ ▄███▓▓█████ 
▓█░ █ ░█░▓█   ▀ ▓██▒    ▒██▀ ▀█  ▒██▒  ██▒▓██▒▀█▀ ██▒▓█   ▀ 
▒█░ █ ░█ ▒███   ▒██░    ▒▓█    ▄ ▒██░  ██▒▓██    ▓██░▒███   
░█░ █ ░█ ▒▓█  ▄ ▒██░    ▒▓▓▄ ▄██▒▒██   ██░▒██    ▒██ ▒▓█  ▄ 
░░██▒██▓ ░▒████▒░██████▒▒ ▓███▀ ░░ ████▓▒░▒██▒   ░██▒░▒████▒
░ ▓░▒ ▒  ░░ ▒░ ░░ ▒░▓  ░░ ░▒ ▒  ░░ ▒░▒░▒░ ░ ▒░   ░  ░░░ ▒░ ░
  ▒ ░ ░   ░ ░  ░░ ░ ▒  ░  ░  ▒     ░ ▒ ▒░ ░  ░      ░ ░ ░  ░
  ░   ░     ░     ░ ░   ░        ░ ░ ░ ▒  ░      ░      ░   
    ░       ░  ░    ░  ░░ ░          ░ ░         ░      ░  ░
                        ░                                   
                      ▄▄▄     ▄▄▄█████▓                     
                     ▒████▄   ▓  ██▒ ▓▒                     
                     ▒██  ▀█▄ ▒ ▓██░ ▒░                     
                     ░██▄▄▄▄██░ ▓██▓ ░                      
                      ▓█   ▓██▒ ▒██▒ ░                      
                      ▒▒   ▓▒█░ ▒ ░░                        
                       ▒   ▒▒ ░   ░                         
                       ░   ▒    ░                           
                           ░  ░                             
                                                            
                   ██▓▄▄▄█████▓ █    ██                     
                  ▓██▒▓  ██▒ ▓▒ ██  ▓██▒                    
                  ▒██▒▒ ▓██░ ▒░▓██  ▒██░                    
                  ░██░░ ▓██▓ ░ ▓▓█  ░██░                    
                  ░██░  ▒██▒ ░ ▒▒█████▓                     
                  ░▓    ▒ ░░   ░▒▓▒ ▒ ▒                     
                   ▒ ░    ░    ░░▒░ ░ ░                     
                   ▒ ░  ░       ░░░ ░ ░                     
                   ░              ░                         
                                                          
";
Console.WriteLine(msg);
```

---

### Building Software

So far, you likely invoked your programs via `dotnet run` in your project folder.
That command calls internally `dotnet build`, the actual build step.

---

### Exercise: Download and run a program

I have built the program for you. Download it from here: [`hello.zip`](./hello.zip), can you run it?

  - If yes, what do you see?
  - If not, what do you see?

In a terminal, run the program via

```bash
cd ~/Downloads
unzip hello.zip
./hello
```

---

### Exercise: Download and run a program

I have built the program for you. Download it from here: [`hello.zip`](./hello.zip), can you run it?

  - If yes, what do you see?
  - If not, what do you see?

In a terminal, run the program via

```bash
cd ~/Downloads
unzip hello.zip
dotnet hello.dll
```

---

### Recap: What was software again?

  > "Software is the collection of **all artifacts**, which allow (a) **suitably educated person(s)** with access to specified and suitable **hardware** to **instantiate** a running system.
  > 
  > Additionally, the collection of such artifacts allow such suitably educated person(s) to **understand** and **reason** about a systems' working and properties and let them **understand** why the system is as it is and why it behaves the way it does."
  > 
  > Helge

--

That is, software is first software once it can be run (instantiated).
To be able to run software you have to first build (and distribute) it.

---

### Detour: Disabling Telemetry when building with .Net

Microsoft (and other big organizations) wants to know what you are doing, see https://learn.microsoft.com/en-us/dotnet/core/tools/telemetry.


<iframe src="https://dotnet.microsoft.com/en-us/platform/telemetry" width="100%" height="500px"></iframe>

---

### Detour: Disabling Telemetry when building with .Net

Create an environment variable (in your `.bashrc` or whichever shell you are using) before invoking any `dotnet` commands:

  * `DOTNET_CLI_TELEMETRY_OPTOUT=1` or 
  * `DOTNET_CLI_TELEMETRY_OPTOUT=true`

---


### How to build .Net software?


  > `dotnet publish` compiles the application, reads through its dependencies specified in the project file, and publishes the resulting set of files to a directory. The output includes the following assets:
  > 
  >  * Intermediate Language (IL) code in an assembly with a `dll` extension.
  >  * A .deps.json file that includes all of the dependencies of the project.
  >  * A .runtimeconfig.json file that specifies the shared runtime that the application expects, as well as other configuration options for the runtime (for example, garbage collection type).
  >  * The application's dependencies, which are copied from the NuGet cache into the output folder.
  >
  > Publishing an app as framework-dependent produces a cross-platform binary as a `dll` file, and a platform-specific executable that targets your current platform. The `dll` is cross-platform while the executable isn't. 
  >
  > https://learn.microsoft.com/en-us/dotnet/core/deploying/#publish-framework-dependent

---

### How to build .Net software?

  > Applications you create with .NET can be published in two different modes, and the mode affects how a user runs your app.
  > 
  > Publishing your app as **self-contained** produces an **application that includes the .NET runtime and libraries**, and your application and its dependencies. Users of the application can run it on a machine that doesn't have the .NET runtime installed.
  > 
  > Publishing your app as **framework-dependent** produces an **application that includes only your application itself and its dependencies**. Users of the application have to separately install the .NET runtime.
  >
  > <https://learn.microsoft.com/en-us/dotnet/core/deploying/>

---

### How to build .Net software?

```bash
dotnet publish --help
```

---

### Building framework-dependent non-self-contained software


```bash
dotnet publish -p:AssemblyName=hello -o dist/linux/nosc --no-self-contained
```

```bash
 ls -lah dist/linux/
total 120K
drwxrwxr-x 2 me me  12K Oct 24 15:20 .
drwxrwxr-x 5 me me 4.0K Oct 24 13:26 ..
-rwxr-xr-x 1 me me  76K Oct 24 14:49 hello
-rw-rw-r-- 1 me me  385 Oct 24 15:10 hello.deps.json
-rw-rw-r-- 1 me me 8.0K Oct 24 14:49 hello.dll
-rw-rw-r-- 1 me me  11K Oct 24 14:49 hello.pdb
-rw-rw-r-- 1 me me  139 Oct 24 15:10 hello.runtimeconfig.json
```

---

### Building framework-dependent self-contained software

```bash
dotnet publish -p:AssemblyName=hello -o dist/linux/sc --self-contained
```

```bash
ls dist/linux/nosc
createdump
hello
hello.deps.json
hello.dll
hello.pdb
hello.runtimeconfig.json
libclrjit.so
libcoreclr.so
libcoreclrtraceptprovider.so
libdbgshim.so
libhostfxr.so
libhostpolicy.so
libmscordaccore.so
libmscordbi.so
libSystem.Globalization.Native.so
libSystem.IO.Compression.Native.so
libSystem.Native.so
libSystem.Net.Security.Native.so
libSystem.Security.Cryptography.Native.OpenSsl.so
Microsoft.CSharp.dll
Microsoft.VisualBasic.Core.dll
Microsoft.VisualBasic.dll
```

---

### Building single-file (not framework-dependent) self-contained software

```bash
dotnet publish -p:AssemblyName=hello -p:PublishSingleFile=true \
    -o dist/linux/sf-sc -r linux-x64 --self-contained
```

The result is the executable [`hello`](./hello)

--

What is the issue with it?

  * It is big
  * It contains a specific version of .Net runtime and all dependencies. That might pose a security threat when deployed.

---

### What is this `hello` program actually?

--

```bash
file hello
hello: ELF 64-bit LSB pie executable, x86-64, version 1 (GNU/Linux), dynamically linked, interpreter /lib64/ld-linux-x86-64.so.2, BuildID[sha1]=d1cc43e1dffa1afdaebd5a0f79af0201bb63a2f5, for GNU/Linux 2.6.32, stripped, too many notes (256)
```

---

### Does it contain everything it needs to run?

--

```bash
ldd hello
    linux-vdso.so.1 (0x00007ffcce0c8000)
    libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007f06df6e0000)
    libdl.so.2 => /lib/x86_64-linux-gnu/libdl.so.2 (0x00007f06df6db000)
    libz.so.1 => /lib/x86_64-linux-gnu/libz.so.1 (0x00007f06df6bf000)
    librt.so.1 => /lib/x86_64-linux-gnu/librt.so.1 (0x00007f06df6ba000)
    libgcc_s.so.1 => /lib/x86_64-linux-gnu/libgcc_s.so.1 (0x00007f06df69a000)
    libstdc++.so.6 => /lib/x86_64-linux-gnu/libstdc++.so.6 (0x00007f06de800000)
    libm.so.6 => /lib/x86_64-linux-gnu/libm.so.6 (0x00007f06deb19000)
    libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f06de400000)
    /lib64/ld-linux-x86-64.so.2 (0x00007f06df704000)
```

---

### What are these `.so` files?

--

Unfortunately, such shared libraries are also called DLLs on Windows...

--

Figure out which Debian package contains that file
```bash
dpkg -S /lib64/ld-linux-x86-64.so.2
libc6:amd64: /lib64/ld-linux-x86-64.so.2
```

--

What is `/lib64/ld-linux-x86-64.so.2`?

```bash
$ apt info libc6
Package: libc6
Version: 2.35-0ubuntu3.1
Priority: required
Section: libs
Source: glibc
Origin: Ubuntu
Maintainer: Ubuntu Developers <ubuntu-devel-discuss@lists.ubuntu.com>
Original-Maintainer: GNU Libc Maintainers <debian-glibc@lists.debian.org>
Bugs: https://bugs.launchpad.net/ubuntu/+filebug
Installed-Size: 13.9 MB
Depends: libgcc-s1, libcrypt1 (>= 1:4.4.10-10ubuntu4)
Recommends: libidn2-0 (>= 2.0.5~), libnss-nis, libnss-nisplus
Suggests: glibc-doc, debconf | debconf-2.0, locales
Breaks: busybox (<< 1.30.1-6), fakeroot (<< 1.25.3-1.1ubuntu2~), hurd (<< 1:0.9.git20170910-1), ioquake3 (<< 1.36+u20200211.f2c61c1~dfsg-2~), iraf-fitsutil (<< 2018.07.06-4), libgegl-0.4-0 (<< 0.4.18), libtirpc1 (<< 0.2.3), locales (<< 2.35), locales-all (<< 2.35), macs (<< 2.2.7.1-3~), nocache (<< 1.1-1~), nscd (<< 2.35), openarena (<< 0.8.8+dfsg-4~), openssh-server (<< 1:8.2p1-4), r-cran-later (<< 0.7.5+dfsg-2), wcc (<< 0.0.2+dfsg-3)
Replaces: libc6-amd64
Homepage: https://www.gnu.org/software/libc/libc.html
Task: minimal, server-minimal
Original-Vcs-Browser: https://salsa.debian.org/glibc-team/glibc
Original-Vcs-Git: https://salsa.debian.org/glibc-team/glibc.git
Download-Size: 3,235 kB
APT-Manual-Installed: yes
APT-Sources: http://apt.pop-os.org/ubuntu jammy-updates/main amd64 Packages
Description: GNU C Library: Shared libraries
 Contains the standard libraries that are used by nearly all programs on
 the system. This package includes shared versions of the standard C library
 and the standard math library, as well as many others.
```

---

### Dynamically linked binaries

The .Net compiler creates dynamically linked binaries as the previous example illustrates.

Other languages that you know of, e.g., Go, create statically linked binaries.

---

### Static vs. dynamic linking

<img src="https://www.baeldung.com/wp-content/uploads/sites/4/2021/01/StaticVsDynamicLinking-1536x1223-1.png" width="80%">

Source: https://www.baeldung.com/cs/dynamic-linking-vs-dynamic-loading

---

### Why does it matter?

Design your software to be **executed**.

Code that cannot be build and executed is no software.
Design consideration how to build easily and continuously?
Choose languages and tools to that allow you to iterate swiftly.
Design to build components:

#### Build components

Remember that components can be deployed separately.
That might be your goal, build components.

---

### "Fat" statically linked binaries

Does a statically linked binary actually contain everything we need to run a program?

--

No, not necessarily. 
But languages like Go support you in building "Fat" statically linked binaries.
These are binaries that also include all static artifacts that a program may need during execution.

```go
import _ "embed"

//go:embed hello.txt
var s string
print(s)
```

Source: https://pkg.go.dev/embed
<!-- http://www.inanzzz.com/index.php/post/1rwm/including-and-reading-static-files-with-embed-directive-at-compile-time-in-golang
 -->

I believe .Net does not support that.

---

### `make` and `Makefile`s

These were a lot of commands to build software.
I cannot remember them!
Since building software is actually quite difficult, build systems were invented, which in turn are [usually also difficult](https://www.microsoft.com/en-us/research/uploads/prod/2018/03/build-systems.pdf).

One of the oldest and most widely used is the `make` tool.

--

  > Make originated with a visit from Steve Johnson (author of yacc, etc.), storming into my office, cursing the Fates that had caused him to waste a morning debugging a correct program (bug had been fixed, file hadn't been compiled, cc *.o was therefore unaffected). As I had spent a part of the previous evening coping with the same disaster on a project I was working on, the idea of a tool to solve it came up. It began with an elaborate idea of a dependency analyzer, boiled down to something much simpler, and turned into Make that weekend. Use of tools that were still wet was part of the culture. Makefiles were text files, not magically encoded binaries, because that was the Unix ethos: printable, debuggable, understandable stuff.
  >
  > — Stuart Feldman, The Art of Unix Programming, Eric S. Raymond 2003

Source: https://en.wikipedia.org/wiki/Make_(software)

---

### The `make` Tool

```
MAKE(1)                          User Commands                         MAKE(1)

NAME
       make - GNU make utility to maintain groups of programs

SYNOPSIS
       make [OPTION]... [TARGET]...

DESCRIPTION
       The  make  utility will determine automatically which pieces of a large
       program need to be recompiled, and  issue  the  commands  to  recompile
       them.   The  manual describes the GNU implementation of make, which was
       written by Richard Stallman and Roland McGrath, and is currently  main‐
       tained  by  Paul  Smith.   Our examples show C programs, since they are
       very common, but you can use make with any programming  language  whose
       compiler can be run with a shell command.  In fact, make is not limited
       to programs.  You can use it to describe any task where some files must
       be updated automatically from others whenever the others change.

       To  prepare to use make, you must write a file called the makefile that
       describes the relationships among files in your program, and the states
       the  commands for updating each file.
```

---

### The `make` Tool

<img src="https://makefiletutorial.com/assets/dependency_graph.png" width="100%">

Source: https://makefiletutorial.com/

---

### Investigating some Makefiles 

  * [.Net project `Makefile`](bdsa_greeter/csharp/Makefile)
  * [Go project `Makefile`](bdsa_greeter/go/Makefile)
  * [Python project `Makefile`](bdsa_greeter/python/Makefile)

A more thorough introduction of `Makefile`s and their syntax can be found here: <https://makefiletutorial.com/>

--

What is the problem with these `Makefile`s?

--

It builds for me...
Do you have the right environment setup?

  * Do you have the `dotnet format` tool installed?
  * Do you have the right directory structure?
  * Do you have the right version of .Net Core, Python, Go, etc. installed?
  * Can you build the right version of your program for the target environment?

---

### Cross-compilation

Languages like C# and Go can be compiled to other target architectures and operating systems than the one on which the software is built.

```bash
# MacOS
dotnet publish -p:AssemblyName=hello -p:PublishSingleFile=true \
    -o dist/macos/ -r osx-x64 --self-contained

# Windows
dotnet publish -p:AssemblyName=hello -p:PublishSingleFile=true \
    -o dist/windows/ -r win-x64 --self-contained

# Linux
dotnet publish -p:AssemblyName=hello -p:PublishSingleFile=true \
    -o dist/linux/sf-sc -r linux-x64 --self-contained
```

Not all programming languages support this.

---

### GitHub Actions

More elaborate build systems, such as, GitHub Actions, attempt to "abstract away" the environment by establishing one build environment for the entire development team.

--

<img src="https://learn.microsoft.com/en-us/training/wwl-azure/introduction-to-github-actions/media/actions-structure-1b8448db.png" width="50%">

Source: https://learn.microsoft.com/en-us/training/modules/introduction-to-github-actions/3-explore-actions-flow

---

### Alternatives to GitHub Actions

Self-hosted most often in bigger organizations and companies:

  * [Jenkins](https://jenkins.io/index.html)
  * [Bamboo](https://www.atlassian.com/software/bamboo)
  * [TeamCity](https://www.jetbrains.com/teamcity/)
  * [Concourse](https://concourse.ci)
  * [Azure DevOps Server](https://azure.microsoft.com/en-us/services/devops/server/)
  
Build systems as a service:

  * [Travis CI](https://travis-ci.org/)
  * [CircleCI](https://circleci.com)
  * [Github Actions](https://github.com/features/actions)
  * [Wercker](http://www.wercker.com)
  * [Azure Pipelines](https://azure.microsoft.com/en-in/products/devops/pipelines/)

---

### GitHub Actions

  > GitHub tracks events that occur. Events can trigger the start of workflows.
  > 
  > Workflows can also start on cron-based schedules and can be triggered by events outside of GitHub.
  > 
  > They can be manually triggered.
  > 
  > Workflows are the unit of automation. They contain Jobs.
  > 
  > Jobs use Actions to get work done.
  > 
  > <https://learn.microsoft.com/en-us/training/modules/introduction-to-github-actions/>

Source: https://learn.microsoft.com/en-us/training/modules/introduction-to-github-actions/

---

### GitHub Actions Starter Workflows

```yaml
name: .NET

on:
  push:
    branches: [ $default-branch ]
  pull_request:
    branches: [ $default-branch ]

jobs:
  build:

    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v3
    - name: Setup .NET
      uses: actions/setup-dotnet@v3
      with:
        dotnet-version: 6.0.x
    - name: Restore dependencies
      run: dotnet restore
    - name: Build
      run: dotnet build --no-restore
    - name: Test
      run: dotnet test --no-build --verbosity normal
```

Source: https://github.com/actions/starter-workflows/blob/main/ci/dotnet.yml

---

### GitHub Actions Workflows


  > Workflows include several standard syntax elements.
  > 
  >  **Name**: is the name of the workflow. It's optional but is highly recommended. It appears in several places within the GitHub UI.
  > 
  >  **On**: is the event or list of events that will trigger the workflow.
  >
  >  **Jobs**: is the list of jobs to be executed. Workflows can contain one or more jobs.
  >
  >  **Runs-on**: tells Actions which runner to use.
  >
  >  **Steps**: It's the list of steps for the job. Steps within a job execute on the same runner.
  >
  >  **Uses**: tells Actions, which predefined action needs to be retrieved. For example, you might have an action that installs node.js.
  >
  >  **Run**: tells the job to execute a command on the runner. For example, you might execute an NPM command.
  >
  > <https://learn.microsoft.com/en-us/training/modules/introduction-to-github-actions/5-describe-standard-workflow-syntax-elements>

--

### GitHub Actions Workflow Syntax

A more thorough description of GitHub Actions workflow syntax: <https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions>

<!-- https://www.jenkins.io/doc/book/pipeline/jenkinsfile/#using-a-jenkinsfile -->

---

### GitHub Actions Events

  * Scheduled events
  * Code events
  * Manual events
  * Webhook events
  * External events

Source: https://docs.github.com/actions/learn-github-actions/events-that-trigger-workflows

---

### GitHub Actions Jobs

  > By default, if a workflow contains multiple jobs, they run in parallel.

--

`needs` allows to define an order between jobs:

```yaml
jobs:
  startup:
    runs-on: ubuntu-latest
    steps:

      - run: ./setup_build_env.sh
  build:
    needs: startup
    steps:

      - run: ./build.sh
```

Source: https://learn.microsoft.com/en-us/training/modules/introduction-to-github-actions/7-explore-jobs

---

### GitHub Actions Runners

  > When you execute jobs, the steps execute on a Runner.
  > 
  > The steps can be the execution of a shell script or the execution of a predefined Action.
  > 
  > GitHub provides several hosted runners to avoid you needing to spin up your infrastructure to run actions.
  > 
  > Now, the maximum duration of a job is 6 hours, and for a workflow is 72 hours.
  >
  > <https://learn.microsoft.com/en-us/training/modules/introduction-to-github-actions/8-explore-runners>

---

### GitHub Actions in the Wild: `libgit2sharp`

```yaml
name: CI
on:
  push:
    branches: [master, release-*]
    tags:
      - '[0-9]+.[0-9]+.[0-9]+'
      - '[0-9]+.[0-9]+.[0-9]+-*'
  pull_request:
  workflow_dispatch:
env:
  DOTNET_NOLOGO: true
jobs:
  build:
    name: Build
    runs-on: ubuntu-20.04
    steps:
      - name: Checkout
        uses: actions/checkout@v2.3.4
        with:
          fetch-depth: 0
...          
```

Source: 
<https://github.com/libgit2/libgit2sharp/blob/master/.github/workflows/ci.yml>

<!-- https://github.com/libgit2/libgit2sharp/actions/runs/2083148186/jobs/3024758433
 -->

---

### Build Matrix???

Run multiple jobs in parallel to build, test, etc. software on different versions of different operating systems, frameworks etc.

```yaml
jobs:
  example_matrix:
    strategy:
      matrix:
        os: [ubuntu-22.04, ubuntu-20.04]
        version: [10, 12, 14]
    runs-on: ${{ matrix.os }}
    steps:
      - uses: actions/setup-node@v3
        with:
          node-version: ${{ matrix.version }}
```

---

### Azure Pipeline in the Wild: Coronapas App

Azure Pipelines and GitHub Actions are similar but different: <https://learn.microsoft.com/en-us/dotnet/architecture/devops-for-aspnet-developers/actions-vs-pipelines
>

```yaml
stages:

## ===== CI START=====
  - stage: 'CI_build_and_unit_tests'
    jobs:
    - job: 'build_and_test_debug'
      displayName: Build solution and run tests (Unit and UI)
      pool: 
        name: $(macAgentsPool)

      variables:
        buildPlatform: 'Any CPU'
        buildConfiguration: 'Debug'

      steps:
        - task: NuGetToolInstaller@1
          displayName: 'Install NuGet'

        - task: NuGetCommand@2
          displayName: 'Restore NuGet packages'
          inputs:
            command: 'restore'
            restoreSolution: '$(solution)'

        - task: MSBuild@1
          displayName: 'Build solution'
          inputs:
            solution: '$(solution)'
            platform: '$(buildPlatform)'
            configuration: '$(buildConfiguration)'
            msbuildArguments: '/maxcpucount /nodeReuse:False /p:PackageLocation="$(build.artifactStagingDirectory)"'
            clean: true
            restoreNugetPackages: true
            createLogFile: true
```

Source: https://github.com/Sundhedsdatastyrelsen/Coronapas.Mobile/blob/main/azure-pipelines.yml




---

### GitHub Actions in the Wild: Altinn

You can do more than building software with GitHub Actions:

```
name: Auto Assign to Project

on:
  issues:
    types:
      - opened

jobs:
  add-to-project:
    name: Add issue to Team Platform project
    runs-on: ubuntu-latest
    steps:
      - uses: actions/add-to-project@main
        with:
          project-url: https://github.com/orgs/Altinn/projects/20
          github-token: ${{ secrets.ASSIGN_PROJECT_TOKEN }}
```

Source: https://github.com/Altinn/altinn-platform/blob/main/.github/workflows/assign-issues-to-projects.yml

---

### More GitHub Actions ...

  * Marketplace of Actions: <https://github.com/marketplace?type=actions>
  * Getting started with GitHub Actions: https://github.com/skills/hello-github-actions

---

### GitHub Actions, what is the problem?

--

Cloud-computing, the mainframes of today...

<img src="https://upload.wikimedia.org/wikipedia/commons/6/61/IBM_System_Z9_%28type_2094_inside%29.jpg" width="30%">

---

### Run GitHub Actions Locally

<img src="https://github.com/nektos/act/wiki/quickstart/act-quickstart-2.gif" width="90%">

Source: https://github.com/nektos/act

--

Alternatively, you might want to combine `Makefile`s with other build systems.

---

### Your turn!

![](https://i.gifer.com/6E2.gif)

  * Go to the exercise session.
  * Register your project groups **best today**, see <https://github.com/itu-bdsa/project-description#week-zero-week-42>
  * Work on the project, see <https://github.com/itu-bdsa/project-description>
