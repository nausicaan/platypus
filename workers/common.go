package workers

import (
	"io"
	"log"
	"os"
	"os/exec"
)

const (
	sender, recipient = "interuser@cactuar.dmz", "byron.stuike@gov.bc.ca,jeff.stewart@gov.bc.ca,david.kelsey@gov.bc.ca"
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
func rtnByte(cmd *exec.Cmd) []byte {
	output, err := cmd.Output()
	trouble(err)
	return output
}

// Check for errors, halt the program if found, and log the result
func trouble(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Run the Linux mail command and email the result to the configured recipent(s)
func mailman(list string) {
	cmd := exec.Command("mail", "-s", "WordPress updates for "+site, "-r", "Delivery Cactuar <"+sender+">", recipient)
	stdin, err := cmd.StdinPipe()
	trouble(err)

	go func() {
		defer stdin.Close()
		_, err := io.WriteString(stdin, "Below is the current list of plugins requiring updates for "+site+". Have a magical day!\n\n"+list)
		trouble(err)
	}()

	out, err := cmd.CombinedOutput()
	trouble(err)

	Logging("Email sent" + string(out))
}

// Pipe together commands using the exec.Command function
func concat(method, flag, task, pipe string) []byte {
	cmd := exec.Command(method, flag, task)
	stdin, err := cmd.StdinPipe()
	trouble(err)

	go func() {
		defer stdin.Close()
		_, err := io.WriteString(stdin, pipe)
		trouble(err)
	}()

	out, err := cmd.CombinedOutput()
	trouble(err)
	return out
}

func cleanup(cut string) {
	trouble(os.Remove(cut))
}
