package markdown

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func MdToHtml(str string) ([]byte, error) {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	par := parser.NewWithExtensions(extensions)

	md := []byte("## markdown text")
	html := markdown.ToHTML(md, par, nil)

	return html, nil
}

func HtmlToMd(str string) error {
	return nil
}
