package capability

import (
	"github.com/ctrsploit/ctrsploit/pkg/capability"
	"github.com/ctrsploit/sploit-spec/pkg/env/container"
)

const (
	CommandCapabilityName = "capability"
)

func Capability() (cap container.Capability, err error) {
	pid1, err := capability.GetPid1Capability()
	if err != nil {
		return
	}
	current, err := capability.GetCurrentCapability()
	if err != nil {
		return
	}
	cap = container.Capability{
		Pid1: pid1,
		Self: current,
	}
	return
}
