package mountinfo

import (
	pkgMountinfo "github.com/ctrsploit/ctrsploit/pkg/mountinfo"
	"github.com/moby/sys/mountinfo"
)

func MountInfo() (machine []*mountinfo.Info, err error) {
	return pkgMountinfo.MountInfo()
}
