=============
Issue Tracker
=============


Building with value
===================

I've been thinking what is the best way of developing new things for a while.
Finally I think I've found a process and hopefully it will help you:

1. find and understand value,
2. contain,
3. sketch,
4. clarify,
5. solidify,
6. go back to 1.

That's it...

Yes it's stupidly simple, but it actually works quite well once you understand it.
The easiest way to get a feeling for it is to see it in action.

This process is explained with Go, but it works also in other languages.

Premise
=======

As an example case-study we'll take *issue tracker*. Probably you have seen
one and maybe used one, so it's an easy enough example to understand.

We start with a simple premise:

A company of 40 people have problems seeing all the issues that the company
is facing and that need to be solved. Traditional issue trackers don't have
the clarity that they need.

> Note: there are issue trackers that probably do what we are going to build.
> It can make more business sense to buy an Issue Tracker instead of building
> your own. This "Issue Tracker" is for learning purposes and we make
> decisions that may not be the best business decisions, although
> they will provide more learning experience.


Value
=====

> ... starting in The Mist, you are seeking ways to create Value,
> where value is defined as something of worth to some person or set
> of people whom we wish to serve. - [Value Stream - ScrumPLoP](https://sites.google.com/a/scrumplop.org/published-patterns/value-stream)

I'm going to use the word `value` a lot. Why? Because it's the most important
part of the software you are building. If you are not creating value you are
wasting resources: your own time, your money, other people's time...

To "creating something of value" means that people can benefit from it, even
when it's not providing all the value it could provide. The longer you are
not providing value, the longer you don't get feedback of usefulness.

`Value` can take many forms:

* Functionality for an end-user
* Functionality for a developer
* Knowledge from research
* Knowledge from a prototype

It's obvious how functionality helps to create value. Knowledge makes it easier
to deliver more value to the users.

As a thought exercise:

Does "login feature" provide value? Well it's a feature, so it must have some value.
A great question is to ask, has anyone triumphantly said "I logged into the system
today, I got so much done." Not really, so does that mean it doesn't have value?

Logging in is more of an annoyance than something of value, but it does prevent
loss of value. The only goal of it is to protect value. Logging in is not the
only way of protecting value, you could make backups or make it only accessible
from internal network etc.

Logging in has no value, if you don't have any value to protect.
If you have already protected your value, then it also doesn't have value.

There are several important pieces here:
1. Not all features have value.
2. Value is not a constant.
3. Value of things changes.


Issue
=====

1. Finding and understanding value
----------------------------------

What provides value in an "Issue Tracker". Often you can determine the importance
of things by seeing how much it is mentioned.

> "What things are we working on?"
> "I need to finish this issue."
> "I don't have good visibility on the issues."

It's obvious that we have significant value in an "issue".

Why is "issue" valueable?

* It shows what we still need to do.
* It shows what we have done.
* It shows what we are doing.
* They help to prioritize things.


2. Contain it
-------------

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

3. Sketch the code
------------------

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


4. Clarify the code
-------------------

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


5. Solidify the code
--------------------

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


6. Value done
--------------------

Now we have captured something of value in code. It can't be used easily
right now, but we have something that someone would like to use.

It might look that it was a really involved process that created only these
few lines of code. In reality, it's pretty fluid and moves quite quickly. The
only reason it looks involved is because I tried to write down everything I
was thinking while writing the code.