Here's the translated chat about the project:

C - Customer; A: Analyst.

```
C: Hi, how difficult would it to make a webpage where
   workers can login and add worked hours to different
   projects plus any materials that were used and have
   an option to add pictures.
A: It depends how much different things are added.
   It shouldn't be too difficult, probably there
   already exist things that do that.
C: Yes, there is, it costs about ~150€ per month for
   single user, also they don't solve our problems very
   well. I asked the pricing for some 20 users, it was
   ~6000€.
A: If there isn't any complicated business problems
   and there is no need for old browser support, then it
   should be fairly simple.
C: Wait I'll make a quick diagram, how it should work.
```

![Project sketch](programming/timeclock/images/project-sketch.png)

```
C: Every user gets a login user and password.
   Only supervisor can change things.
A: Do you need any summary/statistics pages?
C: Yes, that would be useful.

A: Who will create the projects?
C: The supervisor.
A: What information needs to be associated with each project?
C: If it's a project with fixed pricing then it calculates
   how much has been spent and whether balance is positive
   or negative.

A: How will the balance be calculated?
C: Hours and materials.
A: Does every worker have the same hour cost?
   Or does it vary between projects?
C: Usually the cost per hour is same, except workshop hours.
A: Always or usually?
C: For different orders it can be different. You should be
   able to change the pricing.

A: Can the whole thing be in English or does it need to
   be localized?
C: Translation is a small problem.
A: It's not about the translating part, rather whether
   there should be built-in support for multiple languages.

A: How should changing pricing look like?
   Should it need some default values that will be used
   and can be overridden depending on the project?
   Will they change during project?
C: The pricing will be fixed when the project is created
   and won't be changed later.

A: Do you need something for sub-projects or sub-tasks?
C: Yes, that might come in handy.
A: Do you need multi-level or is a two-level separation
   sufficient?
   Project -> Objective X -> Task A;
   or Project -> Objective A.
C: Two-level separation would be sufficient.

A: How will the worker mark down material usage?
   By price, amount or something else?
C: When project is created then the material usage is
   already known. But worker should be able to say that
   instead of 4 bolts we actually needed 6.
A: So will he write:
   2 bolts priced at 8€ per piece
   2 bolts, type XYZ and get the price from system?
C: Yup, the latter one.

A: Can the bolt price change during the project and how
   should it be handled, let's say there is:

   1. In a project it's marked that 2 bolts were used.
   2. Then price of that bolt changes.

   Should the price be at the time 1. or 2.
C: At moment 1.

A: How does the worker mark down his status?
   I worked on X for 4hours?
   How should the work hour pricing be calculated?
   Does it have to be legally valid from an accounting
   standpoint or is it only for internal tracking?
C: For workers we only track the hours, accounting
   manages the payments separately. Important are
   materials and cost per hour.
```

This whole discussion took about 1hour. This seems like
a lot of jumping around instead of a rigid discussion, but
it conveys sufficient information about the project to
make a prototype.

### Artifacts

While discussing I was simultanously desiging different
artifacts that need to be tracked:

```
ResourceType (Worker, Material)

[]Resource {
	Name
	ResourceType
	Unit
	UnitPrice
}

Unit (Hour, Piece, Grams, Litre)

Customer {
	Name
	Phone
}

Worker {
	Name
	Supervisor
}

Project {
	Customer
	Title
	Pricing {
		Hours
		Price
	}
	Status
	Description
}

Task {
	Project
	Title
	Description
}

Expense {
	Project
	Task
	Worker
	Date
	Resource
	Amount
	Unit
	UnitPrice
}

Comment {
	Task
	Worker
	Date
	Comment
}

Attachment {
	Task
	Worker
	Date
	Comment
}
```

This isn't some magic syntax, it's pseudo-code of artifacts.
Here I decided to track work-hours and materials as "Expense"
instead of separate entities, it will add some uniformity to the discussion.

### Prototype

I'm sure that we don't have full understanding of the business logic,
but we do understand how things are related. It will be sufficient to create
a prototype.

Note, that it is a concious desicion to stop analysis at this point.
This is the minimal viable product that provides value to the customer.
I'm not sure whether customer would be to imagine all the intricacies
even if I asked them. I might get wrong or guesses instead of good answers.
We need something solid, something that the customer can try out.

There are several ways we could build the prototype: paper mock-ups, wireframes,
some prototyping software, HTML page, actual code.
Use whatever you feel most comfortable fastest with.
In this case the webpage seems pretty trivial so there is no need for paper
mock-ups or wireframes. We can go directly to HTML or code.