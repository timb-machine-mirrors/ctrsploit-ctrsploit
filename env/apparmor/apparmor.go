package apparmor

import (
	"github.com/ctrsploit/ctrsploit/pkg/apparmor"
	"github.com/ctrsploit/ctrsploit/pkg/lsm"
	"github.com/ctrsploit/sploit-spec/pkg/env/container"
)

func Apparmor() (a container.Apparmor, err error) {
	a = container.Apparmor{
		KernelSupported:  apparmor.IsSupport(),
		ContainerEnabled: apparmor.IsEnabled(),
	}
	if !a.KernelSupported {
		return
	}
	current, err := lsm.Current()
	if err != nil {
		return
	}
	mode, err := apparmor.Mode()
	if err != nil {
		return
	}
	a.Profile = current
	a.Mode = mode
	return
}
