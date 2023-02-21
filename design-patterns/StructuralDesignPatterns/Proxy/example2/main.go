package main

import (
	"fmt"
)

/*
		Beispiel:

	   	Ich muss einen Befehl ausführen, also muss ich vorher eine neue Terminalsitzung erstellen
	   	und unseren Benutzernamen und das Kommando eingeben
*/

// ITerminal is interface
// It is a public method whose implemented in Terminal (Proxy) and Gopher Terminal
type ITerminal interface {
	Execute(cmd string) (resp string, err error)
}

// GopherTerminal
type GopherTerminal struct {
	// user is a current authorized user
	User string
}

// Execute just runs known commands for current authorized user
func (gt *GopherTerminal) Execute(cmd string) (resp string, err error) {
	// Set "terminal" prefix for output
	prefix := fmt.Sprintf("%s@go_term$:", gt.User)

	// Execute some asked commands
	switch cmd {
	case "run":
		resp = fmt.Sprintf("%s run %s", prefix, gt.User)
	case "man":
		resp = fmt.Sprintf("%s visit 'https://golang.org/doc/' for Golang documentation", prefix)
	default:
		err = fmt.Errorf("%s unknown command", prefix)
	}

	return
}

/* Create proxy to deliver user and commands to specific objects */

// Terminal is a implementation of Proxy (validates and send data to GopherTerminal)
// Before sending commands, user must be authorized
type Terminal struct {
	currentUser    string
	gopherTerminal *GopherTerminal
}

// NewTerminal creates a new instance of a terminal
func NewTerminal(user string) (t *Terminal, err error) {
	// Check user if given correctly
	if user == "" {
		err = fmt.Errorf("user can't be empty")
		return
	}

	// Before we execute user commands: validate current user (if rights are given)
	if authErr := authorizeUser(user); authErr != nil {
		err = fmt.Errorf("(%s) not allowed to use terminal and execute commands", user)
		return
	}

	// Create new instance of terminal and set valid user
	t = &Terminal{currentUser: user}

	return
}

// Execute intercepts execution of command, implements authorizing user, validates it and
// poxing command to real terminal (gopherTerminal) method
func (t *Terminal) Execute(command string) (resp string, err error) {
	// If user allowed to execute sending commands, I can decide which terminal can be used, remote or local etc..
	// creating new instance of terminal, set current user and send user command to execution in terminal
	t.gopherTerminal = &GopherTerminal{User: t.currentUser}

	// For example our proxy can log or output intercepted execution... etc
	fmt.Printf("proxy: intercepted execution of user (%s), asked command (%s)\n", t.currentUser, command)

	// Transfer data to original object and execute command
	if resp, err = t.gopherTerminal.Execute(command); err != nil {
		err = fmt.Errorf("proxy: can't execute command: run, man")
		return
	}

	return
}

// authorize validates user right to execute commands
func authorizeUser(user string) (err error) {
	// Use terminal like a proxy, then intercept user name to validate,
	// if it's allowed to execute commands
	if user != "gopher" {
		// Do some logs, notifications etc...
		err = fmt.Errorf("user %s in black list", user)
		return
	}

	return
}

func main() {
	// Create new instance of Proxy terminal
	t, err := NewTerminal("gopher")
	if err != nil {
		panic(err.Error())
	}

	// Execute user command
	excResp, excErr := t.Execute("run")
	if excErr != nil {
		fmt.Printf("ERROR: %s\n", excErr.Error())
	}

	// Show execution response
	fmt.Println(excResp)
}
