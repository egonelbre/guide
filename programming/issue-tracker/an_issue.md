Here is our first important question. *"What provides value in an Issue Tracker?"*

We may notice the important pieces from the language used.

    "What are we working on?"
    "I need to finish this issue."
    "I don't have good visibility on the issues."

We may also get this information from a domain expert. In our case it is
obvious, the most important thing is an ***issue***.

Here we should think, why ***issue*** is important:

* It shows what we are doing.
* It shows what we have done.
* It shows what we still need to do.

***Issue*** is definitely a valuable part of Tracker hence we must capture
it in code.


### Contain it

<a class="sha" href="https://github.com/loov/tracker/tree/bd9af4dcda0207555204bad0addcd0ce4d0a61dc">bd9af4dcda</a>

To ensure that we can capture ***issue*** and do not lose the knowledge
about it's importance we should create a namespace for it. That way we make
it significant and important. At the same time we create a locus of attention
which allows to understand and examine the feature wholly and whether it
is complete.

Of course, namespaces are not the only way to contain things - there are also
packages, classes, functions, methods, constraints etc. What you use to contain
will depend on how large, detailed or important the contained thing is.

It is better to start with a notch larger container than is needed, it is not
difficult to make it smaller. However the reverse, moving from smaller container
to larger, is usually more difficult.

***Issue*** is a very important concept, hence we start with a package `issue`.
Our starting folder structure will be:

{callout="//"}
```
/main.go <1>
/issue/
```

It is clear that at somepoint we will need `main.go` <1>, so we can add it now -
although, when we add it doesn't matter.


### Spike it

<a class="sha" href="https://github.com/loov/tracker/tree/9f47cfaaeadbaa999dc36fd238684151bb9bc6e4">9f47cfaaea</a>

We should try to figure out what our `issue` contains. Create `issue/info.go`
with:

{callout="//"}
```go
package issue

type ID int                     <1>
type Status string              <2>

const (
	Created Status = "Created"  <3>
	Closed         = "Closed"
)

type Info struct {              <4>
	ID      ID
	Caption string
	Desc    string
	Status  Status
}
```

We must have an `issue.ID` <1> to uniquely identify an issue. Each issue usually
has an `issue.Status` <2><3> associated with it. We need something to bring all
the attributes together <4>. Keep in mind we are sketching the code and are not
committed to this structure. We are not looking perfection, but rather a
global view how things will work together.

Notice that I don't use long names such as `IssueStatus`, `IssueInfo` because
the namespace for them already contains name `issue`. We should always leverage
our namespace for naming. The full name for them are `issue.Status` and `issue.Info`
which help to clarify.

We also need some way to store and load those issues. The way we store and load
them can change, hence we should abstract this knowledge away. We create an
interface `issue.Manager` for it. We put it into `issue/manager.go`:

{callout="//"}
```go
package issue

type Manager interface {
	Create(info Info) (ID, error)
	Load(id ID) (Info, error)
	Close(id ID) (error)
	List() (issues []Info, error)
}
```

To get a overview how we will use it, we write some usage code into `main.go`:

{callout="//"}
```go
package main

import (
	"fmt"

	"github.com/loov/tracker/issue"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var manager issue.Manager

	id, err := manager.Create(issue.Info{
		Caption: "Hello",
		Desc: "World",
		Status: issue.Created,
	})
	check(err)

	info, err := manager.Load(id)
	check(err)

	fmt.Println(info)

	infos, err := manager.List()
	check(err)
	fmt.Println(infos)
}
```

It doesn't matter in which order you create these pieces, sometimes it is easier
to create the usage code first, other times it is easier to create the implementation
first. The thing that does matter is that both exist to ensure that we have
the implementation details right and that we can integerate it with rest of the
code.

The usage code can also be sketched as a test, this depends on how the sketched
code will be used, how it needs to integerate with the rest of the system and
other factors.

### Gradual stiffening

Notice that we actually don't have any runnable code yet, it's fine, because
until now we were trying to grasp what we are implementing and that all the
pieces work together as intended.

Now we will step-by-step start to flesh out the actual structure, until we
have solid and good runnable code. We are in our beginning stages of our project
so there really isn't much to worry about. We should skim over our code and
notice anything that doesn't feel nice.

The first thing we may notice is `issue.Created`. What would
`info.Status == issue.Created` mean? This suggests that we haven't captured the
intent as well as we should have. Let's refine our sketch, `info.Status == issue.Open`
sounds much better, hence we change `issue/info.go`:

{callout="//"}
```go
const (
	Open Status = "Open"
	Done        = "Done"
)
```

In `main.go` the `manager` doesn't feel solid, it feels like a fuzzy concept
without specific meaning. There probably will be more things that need to "manage"
things. Is there a better name for it?

What does the `manager` do? *"It manages and tracks issues."* Here is a clue
for a nicer name `Tracker`. We shall refine `issue/manager.go` into `issue/tracker.go`
and change:

{callout="//"}
```go
package issue

type Tracker interface {
	Create(info Info) (ID, error)
	Load(id ID) (Info, error)
	Close(id ID) (error)
	List() (issues []Info, error)
}
```

We also do all the necessary adjustments to `main.go`. At the end of this we
should have code that compiles however it's fine if it is not yet completely
bug-free. We will do this in the next step, however gradual stiffening together
with cleanup should be mixed it will always end with a final cleanup pass.

### Cleanup

<a class="sha" href="https://github.com/loov/tracker/tree/56f7a0930c1715deeef9e1cf18924353d4968d44">56f7a0930c</a>

Now we have a good idea about the feature and how to put it into code,
we shall go over and fill in all the missing details and ensure that we
have comments and a few tests and are able to use it in some form.

Here we add a stub implementation for the tracker and then write some tests
for the tracker.

Cleanup code -- means now that you have figured out the sketch, you can
go and trace over it with nice clean looking lines, using the ruler and
delete all the messy bits.

The other thing what we want to do here is to make easier to understand
and ensure whether the code really behaves as it should. In most cases you would
want unit or behavior tests, but they are not the only way. You could also
write property tests. Or write output that could be verified by hand,
if the correct behavior is difficult to describe in code.

Few interesting bits while solidifying code. When you come across questions,
mark them as such. For example while writing the tracker test case, I made a
mistake while writing:

```
	//file: issue/tracker_test.go
	tracker.Close(id)
	// ...
	expect := Info{
		ID:      id,
		Caption: "Caption",
		Desc:    "Desc",
		Status:  Closed,     // <--- error, should be Done
	}
```

I mixed up two things: the method is called "Close" and the resulting
status is "Done". Because I made a mistake while writing this, it suggests to
me that the code is not clear enough... but I'm not sure how to improve it.
It probably isn't that important, so I'll mark it as a TODO and move on
to other things:

```
// file: issue/info.go
const (
	//TODO: should tracker.Create renamed to "Open", because status is Open
	Open = Status("Open") // Open means that the issue needs to be worked on
	//TODO: should "Done" be renamed to "Closed", because tracker has method Close
	Done = Status("Done") // Done means that the issue is completed and delivered
)
```

I could try to figure out this immediately, but I really don't think I have
necessary information right now and I will probably find out the details
while implementing other things.


### Is it done?

Now we have captured something of value in code. It can't be used easily
right now, but we have something that someone would like to use.

It might look that it was a really involved process that created only these
few lines of code. In reality, it's pretty fluid and moves quite quickly. The
only reason it looks involved is because I tried to write down everything I
was thinking while writing the code.