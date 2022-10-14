---
marp: true
author: Rasmus Lystrøm
theme: default
class: invert
---

# The Gilded Rose (Recap)

![bg right](images/gilded-rose.jpg)

Rasmus Lystrøm
Associate Professor
ITU

---

# Code Metrics (Visual Studio Enterprise)

**Maintainabilityability Index**: Between 0 and 100. Higher is better. Aim for higher than 20

**Cyclomaticmatic Complexity**: Lower is better. Split methods with complexity > 10

**Lines of Source Code** / **Lines of Executable Code**

**Inheritance of Inheritance**: Between 1 and infinity. Lower is better, but sometimes inheritance is good

**Class Coupling**: Lower is better. Aim for max 9

Source: [Code metrics values](https://learn.microsoft.com/en-us/visualstudio/code-quality/code-metrics-values).

---

# Code Metrics: Original

| |Maintainability|Cyclomatic|Inheritance|Coupling|LoC|Exe.Code|
|:----|:---:|:---:|:---:|:---:|:---:|:---:|
|Program|50|22|1|4|131|34|
|Item|100|6|1|0|8|0|
|Total|75|28|1|4|144|34|

---

# Approach

Understand the task at hand – inspect the code
Write tests to ensure the program works to specification
Refactor, refactor, refactor
Extract methods
Implement *Conjured*
Refactor, refactor, refactor

1. Introduce polymorphism?
1. Immutable, functional, pattern matching?

---

# The Gilded Rose in C♯

## Demo

---

# Code Metrics: Refactored (Make the Change Easy)

| |Maintainability|Cyclomatic|Inheritance|Coupling|LoC|Exe.Code|
|:----|:---:|:---:|:---:|:---:|:---:|:---:|
|Program|58|3|1|5|52|9|
|Item|100|6|1|0|8|0|
|GildedRose|72|16|1|3|74|22|
|Total|76|25|1|6|142|31|

---

# Code Metrics: Refactored (Polymorphed)

| |Maintainability|Cyclomatic|Inheritance|Coupling|LoC|Exe.Code|
|:----|:---:|:---:|:---:|:---:|:---:|:---:|
|Program|58|3|1|8|34|9|
|Item|91|8|1|1|21|5|
|&nbsp; AgedBrie|71|2|2|2|15|5|
|&nbsp; Backst.P.|63|4|2|2|23|9|
|&nbsp; Sulfuras|100|1|2|1|6|0|
|&nbsp; Conjured|71|2|2|2|15|5|
|Total|75|20|2|9|132|33|

---

# Code Metrics: Refactored (Immutable/Functional)

| |Maintainability|Cyclomatic|Inheritance|Coupling|LoC|Exe.Code|
|:----|:---:|:---:|:---:|:---:|:---:|:---:|
|Program|68|3|1|8|33|10|
|GildedRose|69|12|1|6|81|30|
|Item|100|1|1|1|1|0|
|ItemType|100|1|1|0|1|0|
|Total|84|17|1|10|125|40|

---

# The Gilded Rose in F♯

## Demo

---

# Summary

| |Maintainability|Cyclomatic|Inheritance|Coupling|LoC|Exe.Code|
|:----|:---:|:---:|:---:|:---:|:---:|:---:|
|Original|75|28|1|4|144|34|
|Refactored|76|25|1|6|142|31|
|Polymorp.|75|20|2|9|132|33|
|Functional|84|17|1|10|125|40|
|F♯| | | | |64| |

---

# Conclusion / Discussion
