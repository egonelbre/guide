{.epigraph}
> The scientists of today think deeply instead of clearly.
> One must be sane to think clearly, but one can think deeply and be quite insane.
>
> {.quote}
> Nikola Tesla

{.premise}
Whenever we are describing or explaining concepts there is a strict limit
how much we can keep in our heads at once. However, if we use few things to
explain things, then we increase the number of "abstraction layers".

{.therefore}
When describing an concept use 4±1 other ideas to explain it.

--- Rationale ---

People have a well known limit in their processing capabilities called magic number four.
(https://en.wikipedia.org/wiki/Working_memory#Capacity) (http://www.ncbi.nlm.nih.gov/pmc/articles/PMC2864034/).
Anytime we exceed this limit then the explanation of the concept doesn't fit into our working memory.

Going the other way, using as few other ideas, doesn't work either. If we
need in total 5 things to explain something, and we use 2 things to explain each
thing. This in the best case would create two additional ideas or concepts.

DIAGRAMS

--- Evaluation ---

Counting `ideas` properly is difficult, but it is good principle to think about.
Here the proof is in the pudding. If you cannot keep the ideas in your head,
this means that you have exceeded the limit.

To test this, use Naked-CRC (link to Legacy Software) cards
(or equivalent for your paradigm) to describe the behavior of the system to
another programmer. Communicate the names/ideas verbally, but keep around a
generic non-colored thing to represent them. If the programmer can explain
afterwards, then it probably is below the average processing limit.

There's also a phone test, see whether you can explain your idea over
the phone and properly carry it over. This requires the other person to hold
all the ideas in their head.

--- Subtleties ---

This 4+-1 doesn't mean that you cannot have more than 5 lines in your method.
Or that you cannot have more than 5 ideas in your method. The main thing we are
concerned about is how many concepts does person have hold in their head to
understand.

By limiting the number lines or ideas per function can do harm. One of such
examples is a game loop (link to Carmack).

```
while(!exited){
	update mouse
	update systems
	update entities
	draw entites
	draw hud
	blit to screen
}
```

```
update(){
	update mouse
	update systems
	update entities
}

draw(){
	draw entites
	draw hud
}

while(!exited){
	update()
	draw()
	blit to screen
}
```

We have increased our LOC, without actually reducing the complexity. The first
case is perfectly understandable, we need to keep in mind that we are in
the `while` loop and any specific action we are doing.

Grouping and comments would have kept the game-loop more concise and easier
to follow.

```
while(!exited){
	{ // updating
		update mouse
		update systems
		update entities
	}
	{ // drawing
		draw entites
		draw hud
	}
	{ // loop update
		blit to screen
	}
}
```

It should also be noted that when a single class, type or function has
two ideas together wihout unifying idea, it should be counted as two ideas.

EXAMPLE NEEDED

--- Extremes ---

This doesn't mean there are cases where that this bounds shouldn't be broken.
Sometimes the essential complexity of the idea requires multiple things.
Sometimes it's not necessary to use more ideas.

Distributed algorithms can easily violate these rules, because you need to account
for different kinds of failures and failovers. We can take (quorum) algorithm.
In those cases human verification becomes difficult or near impossible.
We can use program verification to prove properties of the system.
Alternatively we can substitute them with easier to understand algorithms,
such as Raft, that may not have as good properties, but we are less likely to
make a mistake in it's implementation.

Now when you don't have 3 things to explain your idea, which can often be the
case when implementing trivial things, there's not point in inventing new things
to satisfy the constraint.

--- Notes ---

It should be noted that good programmers tend have a better skill in moving
around these "idea groupings" (LINK CHRIS GRANGER), however, good algorists tend
to be better at being able to hold more things in their head (TODO VERIFY).
It's easy to see why, software development mostly requires creating, manipulating
and making groups of ideas interact; alternatively figuring out why
multiple groups of ideas do not interact properly. Whereas algorithms require
clever manipulation of concurrent ideas.