package main

import (
	"fmt"

	"github.com/skoflok/bandcamp_api_parser/bandcamp"
)

func main() {
	fmt.Println("stat")
	fmt.Println(bandcamp.Hell())
	fmt.Println(bandcamp.NewQueryArgs(2).Encode())
}
