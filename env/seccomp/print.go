package seccomp

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/env/container"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
)

type Result struct {
	Name      result.Title `json:"name"`
	Kernel    item.Bool    `json:"kernel"`
	Container item.Bool    `json:"container"`
	Mode      item.Short   `json:"mode"`
}

func Human(machine container.Seccomp) (human Result) {
	human = Result{
		Name: result.Title{
			Name: "Seccomp",
		},
		Kernel: item.Bool{
			Name:        "Kernel Supported",
			Description: "",
			Result:      machine.KernelSupported,
		},
		Container: item.Bool{
			Name:        "Container Enabled",
			Description: "",
			Result:      machine.ContainerEnabled,
		},
		Mode: item.Short{
			Name:        "Mode",
			Description: "",
			Result:      machine.Mode,
		},
	}
	return
}

func Print() (err error) {
	machine, err := Seccomp()
	if err != nil {
		return
	}
	u := result.Union{
		Machine: machine,
		Human:   Human(machine),
	}
	fmt.Println(printer.Printer.Print(u))
	return
}
