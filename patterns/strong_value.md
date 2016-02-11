{.epigraph}
> ... starting in The Mist, you are seeking ways to create Value,
> where value is defined as something of worth to some person or set
> of people whom we wish to serve.
>
> {.quote}
> [Value Stream - ScrumPLoP](https://sites.google.com/a/scrumplop.org/published-patterns/value-stream)

{.premise}
There are always a lot of things that could be done, but we can do very little.
It is easy to get lost in the ideas, technical issues and technical debt.
We must make choices in what we do and more importantly things that we won't do.

The intent of writing software is to create Value to people. Anything that is not
of Value is waste.

{.therefore}
Focus on creating Value and how it relates to all the stakeholders and its surroundings.
Ensure that the Value is clear to everyone and fits together with other Values around it.

--- Description ---

With Value there are two important aspects that must be considered Context and the  Evaluator.

We cannot understand Value without the Context where it appears. Value is always
connected to other things of Value. We cannot properly evaluate importance
of something without being able to compare it with something else. Hence
we must have a good understanding of the Value and how it relates to other
things surrounding it.

{.example}
A door is valuable when it is front of a house.
A keyboard is valuable when it is connected to a computer.
A web server is valuable when it serves content to a browser.

Value is not created for anyone, it is created for someone. We should keep in mind
how the Value enhances and improves the life of the person using it.

{.example}
Easy to use GUI interface provides more value to beginners, however it
is slower to use for power-users. A command-line interface is often faster
to use, when learned, but it is more difficult to learn.

Value of something is not a binary scale, there are subtle things that improve or
diminish Value. There are also a lot of things that do not significantly add
to the value.

{.example}
A color blind palette choice can make the software easier to use for people with
color deficieny, but is often overlooked. Never relying on only color to distinguish
visual elements is even better.

{.example}
When writing a prototype or learning, there is little value in good documentation,
good naming conventions and tests. Hence when learning something, do not write
comments or polish it, you can iterate quicker and learn more by skipping those
activities.

There are a many things that can be valuable and it is easy to miss the less
obvious parts. There are things directly related to code:

* [Value in Namespaces](value/namespaces.md) - we need to organize our structures
  to ensure that we have clear distinction between ideas.
* [Value in Objects](value/objects.md) - many of the things we write can be
  directly related to the world.
* [Value in Functions](value/functions.md) - pure functions provide a reliable and
  repeatable computations.
* [Value in Logic](value/logic.md) - describing logical relations between allows
  separating the intent from the way things are computeted.
* [Value in Interaction](value/interaction.md) - interaction between things can
  be just as important as the interacting things themselves.
* [Value in Libraries](value/libraries.md) - separating ideas into libraries creates
  a possibility for reuse.
* [Value in Platform](value/platform.md) - a good platform allows other things
  to be easier and simpler understand.

There are also things related to the Visuals:

* [Value in User Experience](value/ux.md) - a great experience leads to less
  frustration in people.
* [Value in User Interface](value/ui.md) - good interface pleases and guides
  people to making better choices.

--- Evaluation ---

To have Strong Value in your code you should understand:

* when it is being used -- in a rush, at work or at leisure time;
* where it is being used -- in office, on a bus or on a couch;
* how it is being used -- on a desktop computer, on a tablet or a TV;
* who is using it -- a power user, a color blind person, a blind person.
* how it relates to other Values around it, how they complement each other.

--- Finding ---

There are several ways that can be used to find what is Valuable.

Good [User Experience](value/ux.md) and [User Interface](value/ui.md) show
pieces that are valuable to the end-user. Hence by building the "best"
user-interface you can find some of the strong values.

Domain Expert terminology and a Domain Expert can directly show what are
the important pieces.

--- Building ---

When building Strong Values into your system, start from the strongest and
iteratively work towards adding them. [1](https://sites.google.com/a/scrumplop.org/published-patterns/value-stream/greatest-value)

You may not get the perfect separation into Value structures from the start,
but over-time, since you started with the Strongest Value, they will be
iterated the most.

--- Notes ---

The notion of Value very strongly correlates to the notion of Centers in
Christopher Alexanders work. Similarly how centers are composed of other
centers and how centers complement each others, so does Value. [2](http://www.regismedina.com/articles/christopher-alexander-theory-of-incremental-design)

--- References ---

* [Lean Architecture](http://www.leansoftwarearchitecture.com/) by James Coplien
* [Greatest Value](https://sites.google.com/a/scrumplop.org/published-patterns/value-stream/greatest-value) - ScrumPLOP
* [Value Stream](https://sites.google.com/a/scrumplop.org/published-patterns/value-stream) - ScrumPLOP
* [The Nature of Order: The Phenomenon of Life](http://www.natureoforder.com/summarybk1.htm) by Christopher Alexander
* [Christopher Alexander - a theory of incremental design](http://www.regismedina.com/articles/christopher-alexander-theory-of-incremental-design) by Regis Medina
* DDD by Eric Evans