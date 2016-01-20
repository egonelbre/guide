<a class="sha" href="https://github.com/loov/timeclock/tree/c420b57de35a12d2ce78c63cf24c030927bccb5b">c420b57de3</a>

Now that we have represented the `Project` we should
try to expose it. This means having a server and frontend.

For the server we shall go with the standard http library
and for the frontend with plain-old-HTML. We shouldn't use
JavaScript in this project because the users are likely
to have older browsers and this will make the implementation
more likely to work on any platform.

### Setup

I'm going with my usual server setup in `main.go`:

{callout="//"}
```go
package main

import (
	"flag"
	"net/http"
	"os"
)

var (
	addr = flag.String("listen", ":8000", "http server `address`") <1>
)

func main() {
	flag.Parse()

	host, port := os.Getenv("HOST"), os.Getenv("PORT") <2>
	if host != "" || port != "" {
		*addr = host + ":" + port
	}

	assets := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets)) <3>

	http.HandleFunc("/", index) <4>

	log.Println("Starting server on", *addr)
	http.ListenAndServe(*addr, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
```

We add `addr`<1> to allow configuring the listen address.
We override<2> the configuration from environment variables -
this way I can use [gin](https://github.com/codegangsta/gin)
to automatically recompile the server.
We serve `assets` <3> for any static files, such as css.
We show a simple `index.html`<4> page to see whether our server actually
works as intended.

`index.html` contains nothing special:

```html
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<title>Timeclock</title>
	<link rel="stylesheet" href="/assets/css/main.css">
</head>
<body>
	<h1>Timeclock</h1>
</body>
</html>
```

Similarly `/assets/css/main.css` contains minimal reset:

```css
* {
	box-sizing: border-box;
}

body, html {
	margin: 0;
	padding: 0;
	color: #333;

	letter-spacing: 0.02em;

	-ms-word-wrap: break-all;
	word-wrap: break-word;
	-webkit-font-smoothing: antialiased;
}
```

The basic html/css setup might vary. You could use some standard css setup
such as [skeleton.css](http://getskeleton.com/) or
[bootstrap](http://getbootstrap.com/), to speed up your workflow.

We test our server with `gin .` and see whether `localhost:3000` shows something.

### Viewing project info

<a class="sha" href="https://github.com/loov/timeclock/tree/4ef17bc635874534b6474f73384bcabb403bf9b2">4ef17bc635</a>

After we ensure that the server is working, we make the page
show some mock project information. We convert server to use
`html/template`.

We also add a function to handle any internal errors... note it should use minimally anything that we build. Otherwise, if we use a broken thing to display broken thing we might cause very severe errors.

```go
func internalError(w http.ResponseWriter, r *http.Request, err error) {
	message := template.HTMLEscapeString(err.Error())
	page := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<title>Timeclock</title>
	<link rel="stylesheet" href="/assets/css/main.css">
</head>
<body>
	<div class="error">%s</div>
</body>
</html>
`, message)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(page))
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Printf("error parsing template: %v", err)
		internalError(w, r, err)
		return
	}

	example := &project.Project{
		Title:    "Alpha",
		Customer: "ACME",
		Pricing: project.Pricing{
			Hours: 480,
			Price: 1000,
		},
		Description: "Implement views",
		Status:      project.InProgress,
	}

	err = t.Execute(w, example)
	if err != nil {
		log.Printf("error parsing template: %v", err)
		internalError(w, r, err)
		return
	}
}
```

We also adjust the `index.html` to use project information:

```html
<div class="project">
	<div class="title">{{.Title}}</div>
	<table class="info" border="0">
		<tr>
			<td><b>Customer</b></td><td>{{.Customer}}</td>
			<td class="description" rowspan="4">{{.Description}}</td>
		</tr>
		<tr><td><b>Status</b></td><td>{{.Status}}</td></tr>
		<tr><td><b>Hours</b></td><td>{{.Pricing.Hours}}</td></tr>
		<tr><td><b>Price</b></td><td>{{.Pricing.Price}}€</td></tr>
	</table>
</div>
```

We also made some adjustments to the .css to make it look nicer.

![Project View 00](images/project-info-00.png)

Then we add a way to view expenses:

![Project View 01](images/project-info-01.png)

Before proceeding we can ask users whether this is
what they were looking for. This is so we can understand
whether we are on the right track.
