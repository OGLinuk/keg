package main

import (
	"fmt"

	"github.com/rwxrob/cmdtab"
	"github.com/rwxrob/keg"
)

func init() {
	x := cmdtab.New("tstamp")
	x.Summary = `outputs an ISO-8601 time stamp`
	x.Description = ``
	x.Method = func(args []string) error {
		fmt.Println(keg.TStamp())
		return nil
	}
}
