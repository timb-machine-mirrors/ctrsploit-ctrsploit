package auto

import (
	"github.com/ctrsploit/ctrsploit/env/apparmor"
	"github.com/ctrsploit/ctrsploit/env/capability"
	"github.com/ctrsploit/ctrsploit/env/cgroups"
	"github.com/ctrsploit/ctrsploit/env/graphdriver"
	"github.com/ctrsploit/ctrsploit/env/namespace"
	"github.com/ctrsploit/ctrsploit/env/seccomp"
	"github.com/ctrsploit/ctrsploit/env/selinux"
	"github.com/ctrsploit/ctrsploit/env/where"
	"github.com/ctrsploit/sploit-spec/pkg/env/container"
)

func Auto() (env container.Env, err error) {
	w, err := where.Where()
	if err != nil {
		return
	}
	cap, err := capability.Capability()
	if err != nil {
		return
	}
	aa, err := apparmor.Apparmor()
	if err != nil {
		return
	}
	se, err := selinux.Selinux()
	if err != nil {
		return
	}
	sc, err := seccomp.Seccomp()
	if err != nil {
		return
	}
	ns, err := namespace.Namespace()
	if err != nil {
		return
	}
	cg, err := cgroups.Cgroups()
	if err != nil {
		return
	}
	fs, err := graphdriver.GraphDrivers()
	if err != nil {
		return
	}
	env = container.Env{
		Basic: container.Basic{
			Where:         w,
			KernelVersion: "", //TODO
		},
		LinuxSecurityFeature: container.LinuxSecurityFeature{
			Credential: container.Credential{}, //TODO
			Capability: cap,
			LSM: container.LSM{
				Apparmor: aa,
				SELinux:  se,
			},
			Seccomp:    sc,
			Namespace:  ns,
			CGroups:    cg,
			Filesystem: fs,
		},
		Cluster: container.Cluster{}, //TODO
		Advance: container.Advance{}, //TODO
	}
	return
}
