{.epigraph}
> Speaking as a builder, if you start something, you must have a vision of the
> thing which arises from your instinct about preserving and enhancing
> what is there... If you're working correctly, the feeling doesn't wander about.
Quote: -- Christopher Alexander

The most important part of programming is that you cannot do everything at once.
The human mind is fairly limited how much it can hold in its head at a time.
To combat this limitation we need to build things piece-by-piece to ensure
that everything will fit together nicely.

Following patterns are invaluable to building something with ease and clearity.


### Structure Follows Value

We dedicated a lot of time explaining what value is and why it is important.
To ensure that we actually build something of value, we need to understand
the value it provides and how it provides it.

By dissecting and deconstructing value we are able to find boundaries that
align with the problem domain.


### Spiking

To ensure that we get a good global view how things will interact we
put up several spikes and imagine everything working together. This helps
to iterate much faster on the larger view without being slowed down by
techinical details.


### Gradual Stiffening

To accomodate all the issues that arises from integrating different pieces
we gradually make the code more solid. This allows better detection on any
problems that may arise.


### Cleanup

This always happens in tandem with gradual stiffening and should be done to
finish that everything is correct and in order.


### Putting them together

It might not be easy to comprehend this without practical examples, but the
easiest way to understand is to imagine this with drawing.

You don't start drawing from the top-left corner and proceed to the right
side, and then the next line. First you figure out, what you are drawing
and what are the things you are going to draw (structure follows value).
Then you do few rough sketches until you see that the drawing will work nicely (spiking).
Then you start to refine the sketches, then you draw properly with a pen (gradual stiffening).
Then you clean all the sketch lines and color it (cleanup).

The principle for programming is the same.