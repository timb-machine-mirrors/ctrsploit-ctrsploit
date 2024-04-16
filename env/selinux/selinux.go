package selinux

import (
	"github.com/ctrsploit/ctrsploit/pkg/selinux"
	"github.com/ctrsploit/sploit-spec/pkg/env/container"
)

func Selinux() (machine container.SELinux, err error) {
	supported, err := selinux.KernelSupported()
	if err != nil {
		return
	}
	machine = container.SELinux{
		KernelSupported:  supported,
		ContainerEnabled: selinux.IsEnabled(),
		Mode:             selinux.Mode().String(),
		MountPoint:       selinux.GetSelinuxMountPoint(),
	}
	return
}
