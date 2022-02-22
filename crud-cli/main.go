package main

import (
	"log"
	"github.com/dixonwille/wmenu/v5"
)

func main() {
	menu := wmenu.NewMenu("What would you like to do?")
	menu.Action(func(opts []wmenu.Opt) error {})
}
