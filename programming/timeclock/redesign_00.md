<a class="sha" href="https://github.com/loov/timeclock/tree/e165e7904fcad53d98b2f7707f713763372255fa">e165e7904f</a>

After putting this project on hold, due to other responsibilities.
After few weeks I was finally able to meet with a user behind the same desk.
The end result of that discussion was a significant drop in complexity.

The main thing we got out of the discussion was how the workflow of the worker
will look like and how he will use the system.

When the worker starts his day, he will be presented with a screen
(quite likely on a mobile device):

![Start Working](programming/timeclock/images/redesign-00-start-working.png "Start Working")

After selecting a project he chooses the activity he is performing:

![Working](programming/timeclock/images/redesign-00-working.png "Working")

There he can submit any additional information such as add an image of the
progress or report some issue (such as some parts were missing or broken).
After completing the task he can either go select another project or
finish his day with a Day Report.

![Day Report](programming/timeclock/images/redesign-00-day-report.png "Day Report")

We try to put together the day-report from existing information as much as possible.
We cannot simply actively confirm the work, because sometimes a worker might forget
to start his timeclock or alternatively he may be somewhere, where there is no
internet access. There might also be some adjustments necessary to the hours,
either way, this additional reporting step is necessary.

All those work hours must be reviewed by an engineer, whether they are up to bar,
or whether they took longer than they should. Engineers review projects and
every week, rather than every day:

![Week Review](programming/timeclock/images/redesign-00-week-review.png "Week Review")

SIDENOTE: The screen prototypes left out some parts when designing,
because it was easy to imagine them being there, but the exact design
and mock-up wasn't important at that point.

Since all the reviews, whether accepted or declined, need to be processed by
the accounting, there is also a way to view all the pending/accepted/declined
week reviews.

![Weekly Reports](programming/timeclock/images/redesign-00-weekly-reports.png "Weekly Reports")

This new approach is much simpler for the workers to work with and easier to implement,
while mostly delivering the same amount or even more value.

The main thing that was removed, was "materials" tracking. It would be inconvenient
to continously track the materials missing and used, it's easier to do that with talking.

One thing that I additionally created was a "notes.txt" file, to track things
that need to be implemented, however not vital at this point. For example, it
contains that there needs to be few statistics pages to get an overview of projects.