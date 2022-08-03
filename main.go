package main

import (
	"fmt"

	"github.com/skoflok/bandcamp_api_parser/bandcamp"
)

func main() {
	q := bandcamp.NewQueryArgs(2)
	r, e := bandcamp.FetchReleasesFromHome(q)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(r.Items[0])
}
