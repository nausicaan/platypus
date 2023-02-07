package workers

import (
	"io"
	"log"
	"os"
	"os/exec"
)

const (
	sender, recipient string = "interuser@cactuar.dmz", "byron.stuike@gov.bc.ca,jeff.stewart@gov.bc.ca,david.kelsey@gov.bc.ca"
)

// Allow only a predefined set of servers
var servers = []string{"cactuar.dmz", "coeurl.dmz", "chimera.dmz", "moogle.dmz", "moblin.dmz", "mimic.dmz"}

// Test if the server value passed to the program is on the list
func contains() bool {
	for _, v := range servers {
		if v == server {
			return true
		}
	}
	return false
}

// Execute terminal commands and return a byte variable
func returnByte(cmd *exec.Cmd) []byte {
	output, err := cmd.Output()
	errors(err)
	return output
}

// Check for errors, halt the program if found, and log the result
func errors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Run the Linux mail command and email the result to the configured recipent(s)
func mailman(list string) {
	cmd := exec.Command("mail", "-s", "WordPress updates for "+site, "-r", "Delivery Cactuar <"+sender+">", recipient)
	stdin, err := cmd.StdinPipe()
	errors(err)

	go func() {
		defer stdin.Close()
		_, err := io.WriteString(stdin, "Below is the current list of plugins requiring updates for "+site+". Have a magical day!\n\n"+list)
		errors(err)
	}()

	out, err := cmd.CombinedOutput()
	errors(err)

	Logging("Email sent" + string(out))
}

// Pipe together commands using the exec.Command function
func concat(method, flag, task, pipe string) []byte {
	cmd := exec.Command(method, flag, task)
	stdin, err := cmd.StdinPipe()
	errors(err)

	go func() {
		defer stdin.Close()
		_, err := io.WriteString(stdin, pipe)
		errors(err)
	}()

	out, err := cmd.CombinedOutput()
	errors(err)
	return out
}

func cleanup(cut string) {
	errors(os.Remove(cut))
}
