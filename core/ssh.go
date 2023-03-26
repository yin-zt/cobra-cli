package core

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"strings"
)

func (this *Common) Ssh(host, username, password, port, command string) {

	if port == "" {
		port = "22"
	}

	// Set up SSH client configuration
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the remote host
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", host, port), config)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	defer conn.Close()

	// Open a session on the remote host
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	// Set up a pipe for capturing the command output
	out, err := session.StdoutPipe()
	if err != nil {
		log.Fatalf("Failed to create stdout pipe: %s", err)
	}

	// Start the command execution
	//err = session.Start("top -b -n 1")
	err = session.Start("ls -la ")
	if err != nil {
		log.Fatalf("Failed to start command: %s", err)
	}

	// Read the command output and print it to the console
	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		fmt.Println(strings.TrimSpace(scanner.Text()))
	}

	// Wait for the command to finish
	err = session.Wait()
	if err != nil {
		log.Fatalf("Failed to wait for command: %s", err)
	}

}
