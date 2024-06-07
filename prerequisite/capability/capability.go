package capability

import (
	"github.com/containerd/containerd/pkg/cap"
	"github.com/ctrsploit/ctrsploit/pkg/capability"
	"github.com/ctrsploit/sploit-spec/pkg/prerequisite"
	"github.com/ssst0n3/awesome_libs/slice"
)

type Capability struct {
	ExpectedCapability string
	prerequisite.BasePrerequisite
}

var (
	ContainsCapSysAdmin = Capability{
		ExpectedCapability: "CAP_SYS_ADMIN",
		BasePrerequisite: prerequisite.BasePrerequisite{
			Name: "CAP_SYS_ADMIN",
			Info: "Container with cap_sys_admin is dangerous",
		},
	}
	ContainsCapDacReadSearch = Capability{
		ExpectedCapability: "CAP_DAC_READ_SEARCH",
		BasePrerequisite: prerequisite.BasePrerequisite{
			Name: "CAP_DAC_READ_SEARCH",
			Info: "Container with cap_dac_read_search is escapable",
		},
	}
)

func (p *Capability) Check() (err error) {
	err = p.BasePrerequisite.Check()
	if err != nil {
		return
	}
	var pid1, self bool
	{
		caps, err := capability.GetPid1Capability()
		if err != nil {
			return err
		}
		capsParsed, _ := cap.FromBitmap(caps)
		pid1 = slice.In(p.ExpectedCapability, capsParsed)
	}
	{
		caps, err := capability.GetCurrentCapability()
		if err != nil {
			return err
		}
		capsParsed, _ := cap.FromBitmap(caps)
		self = slice.In(p.ExpectedCapability, capsParsed)
	}
	p.Satisfied = pid1 || self
	return
}
