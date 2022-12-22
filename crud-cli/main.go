package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "world", "A name to say hello to.")
var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "use spanish language.")
	flag.BoolVar(&spanish, "s", false, "use spanish language.")
}
func main() {
	flag.Parse()

	if spanish {
		fmt.Printf("hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s\n", *name)
	}
}
