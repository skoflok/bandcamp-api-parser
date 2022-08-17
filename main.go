package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/skoflok/bandcamp_api_parser/api"
)

func main() {
	flag.Parse()
	page, _ := strconv.Atoi(flag.Arg(0))
	q := api.NewQueryArgs(page)
	r, e := api.FetchReleasesFromHome(q)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(r.Items[0])
}
