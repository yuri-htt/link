/*
HTMLから構造を探索し、Linkを見つけてパースする
*/
package Link

import (
	"io"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}