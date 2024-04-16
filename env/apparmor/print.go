package apparmor

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
	Profile   item.Short   `json:"profile"`
	Mode      item.Short   `json:"mode"`
}

func Human(a container.Apparmor) (r Result) {
	r = Result{
		Name: result.Title{
			Name: "AppArmor",
		},
		Kernel: item.Bool{
			Name:        "Kernel Supported",
			Description: "Kernel enabled apparmor module",
			Result:      a.KernelSupported,
		},
		Container: item.Bool{
			Name:        "Container Enabled",
			Description: "Current container enabled apparmor",
			Result:      a.ContainerEnabled,
		},
	}
	if !a.ContainerEnabled {
		return
	}
	r.Profile = item.Short{
		Name:        "Profile",
		Description: "",
		Result:      a.Profile,
	}
	r.Mode = item.Short{
		Name:        "Mode",
		Description: "",
		Result:      a.Mode,
	}
	return
}

func Print() (err error) {
	m, err := Apparmor()
	if err != nil {
		return
	}
	u := result.Union{
		Machine: m,
		Human:   Human(m),
	}
	fmt.Println(printer.Printer.PrintDropAfterFalse(u))
	return
}
