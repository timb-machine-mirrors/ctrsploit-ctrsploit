package capability

import (
	"fmt"
	"github.com/containerd/containerd/pkg/cap"
	"github.com/ctrsploit/sploit-spec/pkg/colorful"
	"github.com/ctrsploit/sploit-spec/pkg/env/container"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
)

const (
	standardCaps uint64 = 0xa80425fb
)

type Cap struct {
	SubTitle     result.SubTitle `json:"-"`
	Capabilities item.Short      `json:"capabilities"`
	NotDefault   item.Bool       `json:"not_default"`
	Additional   item.Long       `json:"additional"`
}

type Caps struct {
	Name    result.Title `json:"name"`
	Pid1    Cap          `json:"pid1"`
	Current Cap          `json:"current"`
}

func getInfoFromCaps(caps uint64, subtitle string) (c Cap) {
	c.SubTitle = result.SubTitle{
		Name: subtitle,
	}
	c.Capabilities = item.Short{
		Name:   "capabilities",
		Result: fmt.Sprintf("0x%x", caps),
	}
	c.NotDefault = item.Bool{
		Name:        fmt.Sprintf("Not Equal to Docker's Default Capability (0x%x)", standardCaps),
		Description: fmt.Sprintf("0x%x", caps),
		Result:      caps != standardCaps,
	}
	if caps != standardCaps {
		capsDiff, _ := cap.FromBitmap(caps & (^standardCaps))
		c.Additional = item.Long{
			Name:        "[Additional]",
			Description: "",
			Result:      colorful.O.Danger(fmt.Sprintf("%q", capsDiff)),
		}
	}
	return
}

func Human(machine container.Capability) (h Caps) {
	pid1 := getInfoFromCaps(machine.Pid1, "pid1")
	current := getInfoFromCaps(machine.Self, "current")
	h = Caps{
		Name: result.Title{
			Name: "Capability",
		},
		Pid1:    pid1,
		Current: current,
	}
	return
}

func Print() (err error) {
	machine, err := Capability()
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
