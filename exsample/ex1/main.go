package main

import (
	"fmt"
	"strings"
	"github.com/yuri-swift/link"
)

var htmlExsample = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`

func main() {
	r = strings.NewReader(htmlExsample)
	links, error := link.Parse(r)
	if error != nil {
		panic(error)
	}
	// %+v:構造体のフィールド名を表示
	fmt.PrinfF("%+v", links)
}