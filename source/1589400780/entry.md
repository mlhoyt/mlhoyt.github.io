---
title: Parse, Don't Validate
date: 2020-05-13T13:13:00-07:00
tags: FP
---

[lexi-lambda Blog: Parse, Don't Validate](https://lexi-lambda.github.io/blog/2019/11/05/parse-don-t-validate/)

> "... strengthening the type of the argument [...] instead of weakening the type of its result"

In a way this allows the compiler to do more work because you cannot use the
function if you cannot get to the stronger type of the argument.

As opposed to the developer doing more work handling the potential weaker
states of the result.

> "... a refinement of the input type that preserves the knowledge gained in the type system.
> Both of these functions check the same thing, but [...] gives the caller access to the
> information it learned, while [...] just throws it away."

Parsing is a transformation and that transformation can add information gained
during the process yielding a stronger resulting type.

Validation yields a point-in-time exception (or not) after which the validated
type is no more meaningful than before -- and is just as capable of containing
an invalid value as a valid value.
