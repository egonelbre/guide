package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func compile(input, output string) {
	cmd := exec.Command("mmark",
		"-head", filepath.Join("assets", "head.html"),
		//"-css", filepath.Join("assets", "main.css"),
		input,
	)
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(output, out, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	compile("programming/index.md", "programming.html")
}
