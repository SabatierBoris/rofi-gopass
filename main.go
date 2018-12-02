package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/SabatierBoris/rofi-gopass/gopass"
	"github.com/SabatierBoris/rofi-gopass/rofi"
)

var execCommand = exec.Command

func main() {
	//TODO Parse configuration with viper/cobra

	gp := gopass.GoPass{}
	items, _ := gp.List()

	rofi := rofi.Rofi{
		Title: "Password",
		Items: items,
		Actions: map[rofi.Command]func(string) error{
			rofi.Main: func(param string) error {
				infos, err := gp.GetInfos(param)
				if err != nil {
					return err
				}
				return autoType(infos)
			},
			rofi.Alt1: func(param string) error {
				return gp.Clip(param)
			},
		},
	}

	fmt.Println(rofi.Run())
}

func autoType(infos map[string]string) error {
	fmt.Println(infos)
	commands, ok := infos["autotype"]

	if !ok {
		commands = "username :tab pass"
	}

	for _, command := range strings.Split(commands, " ") {
		switch command {
		case ":tab":
			cmd := execCommand("xdotool", "key", "Tab")
			_, err := cmd.CombinedOutput()
			if err != nil {
				return err
			}
		case ":return":
			cmd := execCommand("xdotool", "key", "Return")
			_, err := cmd.CombinedOutput()
			if err != nil {
				return err
			}
		default:
			data, ok := infos[command]
			if !ok {
				return fmt.Errorf("%s not found", command)
			}

			cmd := execCommand("xdotool", "type", data)
			_, err := cmd.CombinedOutput()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
