package main

import (
	"fmt"
	"log"
	"strings"

	link "github.com/learntogo/gophercises/03_linkparser"
)

var exampleHTML = `
<html>
	<body>
		<h1>Hello!</h1>
		<a href="/some-page">A link to some page</a>
		<a href="/other-page">A link to another page</a>
	</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := link.Parse(r)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Printf("%v\n", links)
}
