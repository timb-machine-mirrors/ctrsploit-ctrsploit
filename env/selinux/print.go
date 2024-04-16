package selinux

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/env/container"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
)

type Result struct {
	Name       result.Title `json:"name"`
	Enabled    item.Bool    `json:"enabled"`
	Mode       item.Short   `json:"mode"`
	MountPoint item.Short   `json:"mount_point"`
}

func Human(machine container.SELinux) (human Result) {
	human = Result{
		Name: result.Title{
			Name: "SELinux",
		},
		Enabled: item.Bool{
			Name:        "Enabled",
			Description: "",
			Result:      machine.ContainerEnabled,
		},
		Mode: item.Short{
			Name:        "Mode",
			Description: "",
			Result:      machine.Mode,
		},
		MountPoint: item.Short{
			Name:        "Mount point",
			Description: "",
			Result:      machine.MountPoint,
		},
	}
	return
}

func Print() (err error) {
	machine, err := Selinux()
	if err != nil {
		return
	}
	u := result.Union{
		Machine: machine,
		Human:   Human(machine),
	}
	fmt.Println(printer.Printer.PrintDropAfterFalse(u))
	return
}
