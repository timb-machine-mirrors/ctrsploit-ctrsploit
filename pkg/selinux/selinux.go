package selinux

import (
	"github.com/opencontainers/selinux/go-selinux"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"os"
	"strings"
)

type TypeMode int

func (m TypeMode) String() (mode string) {
	switch m {
	case -1:
		mode = "disabled"
	case 0:
		mode = "permissive"
	case 1:
		mode = "enforcing"
	default:
		mode = "unknown"
	}
	return
}

func KernelSupported() (supported bool, err error) {
	content, err := os.ReadFile("/proc/filesystems")
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	supported = strings.Contains(string(content), "selinuxfs")
	return
}

func IsEnabled() bool {
	return selinux.GetEnabled()
}

func Mode() TypeMode {
	return TypeMode(selinux.EnforceMode())
}
