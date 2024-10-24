package main

import (
	"os"

	"github.com/alecthomas/repr"
	"rsc.io/markdown"
)

func main() {
	input := `
# Heading 1

This is a paragraph.

---

This is another paragraph.
`[1:]
	parser := markdown.Parser{}
	doc := parser.Parse(input)
	opts := []repr.Option{
		repr.OmitEmpty(true),
		repr.IgnorePrivate(),
		repr.Hide[markdown.Position](),
	}
	printer := repr.New(os.Stdout, opts...)
	printer.Print(doc)
}
