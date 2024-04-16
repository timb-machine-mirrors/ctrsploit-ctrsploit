package namespace

import (
	"github.com/ctrsploit/ctrsploit/pkg/namespace"
	"github.com/ctrsploit/sploit-spec/pkg/env/container"
)

const CommandName = "namespace"

func Namespace() (machine container.Namespace, err error) {
	machine = container.Namespace{Levels: map[string]container.NamespaceLevel{}}
	arbitrator, err := namespace.NewInoArbitrator()
	if err != nil {
		return
	}
	namespaceLevels, names, err := namespace.CheckNamespaceLevel(arbitrator)
	if err != nil {
		return
	}
	for name, level := range namespaceLevels {
		machine.Levels[name] = container.NamespaceLevel(level)
	}
	machine.Names = names
	return
}
