package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/rwxrob/cmdtab"
	"github.com/rwxrob/keg/md"
)

func init() {
	x := cmdtab.New("html")
	x.Summary = `converts KEG Mark to HTML`
	x.Usage = `[<file>]`
	x.Description = `
		Converts KEG Mark (Markdown) into HTML. Reads from stdin if no *file*
		argument passed.`

	x.Method = func(args []string) (err error) {
		var mark []byte
		var buf bytes.Buffer

		switch {
		case len(args) >= 1:
			mark, err = os.ReadFile(args[0])
		default:
			mark, err = io.ReadAll(os.Stdin)
		}
		if err != nil {
			return err
		}

		err = md.ToHTML(mark, &buf)
		if err != nil {
			return err
		}
		fmt.Print(buf.String())

		return nil
	}
}
