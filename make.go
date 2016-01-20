package main

import (
	"fmt"
	"io/ioutil"

	"github.com/loov/mark"
	"github.com/loov/mark/html"
)

func compile(input, output string) {
	sequence, errs := mark.ParseFile(mark.Dir("."), input)
	if len(errs) > 0 {
		fmt.Printf("\n\n= %s\n", input)
		for _, err := range errs {
			fmt.Println(err)
		}
	}

	result := html.Convert(sequence)

	err := ioutil.WriteFile(output, []byte(`
<html>
<head>
	<title>A Guide to ...</title>
	<link rel="stylesheet" href="assets/main.css">
</head>
<body>`+result+`</body></html>`), 0755)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	compile("programming/_index.md", "programming.html")
	compile("patterns/_index.md", "patterns.html")
	compile("coding/_index.md", "coding.html")
	compile("algorithms/_index.md", "algorithms.html")
	compile("software/_index.md", "software.html")
}
