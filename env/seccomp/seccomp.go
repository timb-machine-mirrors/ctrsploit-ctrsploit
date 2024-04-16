package seccomp

import (
	"github.com/ctrsploit/ctrsploit/pkg/seccomp"
	"github.com/ctrsploit/sploit-spec/pkg/env/container"
)

const (
	CommandName = "seccomp"
)

// Seccomp
// reference: https://lwn.net/Articles/656307/
func Seccomp() (machine container.Seccomp, err error) {
	seccompMode, _, err := seccomp.GetStatus()
	if err != nil {
		return
	}
	machine = container.Seccomp{
		KernelSupported:  seccomp.CheckSupported(),
		ContainerEnabled: seccompMode > 0,
		Mode:             seccompMode.String(),
	}
	return
}
