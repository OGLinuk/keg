package md

import (
	"fmt"
	"io"

	mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func ToHTML(md []byte, buf io.Writer) error {
	// TODO detect math in the input
	mjax := `
<script src="https://polyfill.io/v3/polyfill.min.js?features=es6"></script>
<script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script>
`
	fmt.Fprintln(buf, mjax)

	gold := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithExtensions(mathjax.MathJax),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
		),
	)
	if err := gold.Convert(md, buf); err != nil {
		return err
	}
	return nil
}
