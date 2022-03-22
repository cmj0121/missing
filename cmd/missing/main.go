package main

import (
	"os"

	"github.com/cmj0121/missing"
)

func main() {
	agent := missing.New()
	os.Exit(agent.Run())
}
