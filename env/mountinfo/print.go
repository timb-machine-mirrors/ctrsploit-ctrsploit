package mountinfo

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result"
	"github.com/moby/sys/mountinfo"
)

type Result struct {
	Name  result.Title
	Infos []String
}
type String struct {
	Content string
}

func (s String) Text() string {
	return s.Content
}
func (s String) Colorful() string {
	return s.Content
}
func (s String) IsEmpty() bool {
	return s.Content == ""
}

func Human(machine []*mountinfo.Info) (human Result) {
	human = Result{
		Name: result.Title{
			Name: "mountinfo",
		},
		Infos: []String{},
	}
	for _, info := range machine {
		human.Infos = append(human.Infos, String{Content: fmt.Sprintf("(%s) %s => %s", info.Source, info.Root, info.Mountpoint)})
	}
	return
}

func Print() (err error) {
	machine, err := MountInfo()
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
