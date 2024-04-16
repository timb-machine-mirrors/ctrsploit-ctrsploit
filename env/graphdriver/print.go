package graphdriver

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/env/container"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
)

type ResultItem struct {
	Name     result.Title `json:"name"`
	Enabled  item.Bool    `json:"enabled"`
	Used     item.Bool    `json:"used"`
	Number   item.Long    `json:"number"`
	HostPath item.Long    `json:"host_path"`
}

type Result map[string]ResultItem

func graphDriver(name string, graphDriver container.GraphDriver) (human ResultItem) {
	human = ResultItem{
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

func Human(machine container.Filesystem) (human Result) {
	human = Result{
		"overlay":      graphDriver("Overlay", machine.Overlay),
		"devicemapper": graphDriver("DeviceMapper", machine.DeviceMapper),
	}
	return
}

func Print() (err error) {
	machine, err := GraphDrivers()
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
