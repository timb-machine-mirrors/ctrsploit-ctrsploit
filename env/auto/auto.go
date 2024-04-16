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
)

func Auto() (err error) {
	_ = where.Where()
	_ = apparmor.Print()
	_ = selinux.Print()
	_ = capability.Print()
	_ = cgroups.Print()
	_ = graphdriver.Print()
	_ = namespace.Print("")
	_ = seccomp.Print()
	return
}
