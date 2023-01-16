package main

import (
	w "github.com/nausicaan/wp-checker/workers"
)

var BuildVersion = "1.0.0"

// Launch the program and execute according to the supplied flag
func main() {
	w.Plugin()
}
