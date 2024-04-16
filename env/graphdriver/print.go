package graphdriver

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/env/container"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
)

type Result struct {
	Name     result.Title `json:"name"`
	Enabled  item.Bool    `json:"enabled"`
	Used     item.Bool    `json:"used"`
	Number   item.Long    `json:"number"`
	HostPath item.Long    `json:"host_path"`
}

func graphDriver(name string, graphDriver container.GraphDriver) (human Result, err error) {
	human = Result{
		Name: result.Title{
			Name: name,
		},
		Enabled: item.Bool{
			Name:        "Enabled",
			Description: "",
			Result:      graphDriver.Loaded,
		},
		Used: item.Bool{
			Name:        "Used",
			Description: "",
			Result:      graphDriver.Used,
		},
		Number: item.Long{
			Name:        "The number of graph driver mounted",
			Description: "equal to the number of containers",
			Result:      fmt.Sprintf("%d", graphDriver.Refcnt),
		},
		HostPath: item.Long{
			Name:        "The host path of container's rootfs",
			Description: "",
			Result:      graphDriver.HostPath,
		},
	}
	return
}

func Human(machine container.Filesystem) (human map[string]Result, err error) {
	o, err := graphDriver("Overlay", machine.Overlay)
	if err != nil {
		return
	}
	d, err := graphDriver("DeviceMapper", machine.DeviceMapper)
	if err != nil {
		return
	}
	human = map[string]Result{
		"overlay":      o,
		"devicemapper": d,
	}
	return
}

func Print() (err error) {
	machine, err := GraphDrivers()
	if err != nil {
		return
	}
	human, err := Human(machine)
	if err != nil {
		return
	}
	u := result.Union{
		Machine: machine,
		Human:   human,
	}
	fmt.Println(printer.Printer.Print(u))
	return
}
