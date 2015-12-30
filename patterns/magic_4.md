Whenever we are describing or explaining concepts there is a strict limit
how much we can keep in our heads at once. However, if we use few things to
explain things, then we increase the number of "abstraction layers".

Therefore

When describing an concept use 4+-1 other ideas to explain it.

### Rationale

People have a well known limit in their processing capabilities called magic number four.
(https://en.wikipedia.org/wiki/Working_memory#Capacity) (http://www.ncbi.nlm.nih.gov/pmc/articles/PMC2864034/).
Anytime we exceed this limit then the explanation of the concept doesn't fit into our working memory.

Going the other way, using as few other ideas, doesn't work either. If we
need in total 5 things to explain something, and we use 2 things to explain each
thing. This in the best case would create two additional ideas or concepts.

DIAGRAMS

### Subtleties

This 4+-1 doesn't mean that you cannot have more than 5 lines in your method.
Or that you cannot have more than 5 ideas in your method. The main thing we are
concerned about is how many concepts does person have hold in their head to
understand.

By limiting the number lines or ideas per function can do harm. One of such
example is a game loop (link to Carmack).


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

If grouping is necessary, we can use empty blocks:

```
while(!exited){
	{
		update mouse
		update systems
		update entities
	}
	{
		draw entites
		draw hud
	}
	blit to screen
}
```

### Extremes

This doesn't mean there are cases where that this bounds shouldn't be broken.
Sometimes the essential complexity of the idea requires multiple things. Sometimes it's not necessary to use more ideas.

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