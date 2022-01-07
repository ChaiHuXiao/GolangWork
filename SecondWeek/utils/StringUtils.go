package main

import (
	"bytes"
)

func main() {
	a.say()
}

func StringSplicing(cap int, joins ...string) string {
	buffer := bytes.Buffer{}
	buffer.Grow(cap)
	for _, v := range joins {
		buffer.WriteString(v)
	}
	return buffer.String()
}
