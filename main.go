package main

import (
	"fmt"

	"github.com/skoflok/bandcamp_api_parser/api"
)

func main() {
	q := api.NewQueryArgs(2)
	r, e := api.FetchReleasesFromHome(q)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(r.Items[0])
}
