package checksec

import (
	"github.com/ctrsploit/ctrsploit/cmd/ctrsploit/env"
	"github.com/ctrsploit/ctrsploit/vul"
	"github.com/ctrsploit/ctrsploit/vul/sys_admin"
	"github.com/ctrsploit/sploit-spec/pkg/app"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:    "checksec",
	Aliases: []string{"c"},
	Usage:   "check security inside a container",
	Subcommands: []*cli.Command{
		Auto,
		env.Command,
		app.Vul2ChecksecCmd(&sys_admin.SysadminCgroupV1, []string{"sys_admin", "release_agent", "ra"}),
		app.Vul2ChecksecCmd(&vul.NetworkNamespaceHostLevel, []string{"host"}),
	},
}
