package env

import (
	"github.com/ctrsploit/ctrsploit/env/mountinfo"
	"github.com/urfave/cli/v2"
)

var Mountinfo = &cli.Command{
	Name:    "mountinfo",
	Aliases: []string{"m"},
	Usage:   "list mount points",
	Action: func(c *cli.Context) (err error) {
		err = mountinfo.Print()
		if err != nil {
			return
		}
		return
	},
}
