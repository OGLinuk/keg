package main

import (
	"github.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("kn", "datepath", "cd", "config", "add", "tstamp", "html", "jax")
	x.Summary = ``
	x.Version = "1.0.0"
	x.Author = "Rob Muhlestein <rob@rwx.gg> (rob.rwx.gg)"
	x.Git = "github.com/afkworks/kn/cmd/kn/kn"
	x.Copyright = "(c) Rob Muhlestein"
	x.License = "MPL-2"
	x.Description = ``
}
