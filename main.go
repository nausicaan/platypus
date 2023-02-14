package main

import (
	"fmt"
	"os"

	w "github.com/nausicaan/platypus/workers"
)

const (
	bv     string = "1.0.0"
	reset  string = "\033[0m"
	red    string = "\033[41m"
	green  string = "\033[32m"
	yellow string = "\033[33m"
	halt   string = "program halted "
	zero   string = "Insufficient arguments supplied - " + halt
)

// Launch the program and execute according to the supplied flag
func main() {
	var flag string

	if len(os.Args) == 1 {
		flag = "--zero"
	} else {
		flag = os.Args[1]
	}

	switch flag {
	case "-v", "--version":
		fmt.Println(yellow+"Platypus", green+bv)
	case "-h", "--help":
		helpMenu()
	case "-p", "--plugin":
		w.Plugin()
	case "-t", "--theme":
		w.Theme()
	case "-c", "--core":
		w.Core()
	case "--zero":
		fmt.Println(red, "No flag detected -", halt)
		fmt.Println(reset)
	default:
		fmt.Println(red, "Bad flag detected -", halt)
		fmt.Println(reset)
	}
}

// Print the help information
func helpMenu() {
	fmt.Println(yellow, "\nUsage:", reset)
	fmt.Println("  ./[program] [flag] [target server] [wordpress path] [wordpress url]")
	fmt.Println(yellow, "\nOptions:")
	fmt.Println(green, " -p, --plugin", reset, " Search for Plugin Updates")
	fmt.Println(green, " -t, --theme", reset, "	 Search for Theme Updates")
	fmt.Println(green, " -c, --core", reset, "	 Search for Core Updates")
	fmt.Println(green, " -v, --version", reset, "Display App Version")
	fmt.Println(green, " -h, --help", reset, "	 Help Information")
	fmt.Println(yellow, "\nExample:", reset)
	fmt.Println("  In your WordPress installation folder, run:")
	fmt.Println(green, "\n    ./platypus -p coeurl.dmz /data/www-app/test_blog_gov_bc_ca/current/web/wp test.blog.gov.bc.ca")
	fmt.Println(yellow, "\nHelp:", reset)
	fmt.Println("  For more information go to:")
	fmt.Println(green, "\n    https://github.com/nausicaan/platypus.git")
	fmt.Println(reset)
}
