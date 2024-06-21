package capability

import (
	"github.com/containerd/containerd/pkg/cap"
	"github.com/ctrsploit/ctrsploit/pkg/capability"
	"github.com/ctrsploit/sploit-spec/pkg/env/container"
)

const (
	CommandCapabilityName = "capability"
)

func Capability() (caps container.Capabilities, err error) {
	pid1Eff, err := capability.GetPid1Capability(cap.Effective)
	if err != nil {
		return
	}
	currentEff, err := capability.GetCurrentCapability(cap.Effective)
	if err != nil {
		return
	}
	pid1Bnd, err := capability.GetPid1Capability(cap.Bounding)
	if err != nil {
		return
	}
	currentBnd, err := capability.GetCurrentCapability(cap.Bounding)
	if err != nil {
		return
	}
	caps = container.Capabilities{
		Pid1: container.Capability{
			Eff: pid1Eff,
			Bnd: pid1Bnd,
		},
		Self: container.Capability{
			Eff: currentEff,
			Bnd: currentBnd,
		},
	}
	return
}
