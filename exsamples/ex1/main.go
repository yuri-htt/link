package main

import (
	"fmt"
	"link"
	"strings"
)

var exsampleHtml = `
<html>
	<body>
		<h1>h1 tag text</h1>
		<a href="/other-page">
			A link tag text
			<span> some span </span>
		</a>
		<a href="/other-page-2">A link tag text</a>
	</body>
</html>
`

func main() {
	// strings package > func NewReader(s string) *Reader
	// type Reader: io.Reader interfaceを実装
	r := strings.NewReader(exsampleHtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	// %+v:構造体のフィールド名を表示
	fmt.Printf("%+v", links)
}
