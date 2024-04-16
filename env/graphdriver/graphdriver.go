package graphdriver

import (
	"github.com/ctrsploit/ctrsploit/pkg/graphdriver"
	"github.com/ctrsploit/ctrsploit/pkg/graphdriver/devicemapper"
	"github.com/ctrsploit/ctrsploit/pkg/graphdriver/overlay"
	"github.com/ctrsploit/sploit-spec/pkg/env/container"
)

const CommandName = "graphdriver"

func GraphDriver(g graphdriver.Interface) (graphDriver container.GraphDriver, err error) {
	err = g.Init()
	if err != nil {
		return
	}
	enabled, err := g.IsEnabled()
	if err != nil {
		return
	}

	var used bool
	var number int
	var hostPath string
	if enabled {
		used, err = g.IsUsed()
		if err != nil {
			return
		}
		number, err = g.Number()
		if err != nil {
			return
		}
		hostPath, err = g.HostPathOfCtrRootfs()
		if err != nil {
			return
		}
		graphDriver = container.GraphDriver{
			Loaded:   enabled,
			Used:     used,
			Refcnt:   number,
			HostPath: hostPath,
		}
	}
	return
}

func GraphDrivers() (filesystem container.Filesystem, err error) {
	o, err := GraphDriver(&overlay.Overlay{})
	if err != nil {
		return
	}
	d, err := GraphDriver(&devicemapper.DeviceMapper{})
	if err != nil {
		return
	}
	filesystem = container.Filesystem{
		Overlay:      o,
		DeviceMapper: d,
	}
	return
}
