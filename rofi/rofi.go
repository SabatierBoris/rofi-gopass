package rofi

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
	"syscall"
)

var execCommand = exec.Command //TODO For unit test

//Command is the enum of all rofi user action possible
type Command int

const (
	// Main is the main rofi action
	Main Command = iota
	// Alt1 is the first alternative rofi action
	Alt1
	// Alt2 is the second alternative rofi action
	Alt2
	// Alt3 is the third alternative rofi action
	Alt3
	// Alt4 is the fourth alternative rofi action
	Alt4
	// Alt5 is the fith alternative rofi action
	Alt5
	// Alt6 is the sixth alternative rofi action
	Alt6
	// Alt7 is the seventh alternative rofi action
	Alt7
	// Alt8 is the eighth alternative rofi action
	Alt8
	// Alt9 is the ninth alternative rofi action
	Alt9
	// Alt10 is the tenth alternative rofi action
	Alt10
)

// Rofi object permit to inteact with the Rofi
type Rofi struct {
	Title   string
	Items   []string
	Actions map[Command]func(string) error
}

//Run rofi with menu fill with items
func (r Rofi) Run() error {
	data, command, err := r.display()
	if err != nil {
		return err
	}

	action, ok := r.Actions[command]
	if !ok {
		return fmt.Errorf("unknown action")
	}

	return action(data)
}

func (r Rofi) display() (string, Command, error) {
	cmd := execCommand("rofi", "-dmenu", "-p", r.Title)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", Main, err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, strings.Join(r.Items, "\n"))
	}()

	out, err := cmd.CombinedOutput()

	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				exitCode := status.ExitStatus()
				if exitCode == 1 {
					// Abort by the user
					return "", Main, err
				}

				if exitCode < 10 && exitCode > 19 {
					// Unhandled exit code
					return "", Main, err
				}
				return strings.Trim(string(out), "\n"), Command(exitCode - 9), nil

			}
		}
		return "", Main, err
	}

	return strings.Trim(string(out), "\n"), Main, nil
}
