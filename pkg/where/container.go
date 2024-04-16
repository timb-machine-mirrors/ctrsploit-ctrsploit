package where

import (
	"github.com/ctrsploit/ctrsploit/pkg/namespace"
	"github.com/ctrsploit/sploit-spec/pkg/env/container"
)

type Container struct {
}

// IsIn
// We believe that being in a child mount namespace is equivalent to being inside a container.
// This holds true even if you are in an unshare environment or a chroot environment.
func (c Container) IsIn() (in bool, err error) {
	arbitrator, err := namespace.NewInoArbitrator()
	if err != nil {
		return
	}
	level, err := namespace.GetNamespaceLevel(arbitrator, container.NamespaceNameMnt)
	if err != nil {
		return
	}
	if level == container.NamespaceLevelChild {
		in = true
	}
	return
}
