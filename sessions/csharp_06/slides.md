---
marp: true
author: Rasmus Lystrøm
theme: default
class: invert
---

# C♯ 06: Asynchronous and Parallel Programming

![bg right](https://upload.wikimedia.org/wikipedia/commons/5/55/Overhead_Communication_Cables_-_Kolkata_2011-08-29_4821.JPG)

Rasmus Lystrøm
Associate Professor
ITU

---

# The Gilded Rose

Introduction

---

![bg](https://www.macmillandictionaryblog.com/wp-content/uploads/2018/01/dictionary-1024x575.jpg)

# Dictionary

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

# Async ≠ Parallel ≠ Threads

---

![bg](images/concurrency.jpg)

# Concurrency

---

# Concurrency I

> A property of systems in which several computations are executing simultaneously, and potentially interacting with each other. The computations may be executing on multiple cores in the same chip, preemptively time-shared threads on the same processor, or executed on physically separated processors.

---

# Concurrency II

> Multiple tasks which start, run, and complete in overlapping time periods, in no specific order.

---

![bg](https://live.staticflickr.com/7517/15550717054_a3a85a6573_z.jpg)

# Parallelism

---

# Parallelism

> When multiple tasks OR several parts of a unique task literally run at the same time, e.g., on a multi-core processor.

---

<!-- _class: default -->

![bg](images/thread.jpg)

# Multithreading

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

# Multithreading

> Software implementation which allows different threads to be executed concurrently.
> A multithreaded program appears to be doing several things at the same time even when it's running on a single-core machine.
> Compare to chatting with different people through various IM windows; although you're switching back and forth, the net result is that you're having multiple conversations at the same time.

---

<!-- _class: default -->

![bg](https://www.duperrin.com/english/wp-content/uploads/2020/06/desynchronized-organization.jpg)

# Asynchronous methods

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

# Asynchronous methods

> Asynchrony is used to present the impression of concurrent or parallel tasking.
> Normally used for a process that needs to do work away from the current application where we don't want to wait and block our application awaiting the response.

---

<!-- _class: default -->

![bg](images/nails.jpg)

# Threads

---

<!-- _class: default -->

![bg contain](images/stack-heap.png)

---

<!-- _class: default -->

![bg width:90%](images/main-worker-thread.png)

# Threads Example

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

![bg](images/race-condition.png "Ralf Schumacher, 3 Mar 2002, Reporter Images")

# Race Condition

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

# Race Condition

> Behavior of a program where the output is dependent on the sequence or timing of other uncontrollable events.
--> Bug, when events do not happen in the order the programmer intended.

---

<!-- _class: default -->

![bg width:250%](images/race-condition.png)

# Race Condition

## Demo

---

![bg](images/deadlock.jpg "Deadlock by David Mailtland")

# Deadlock

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

# Deadlock

> A situation in which two or more competing actions are each waiting for the other to finish, and thus neither ever does.

---

![bg right](https://i.stack.imgur.com/kKSUL.gif)

# Deadlock

## Demo

---

<!-- _class: default -->

![bg](https://upload.wikimedia.org/wikipedia/commons/2/21/RyersonUniversityLibrary.JPG)

# Task Parallel Library

## Demo

---

# TPL

Task.Run
Task.Factory…
Task.Delay
Parallel.For
Parallel.ForEach
Parallel.Invoke

Parallel Linq --> .AsParallel()

---

![bg](images/funko_pop_collection_by_needxxx_dazxmlc-fullview.jpg)

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

# Concurrent Collections

---

# `System.Collections.Concurrent`

`ConcurrentQueue<T>`
`ConcurrentStack<T>`
`BlockingCollection<T>`
`ConcurrentDictionary<TKey, TValue>`

---

<!-- _class: default -->

![bg](images/i-am-a-waiting.png)

# Asynchronous Programming

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

---

# Asynchronous Programming

> **Asynchronous programming** is a means of **parallel programming** in which a unit of work runs separately from the main application thread and notifies the calling thread of its completion, failure or progress.

---

# Asynchronous Programming

## Demo

---

# `async` / `await`

`async` -->

Method must return `void`, `Task`, `Task<T>`, `ValueTask`, `ValueTask<T>`, or a task-like type.

Specifically: a type, which satisfy the *async* pattern, meaning a `GetAwaiter` method must be accessible.

`await` --> Await task(s)...

Note: Asynchronous `Main` and *test* methods must return `Task`

---

# Async ≠ Parallel ≠ Threads

*Async*: Non-blocking UI, background tasks, asynchronous

*Parallel*: Speed, Multiprocessor, Parallel execution

*Thread*: Low-level building block. Do not use directly!

---

# Thank you
