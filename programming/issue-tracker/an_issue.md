What provides value in an "Issue Tracker". Often you can determine the importance of things by seeing how much it is mentioned.

> "What things are we working on?"
> "I need to finish this issue."
> "I don't have good visibility on the issues."

It's obvious that we have significant value in an "issue".

Why is "issue" valueable?

* It shows what we still need to do.
* It shows what we have done.
* It shows what we are doing.
* They help to prioritize things.

### Contain it

<a class="sha" href="https://github.com/loov/tracker/tree/bd9af4dcda0207555204bad0addcd0ce4d0a61dc">bd9af4dcda</a>

When something is important we want to capture it in code. We want to keep it
together and not break it up into several pieces. By breaking things apart
we are also breaking the story it tells and how it interacts.

To ensure that important things stay consistent we try to contain them.

There are several ways to contain things: namespaces, modules, packages,
classes, functions, methods etc. What you use to contain depends on how large
the things are and what language you are using.

In Go we would start with a package. It provides the largest scope we would
need, of course if we need, we can always merge it into another package.
Breaking things apart is much more difficult than merging together.

We start with a package issue:

```
/main.go
/issue/
```

We also add a main.go for sketching out our code.

### Sketch the code

<a class="sha" href="https://github.com/loov/tracker/tree/9f47cfaaeadbaa999dc36fd238684151bb9bc6e4">9f47cfaaea</a>

Here we try to figure out what are the essential parts of an issue. Here's
what I've come up with:

```
// file: issue/info.go
package issue

type ID int

type Status string

const (
	Created Status = "Created"
	Closed         = "Closed"
)

type Info struct {
	ID      ID
	Caption	string
	Desc    string
	Status  Status
}
```

Remember we are still sketching the code and we should be free to change things.

Notice that I don't use long names such as `IssueStatus`, `IssueInfo` because
the container for them already contains name `issue`. The full names are
`issue.ID`, `issue.Status`... etc.

Of course we also need to save and load the issues somehow. I really don't care
how exactly, but it's important what we can do with the issues.

How to call "the thing that manages issues", not sure... an easy way out would
be to call it "issue.Manager".

```
// file: issue/manager.go
package issue

type Manager interface {
	Create(info Info) (ID, error)
	Load(id ID) (Info, error)
	Close(id ID) (error)
	List() (issues []Info, error)
}
```

Finally we should try to sketch out how we use it:

```
// file: main.go
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

One important part here is that the order in which I create those doesn't matter,
but it is important that I have both sketches of both example of usage and the thing I'm using.

The usage code could also be written immediately as a test, but I usually don't
want to verify it in isolation, but rather how well it fits together with
rest of the system. So I usually leave the tests for later, unless I'm working
with clearly defined behavior.


### Clarify the code

Notice that we actually don't have runnable code yet, and it's fine, because
we are trying to sketch out how all the pieces interact and make sure that the
code is understandable.

The previous step was simply a "sketch", now we start to refine it and try to
find the places that don't have clarity.

Skimming over the code I find several places that bother me.

What would `if info.Status == issue.Created {` mean? I'm not sure what the
status "created" means, it doesn't capture the intent. I guess
`if info.Status == issue.Open {` would have more clarity, so we refine our sketch:

```
// file: issue/info.go
const (
	Open Status = "Open"
	Done        = "Done"
)
```

In `main.go` the `manager` name will be very confusing, because I suspect
there could be a lot of things that "manage" other things. Is there a better name for
it?

What does the `manager` do? "It manages and tracks issues." Here is a clue for
a nicer name: "Tracker", so we refine it as well:

```
// file: issue/tracker.go
package issue

type Tracker interface {
	Create(info Info) (ID, error)
	Load(id ID) (Info, error)
	Close(id ID) (error)
	List() (issues []Info, error)
}
```

Obviously adjusting the main.go as necessary.


### Solidify the code

<a class="sha" href="https://github.com/loov/tracker/tree/56f7a0930c1715deeef9e1cf18924353d4968d44">56f7a0930c</a>

Now we have a good idea about the feature and how to put it into code,
we shall go over and fill in all the missing details and ensure that we
have comments and a few tests and are able to use it in some form.

Here we add a stub implementation for the tracker and then write some tests
for the tracker.

Solidifying code -- means now that you have figured out the sketch, you can
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