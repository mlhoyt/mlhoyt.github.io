---
title: HOW-TO extensible software design
date: 2020-05-22T11:21:48-07:00
tags: extensible
---

## [Wikipedia: extensibility](https://en.wikipedia.org/wiki/Extensibility)

- Uses a light framework that allows for changes.
  - As opposed to a big up-front design which is supposed to have accounted
    for all potential concerns / functionality.
- Small commands are made to prevent from losing the element of extensibility
- Following the principle of separating work elements into comprehensible units
- Avoid ... low cohesion and high coupling 
  - **Cohesion**: the degree to which the elements inside a module belong together
    In one sense, it is a measure of the strength of relationship between the
    class's methods and data themselves.
  - **Coupling**: is the degree of interdependence between software modules; a
    measure of how closely connected two routines or modules are.

## ACCU: The Philosophy of Extensible Software

[The Philosophy of Extensible Software](https://accu.org/index.php/journals/391)

To some degree this articule seems to blur the physical challenges (why making
software changes can be difficult) and the mental challenges (why we need to be
able to make changes and why we should want to make changes).

The physical challenges are basically just the ripple effect which is minimized
with good abstractions and interfaces but this dovetails into the fact that:

> Development leads to learning more about the *problem domain* and *our solution
> domain* which leads to new insights into both.

### Notes:

- How does software resist change?
  - Ripple effect
    - _Changing the interface of a module means we must make corresponding changes to the use of the interface elsewhere._
  - Friction of change
    - Includes ripple effect
    - Issues arising from multiple simultaneous changes
    - Issues arising from dependency changes
    - _Sound waves cannot travel in a vacuum, likewise software ripples cannot pass from one module to another if they are well spaced._
  - Process roadblocks
    - _Conway's law: Software takes on many of the attributes of the organisation and process that creates it._
      - Organizations can create processes that require work documentation, prioritization, authorization, scheduling, changing, testing, validation, and releasing.
    - _... price of everything, value of nothing._
      - [An explanation of "price of everything, value of nothing"](https://kexino.com/marketing/the-price-of-everything-the-value-of-nothing/#:~:text=Oscar%20Wilde%20is%20credited%20with,Especially%20in%20business.)
      - Summary: **price** is what you pay, **value** is what you get
    - _Like [many things], over time software changes and without active attempts to [maintain and] improve [...] it, it invariably deteriorates._
- Development is a learning process
  - Development leads to learning more about the *problem domain* and *our solution domain* which leads to new insights into both.
- How does extensibility work?
  - **TODO**
- How does extensibility help?
  - **TODO**
- Extensibility is not resuse.
  - **TODO**
- Other
  - _... view software as models [...] or abstractions [...]._
  - _... view of your software as a living, growing, entity._
  - _Extensibility is a technique for reasoning about our software [... which] is often an attribute of other techniques [... e.g.] design patterns._

## StackExchange: Software Engineering: How Can I Improve The Ease Of Which I Can Extend My Software

[How Can I Improve The Ease Of Which I Can Extend My Software](https://softwareengineering.stackexchange.com/questions/62168/how-can-i-improve-the-ease-of-which-i-can-extend-my-software)

## Code Mag: Design for Extensibility

[Design for Extensibility](https://www.codemag.com/article/0801041/Design-for-Extensibility)
