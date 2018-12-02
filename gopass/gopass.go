package gopass

import (
	"fmt"
	"os/exec"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// GoPass object permit to inteact with the GoPass cmd
type GoPass struct{}

var execCommand = exec.Command

// List get all entries
func (GoPass) List() ([]string, error) {
	cmd := execCommand("gopass", "ls", "--flat")

	out, err := cmd.CombinedOutput()

	if err != nil {
		return nil, err
	}

	return strings.Split(string(out), "\n"), nil
}

// GetInfos give all informations about entry
func (GoPass) GetInfos(entry string) (map[string]string, error) {
	cmd := execCommand("gopass", "show", entry)

	out, err := cmd.CombinedOutput()

	if err != nil {
		return nil, err
	}

	data := strings.Split(string(out), "\n")

	result := map[string]string{}
	result["pass"] = data[0]

	var items yaml.MapSlice
	yaml.Unmarshal([]byte(strings.Join(data[1:], "\n")), &items)

	for _, item := range items {
		key, ok := item.Key.(string)
		if !ok {
			continue
		}

		value, ok := item.Value.(string)
		if !ok {
			continue
		}
		result[key] = value
	}

	return result, nil
}

// Clip password to the clipboard
func (GoPass) Clip(entry string) error {
	fmt.Println("Call gopass.Clip")
	cmd := execCommand("gopass", "show", "-c", entry)

	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	return err
}
