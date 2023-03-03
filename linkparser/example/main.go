package main

import (
	"fmt"
	"strings"

	"github.com/EdwinVesga/gophercises/linkparser/parser"
)

const htmlexample = `
<html>
  <body>
    <a href="/dog-cat"
      >dog cat
      <!-- commented text SHOULD NOT be included! --></a
    >
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
