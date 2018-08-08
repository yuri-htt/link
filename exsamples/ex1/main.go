package main

import (
	"fmt"
	"link"
	"strings"
)

var exsampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exsampleHtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	// %+v:構造体のフィールド名を表示
	fmt.Printf("%+v", links)
}
