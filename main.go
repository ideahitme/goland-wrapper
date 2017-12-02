package main

import (
	"os"
	"os/exec"
)

const pathToGoland = "/Applications/GoLand/Contents/MacOS/goland"

func main() {
	d, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cmd := exec.Command(pathToGoland, d)

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
