package cgroups

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/env/container"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
)

type Result struct {
	Name result.Title `json:"name"`
	V1   item.Bool    `json:"v1"`
	V2   item.Bool    `json:"v2"`
	Sub  item.Long    `json:"sub"`
	Top  item.Long    `json:"top"`
}

func Human(machine container.CGroups) (human Result) {
	human = Result{
		Name: result.Title{
			Name: "CGroups",
		},
		V1: item.Bool{
			Name:        "v1",
			Description: "",
			Result:      machine.Version == container.CgroupsV1,
		},
		V2: item.Bool{
			Name:        "v2",
			Description: "",
			Result:      machine.Version == container.CgroupsV2,
		},
	}
	if machine.Version == container.CgroupsV1 {
		human.Sub = item.Long{
			Name:        "sub systems",
			Description: "",
			Result:      fmt.Sprintf("%+q", machine.Subsystems),
		}
		human.Top = item.Long{
			Name:        "top level subsystems",
			Description: "",
			Result:      fmt.Sprintf("%+q", machine.TopLevelSubSystems),
		}
	}
	return
}

func Print() (err error) {
	m, err := Cgroups()
	if err != nil {
		return
	}
	u := result.Union{
		Machine: m,
		Human:   Human(m),
	}
	fmt.Println(printer.Printer.Print(u))
	return
}
