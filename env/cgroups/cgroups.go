package cgroups

import (
	v1 "github.com/ctrsploit/ctrsploit/pkg/cgroup/v1"
	"github.com/ctrsploit/ctrsploit/pkg/cgroup/version"
	"github.com/ctrsploit/sploit-spec/pkg/env/container"
)

const CommandCgroupsName = "cgroups"

func Cgroups() (cgroups container.CGroups, err error) {
	cgroups = container.CGroups{
		Subsystems:         []string{},
		TopLevelSubSystems: []string{},
	}
	if version.IsCgroupV1() {
		cgroups.Version = container.CgroupsV1
	}
	if version.IsCgroupV2() {
		cgroups.Version = container.CgroupsV2
	}

	var c v1.CgroupV1
	subsystemsSupport, err := c.ListSubsystems("/proc/1/cgroup")
	if err != nil {
		return
	}
	for subsystemName, subsystemPath := range subsystemsSupport {
		cgroups.Subsystems = append(cgroups.Subsystems, subsystemName)
		if c.IsTop(subsystemPath) {
			cgroups.TopLevelSubSystems = append(cgroups.TopLevelSubSystems, subsystemName)
		}
	}
	return
}
