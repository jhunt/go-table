package main

import (
	"os"

	"github.com/jhunt/go-table"
)

func main() {
	t := table.NewTable("#", "Name", "Notes", "Status")
	t.Row(nil, 1, "Foo", "lorem ipsum dolor sit amet...", "GOOD")
	t.Row(nil, 2, "Bar", "you can even have\nembedded newlines...", "GOOD")
	t.Output(os.Stdout)
}
