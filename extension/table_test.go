package extension

import (
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/testutil"
)

func TestTable(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			Table,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/table.txt", t, testutil.ParseCliCaseArg()...)
}
