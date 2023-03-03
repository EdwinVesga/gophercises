package main

import (
	"fmt"
	"strings"

	"github.com/EdwinVesga/gophercises/linkparser/parser"
)

const htmlexample = `
<html>
  <body>
    <h1>Hello!</h1>
    <a href="/other-page">A link to another page</a>
  </body>
</html>
`

func main() {
	r := strings.NewReader(htmlexample)

	links, err := parser.Parse(r)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
