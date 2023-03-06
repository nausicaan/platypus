package workers

import (
	"io"
	"log"
	"os"
	"os/exec"
)

var (
	server, site            string
	sender, recipient, user string = "", "", ""
)

// Test if the server value passed to the program is on the list
func contains() bool {
	server = os.Args[2]
	for _, v := range servers {
		if v == server {
			return true
		}
	}
	return false
}

// Run a terminal command, then capture and return the output as a byte
func byteout(name string, task ...string) []byte {
	lpath, err := exec.LookPath(name)
	problem(err)
	osCmd, _ := exec.Command(lpath, task...).CombinedOutput()
	return osCmd
}

// Check for errors, halt the program if found, and log the result
func problem(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Run the Linux mail command and email the result to the configured recipent(s)
func mailman(list string) {
	cmd := exec.Command("mail", "-s", "WordPress updates for "+site, "-r", "Delivery Cactuar <"+sender+">", recipient)
	stdin, err := cmd.StdinPipe()
	problem(err)

	go func() {
		defer stdin.Close()
		_, err := io.WriteString(stdin, "Below is the current list of plugins requiring updates for "+site+". Have a magical day!\n\n"+list)
		problem(err)
	}()

	out, _ := cmd.CombinedOutput()

	journal("Email sent" + string(out))
}

// Pipe together commands using the exec.Command function
func concat(method, flag, task, pipe string) []byte {
	cmd := exec.Command(method, flag, task)
	stdin, err := cmd.StdinPipe()
	problem(err)

	go func() {
		defer stdin.Close()
		_, err := io.WriteString(stdin, pipe)
		problem(err)
	}()

	out, _ := cmd.CombinedOutput()
	return out
}

func cleanup(cut string) {
	problem(os.Remove(cut))
}
