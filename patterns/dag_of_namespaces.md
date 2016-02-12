{.premise}
Structures that contain big loops require knowing about the whole structure
to understand it. However many systems require such circular referencing.

{.therefore}
Ensure that there are no circular imports. If circular structures are required,
keep them contained in a single namespace.

--- Description ---

--- Evaluation ---

When namespace X depends on Y, then by deleting X, without any further modification,
Y should still be usable.

--- Building ---

--- Notes ---

--- References ---
