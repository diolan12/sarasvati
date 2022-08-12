package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func serve(args []string) {
	color.Green("Serving yummy data ...")
	color.Yellow("CTRL+C to stop")

	cmd := exec.Command("php", "-S", "localhost:5412", "-t", "web")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
