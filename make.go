package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func compile(input, output string) {
	fmt.Println("===== ", input)
	cmd := exec.Command("mmark",
		"-head", filepath.Join("assets", "head.html"),
		//"-css", filepath.Join("assets", "main.css"),
		input,
	)
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile(output, out, 0755)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	compile("programming/_index.md", "programming.html")
	compile("patterns/_index.md", "patterns.html")
	compile("coding/_index.md", "coding.html")
	compile("algorithms/_index.md", "algorithms.html")
	compile("software/_index.md", "software.html")
}
