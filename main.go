package main

import (
	"github.com/davidwadden/conductor-go-workers/conductor"
)

func main() {
	conductor.NewConductorHttpClient("http://localhost:8080")
}
